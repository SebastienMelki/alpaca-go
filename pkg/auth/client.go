package auth

import (
	"net/http"

	authv1 "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/auth/v1"
)

const (
	BaseURL = "https://api.alpaca.markets"
)

// Client wraps the generated AuthService client with Alpaca-specific defaults.
type Client struct {
	authv1.AuthServiceClient
}

// Option configures a Client.
type Option func(*options)

type options struct {
	httpClient *http.Client
	baseURL    string
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(c *http.Client) Option {
	return func(o *options) { o.httpClient = c }
}

// WithBaseURL sets a custom base URL (defaults to BaseURL).
func WithBaseURL(url string) Option {
	return func(o *options) { o.baseURL = url }
}

// NewClient creates a new Auth API client.
// Note: The Auth API is typically used for OAuth2 token issuance.
func NewClient(opts ...Option) *Client {
	cfg := &options{
		httpClient: http.DefaultClient,
		baseURL:    BaseURL,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	client := authv1.NewAuthServiceClient(
		cfg.baseURL,
		authv1.WithAuthServiceHTTPClient(cfg.httpClient),
	)

	return &Client{client}
}

// =============================================================================
// Request/Response Types
// =============================================================================

type (
	IssueTokenRequest  = authv1.IssueTokenRequest
	IssueTokenResponse = authv1.IssueTokenResponse
)

