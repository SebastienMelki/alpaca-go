package marketdata

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	marketdatav1beta "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/marketdata/v1beta"
)

func TestGetLogo(t *testing.T) {
	tests := []struct {
		name           string
		symbol         string
		placeholder    bool
		serverResponse func(w http.ResponseWriter, r *http.Request)
		wantErr        bool
		errContains    string
		checkResponse  func(t *testing.T, resp *marketdatav1beta.GetLogoResponse)
	}{
		{
			name:        "success with valid symbol",
			symbol:      "AAPL",
			placeholder: false,
			serverResponse: func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/v1beta1/logos/AAPL" {
					t.Errorf("unexpected path: %s", r.URL.Path)
				}
				if r.Header.Get("APCA-API-KEY-ID") != "test-key" {
					t.Errorf("missing or incorrect API key header")
				}
				if r.Header.Get("APCA-API-SECRET-KEY") != "test-secret" {
					t.Errorf("missing or incorrect API secret header")
				}
				w.Header().Set("Content-Type", "image/png")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte{0x89, 0x50, 0x4E, 0x47}) // PNG magic bytes
			},
			wantErr: false,
			checkResponse: func(t *testing.T, resp *marketdatav1beta.GetLogoResponse) {
				if len(resp.Image) != 4 {
					t.Errorf("expected 4 bytes, got %d", len(resp.Image))
				}
				if resp.ContentType != "image/png" {
					t.Errorf("expected content type image/png, got %s", resp.ContentType)
				}
			},
		},
		{
			name:        "success with placeholder parameter",
			symbol:      "AAPL",
			placeholder: true,
			serverResponse: func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/v1beta1/logos/AAPL" {
					t.Errorf("unexpected path: %s", r.URL.Path)
				}
				if r.URL.Query().Get("placeholder") != "true" {
					t.Errorf("expected placeholder=true query parameter")
				}
				w.Header().Set("Content-Type", "image/png")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte{0x89, 0x50, 0x4E, 0x47})
			},
			wantErr: false,
			checkResponse: func(t *testing.T, resp *marketdatav1beta.GetLogoResponse) {
				if len(resp.Image) != 4 {
					t.Errorf("expected 4 bytes, got %d", len(resp.Image))
				}
			},
		},
		{
			name:        "missing symbol validation",
			symbol:      "",
			placeholder: false,
			serverResponse: func(w http.ResponseWriter, r *http.Request) {
				t.Error("should not make request with empty symbol")
			},
			wantErr:     true,
			errContains: "symbol is required",
		},
		{
			name:        "404 not found",
			symbol:      "INVALID",
			placeholder: false,
			serverResponse: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("symbol not found"))
			},
			wantErr:     true,
			errContains: "HTTP 404",
		},
		{
			name:        "401 unauthorized",
			symbol:      "AAPL",
			placeholder: false,
			serverResponse: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("invalid credentials"))
			},
			wantErr:     true,
			errContains: "HTTP 401",
		},
		{
			name:        "default content type when missing",
			symbol:      "AAPL",
			placeholder: false,
			serverResponse: func(w http.ResponseWriter, r *http.Request) {
				// Explicitly clear Content-Type header
				w.Header().Del("Content-Type")
				w.Write([]byte{0x89, 0x50, 0x4E, 0x47})
			},
			wantErr: false,
			checkResponse: func(t *testing.T, resp *marketdatav1beta.GetLogoResponse) {
				// httptest server may add a default Content-Type, so we just check
				// that our code handles the case when it's empty
				if resp.ContentType == "" {
					t.Errorf("content type should not be empty")
				}
			},
		},
		{
			name:        "symbol with special characters",
			symbol:      "BRK.B",
			placeholder: false,
			serverResponse: func(w http.ResponseWriter, r *http.Request) {
				// URL should be escaped
				if r.URL.Path != "/v1beta1/logos/BRK.B" {
					t.Errorf("unexpected path: %s", r.URL.Path)
				}
				w.Header().Set("Content-Type", "image/png")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte{0x89, 0x50, 0x4E, 0x47})
			},
			wantErr: false,
			checkResponse: func(t *testing.T, resp *marketdatav1beta.GetLogoResponse) {
				if len(resp.Image) != 4 {
					t.Errorf("expected 4 bytes, got %d", len(resp.Image))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.serverResponse))
			defer server.Close()

			client := NewClient(
				"test-key",
				"test-secret",
				WithBaseURL(server.URL),
			)

			resp, err := client.GetLogo(context.Background(), &marketdatav1beta.GetLogoRequest{
				Symbol:      tt.symbol,
				Placeholder: tt.placeholder,
			})

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
					return
				}
				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("expected error to contain %q, got %q", tt.errContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, resp)
			}
		})
	}
}

func TestGetLogoContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		<-r.Context().Done()
	}))
	defer server.Close()

	client := NewClient(
		"test-key",
		"test-secret",
		WithBaseURL(server.URL),
	)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := client.GetLogo(ctx, &marketdatav1beta.GetLogoRequest{
		Symbol: "AAPL",
	})

	if err == nil {
		t.Error("expected context cancellation error, got nil")
	}
	if !strings.Contains(err.Error(), "context canceled") {
		t.Errorf("expected context canceled error, got: %v", err)
	}
}

func TestGetLogoReadError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1000") // Lie about content length
		w.WriteHeader(http.StatusOK)
		// Write less data than promised, then close connection
		w.Write([]byte{0x89})
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
	}))
	defer server.Close()

	client := NewClient(
		"test-key",
		"test-secret",
		WithBaseURL(server.URL),
		WithHTTPClient(&http.Client{
			Transport: &brokenBodyTransport{server.URL},
		}),
	)

	_, err := client.GetLogo(context.Background(), &marketdatav1beta.GetLogoRequest{
		Symbol: "AAPL",
	})

	if err == nil {
		t.Error("expected read error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to read response") {
		t.Errorf("expected read error, got: %v", err)
	}
}

// brokenBodyTransport simulates a connection error during body read.
type brokenBodyTransport struct {
	baseURL string
}

func (t *brokenBodyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(&brokenReader{}),
		Header:     make(http.Header),
	}, nil
}

type brokenReader struct{}

func (r *brokenReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrUnexpectedEOF
}
