package marketdata

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	marketdatav1beta "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/marketdata/v1beta"
)

var (
	errSymbolRequired = errors.New("symbol is required")
)

// GetLogo retrieves the company logo image for a symbol.
// Custom implementation for binary PNG response handling.
func (c *Client) GetLogo(ctx context.Context, req *marketdatav1beta.GetLogoRequest) (*marketdatav1beta.GetLogoResponse, error) {
	if req.Symbol == "" {
		return nil, errSymbolRequired
	}

	path := "/v1beta1/logos/" + url.PathEscape(req.Symbol)
	reqURL := c.baseURL + path

	if req.Placeholder {
		reqURL += "?placeholder=true"
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("APCA-API-KEY-ID", c.apiKey)
	httpReq.Header.Set("APCA-API-SECRET-KEY", c.apiSecret)

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	//nolint:errcheck // Body close errors in defer are not critical
	defer resp.Body.Close()

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, &logoHTTPError{
			StatusCode: resp.StatusCode,
			Body:       string(imageData),
		}
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "image/png"
	}

	return &marketdatav1beta.GetLogoResponse{
		Image:       imageData,
		ContentType: contentType,
	}, nil
}

type logoHTTPError struct {
	StatusCode int
	Body       string
}

func (e *logoHTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Body)
}
