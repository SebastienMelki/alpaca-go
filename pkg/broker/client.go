package broker

import (
	"encoding/base64"
	"net/http"

	brokerv1 "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/broker/v1"
)

const (
	LiveBaseURL    = "https://broker-api.alpaca.markets"
	SandboxBaseURL = "https://broker-api.sandbox.alpaca.markets"
)

// Client wraps the generated BrokerService client with Alpaca-specific defaults.
type Client struct {
	brokerv1.BrokerServiceClient
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

// WithBaseURL sets a custom base URL (defaults to LiveBaseURL).
func WithBaseURL(url string) Option {
	return func(o *options) { o.baseURL = url }
}

// NewClient creates a new Broker API client.
// The Broker API uses HTTP Basic Auth with API key and secret.
func NewClient(apiKey, apiSecret string, opts ...Option) *Client {
	cfg := &options{
		httpClient: http.DefaultClient,
		baseURL:    LiveBaseURL,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	// Broker API uses Basic Auth
	auth := base64.StdEncoding.EncodeToString([]byte(apiKey + ":" + apiSecret))

	client := brokerv1.NewBrokerServiceClient(
		cfg.baseURL,
		brokerv1.WithBrokerServiceHTTPClient(cfg.httpClient),
		brokerv1.WithBrokerServiceAuthorization("Basic "+auth),
	)

	return &Client{client}
}

// NewSandboxClient creates a client for the broker sandbox environment.
func NewSandboxClient(apiKey, apiSecret string, opts ...Option) *Client {
	return NewClient(apiKey, apiSecret, append(opts, WithBaseURL(SandboxBaseURL))...)
}

// Re-export core model types.
type (
	BrokerAccount   = brokerv1.BrokerAccount
	BrokerOrder     = brokerv1.BrokerOrder
	BrokerPosition  = brokerv1.BrokerPosition
	ACHRelationship = brokerv1.ACHRelationship
	Transfer        = brokerv1.Transfer
)

// Re-export account-related types.
type (
	Contact        = brokerv1.Contact
	Identity       = brokerv1.Identity
	Disclosures    = brokerv1.Disclosures
	Agreement      = brokerv1.Agreement
	TrustedContact = brokerv1.TrustedContact
)

// Re-export request types.
type (
	CreateAccountRequest            = brokerv1.CreateAccountRequest
	ListAccountsRequest             = brokerv1.ListAccountsRequest
	GetBrokerAccountRequest         = brokerv1.GetBrokerAccountRequest
	UpdateBrokerAccountRequest      = brokerv1.UpdateBrokerAccountRequest
	CloseBrokerAccountRequest       = brokerv1.CloseBrokerAccountRequest
	CreateACHRelationshipRequest    = brokerv1.CreateACHRelationshipRequest
	ListACHRelationshipsRequest     = brokerv1.ListACHRelationshipsRequest
	DeleteACHRelationshipRequest    = brokerv1.DeleteACHRelationshipRequest
	CreateTransferRequest           = brokerv1.CreateTransferRequest
	ListTransfersRequest            = brokerv1.ListTransfersRequest
	GetTransferRequest              = brokerv1.GetTransferRequest
	CancelTransferRequest           = brokerv1.CancelTransferRequest
	CreateTradingOrderRequest       = brokerv1.CreateTradingOrderRequest
	ListTradingOrdersRequest        = brokerv1.ListTradingOrdersRequest
	GetTradingOrderRequest          = brokerv1.GetTradingOrderRequest
	CancelTradingOrderRequest       = brokerv1.CancelTradingOrderRequest
	ListTradingPositionsRequest     = brokerv1.ListTradingPositionsRequest
	GetTradingPositionRequest       = brokerv1.GetTradingPositionRequest
	CloseTradingPositionRequest     = brokerv1.CloseTradingPositionRequest
	CloseAllTradingPositionsRequest = brokerv1.CloseAllTradingPositionsRequest
)

// Re-export response types.
type (
	ListAccountsResponse             = brokerv1.ListAccountsResponse
	CloseBrokerAccountResponse       = brokerv1.CloseBrokerAccountResponse
	ListACHRelationshipsResponse     = brokerv1.ListACHRelationshipsResponse
	DeleteACHRelationshipResponse    = brokerv1.DeleteACHRelationshipResponse
	ListTransfersResponse            = brokerv1.ListTransfersResponse
	CancelTransferResponse           = brokerv1.CancelTransferResponse
	ListTradingOrdersResponse        = brokerv1.ListTradingOrdersResponse
	CancelTradingOrderResponse       = brokerv1.CancelTradingOrderResponse
	ListTradingPositionsResponse     = brokerv1.ListTradingPositionsResponse
	CloseAllTradingPositionsResponse = brokerv1.CloseAllTradingPositionsResponse
)

// Re-export enums.
type (
	TransferDirection     = brokerv1.TransferDirection
	TransferType          = brokerv1.TransferType
	TransferStatus        = brokerv1.TransferStatus
	ACHRelationshipStatus = brokerv1.ACHRelationshipStatus
)
