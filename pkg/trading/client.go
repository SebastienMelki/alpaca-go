package trading

import (
	"net/http"

	tradingv1 "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/trading/v1"
)

const (
	LiveBaseURL  = "https://api.alpaca.markets"
	PaperBaseURL = "https://paper-api.alpaca.markets"
)

// Client wraps the generated TradingService client with Alpaca-specific defaults.
type Client struct {
	tradingv1.TradingServiceClient
}

// Option configures a Client.
type Option func(*options)

type options struct {
	httpClient           *http.Client
	baseURL              string
	discardUnknownFields bool
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(c *http.Client) Option {
	return func(o *options) { o.httpClient = c }
}

// WithBaseURL sets a custom base URL (defaults to LiveBaseURL).
func WithBaseURL(url string) Option {
	return func(o *options) { o.baseURL = url }
}

// WithDiscardUnknownFields sets whether to discard unknown fields in JSON responses.
// When true, unknown fields are silently ignored instead of causing unmarshal errors.
func WithDiscardUnknownFields(discard bool) Option {
	return func(o *options) { o.discardUnknownFields = discard }
}

// NewClient creates a new Trading API client.
func NewClient(apiKey, apiSecret string, opts ...Option) *Client {
	cfg := &options{
		httpClient: http.DefaultClient,
		baseURL:    LiveBaseURL,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	client := tradingv1.NewTradingServiceClient(
		cfg.baseURL,
		tradingv1.WithTradingServiceHTTPClient(cfg.httpClient),
		tradingv1.WithTradingServiceAPCAAPIKEYID(apiKey),
		tradingv1.WithTradingServiceAPCAAPISECRETKEY(apiSecret),
		tradingv1.WithTradingServiceDiscardUnknownFields(cfg.discardUnknownFields),
	)

	return &Client{client}
}

// NewPaperClient creates a client for paper trading.
func NewPaperClient(apiKey, apiSecret string, opts ...Option) *Client {
	return NewClient(apiKey, apiSecret, append(opts, WithBaseURL(PaperBaseURL))...)
}

// =============================================================================
// Call Option Types
// =============================================================================

// CallOption configures a single RPC call to the Trading API.
type CallOption = tradingv1.TradingServiceCallOption

// WithCallDiscardUnknownFields sets whether to discard unknown fields for a single request.
var WithCallDiscardUnknownFields = tradingv1.WithTradingServiceCallDiscardUnknownFields

// =============================================================================
// Core Model Types
// =============================================================================

type (
	Account               = tradingv1.Account
	AccountConfigurations = tradingv1.AccountConfigurations
	Order                 = tradingv1.Order
	Position              = tradingv1.Position
	Asset                 = tradingv1.Asset
	Clock                 = tradingv1.Clock
	CalendarDay           = tradingv1.CalendarDay
	Watchlist             = tradingv1.Watchlist
	PortfolioHistory      = tradingv1.PortfolioHistory
	AccountActivity       = tradingv1.AccountActivity
)

// =============================================================================
// Order Nested Types
// =============================================================================

type (
	TakeProfitSpec       = tradingv1.TakeProfitSpec
	StopLossSpec         = tradingv1.StopLossSpec
	MLegOrderLeg         = tradingv1.MLegOrderLeg
	AdvancedInstructions = tradingv1.AdvancedInstructions
	CanceledOrder        = tradingv1.CanceledOrder
	ClosedPosition       = tradingv1.ClosedPosition
)

// =============================================================================
// Option Contract Types
// =============================================================================

type (
	OptionContract    = tradingv1.OptionContract
	OptionDeliverable = tradingv1.OptionDeliverable
)

// =============================================================================
// Crypto Funding Types
// =============================================================================

type (
	CryptoWallet           = tradingv1.CryptoWallet
	CryptoTransfer         = tradingv1.CryptoTransfer
	WhitelistedAddress     = tradingv1.WhitelistedAddress
	CryptoTransferEstimate = tradingv1.CryptoTransferEstimate
)

// =============================================================================
// Note: Enums have been converted to strings for better API compatibility
// =============================================================================

// =============================================================================
// Account Request/Response Types
// =============================================================================

type (
	GetAccountRequest                  = tradingv1.GetAccountRequest
	GetAccountConfigurationsRequest    = tradingv1.GetAccountConfigurationsRequest
	UpdateAccountConfigurationsRequest = tradingv1.UpdateAccountConfigurationsRequest
	GetPortfolioHistoryRequest         = tradingv1.GetPortfolioHistoryRequest
	GetAccountActivitiesRequest        = tradingv1.GetAccountActivitiesRequest
	GetAccountActivitiesByTypeRequest  = tradingv1.GetAccountActivitiesByTypeRequest
	GetAccountActivitiesResponse       = tradingv1.GetAccountActivitiesResponse
)

// =============================================================================
// Order Request/Response Types
// =============================================================================

type (
	CreateOrderRequest      = tradingv1.CreateOrderRequest
	ListOrdersRequest       = tradingv1.ListOrdersRequest
	ListOrdersResponse      = tradingv1.ListOrdersResponse
	GetOrderRequest         = tradingv1.GetOrderRequest
	GetOrderByClientIdRequest = tradingv1.GetOrderByClientIdRequest
	ReplaceOrderRequest     = tradingv1.ReplaceOrderRequest
	CancelOrderRequest      = tradingv1.CancelOrderRequest
	CancelOrderResponse     = tradingv1.CancelOrderResponse
	CancelAllOrdersRequest  = tradingv1.CancelAllOrdersRequest
	CancelAllOrdersResponse = tradingv1.CancelAllOrdersResponse
)

// =============================================================================
// Position Request/Response Types
// =============================================================================

type (
	ListPositionsRequest      = tradingv1.ListPositionsRequest
	ListPositionsResponse     = tradingv1.ListPositionsResponse
	GetPositionRequest        = tradingv1.GetPositionRequest
	ClosePositionRequest      = tradingv1.ClosePositionRequest
	CloseAllPositionsRequest  = tradingv1.CloseAllPositionsRequest
	CloseAllPositionsResponse = tradingv1.CloseAllPositionsResponse
)

// =============================================================================
// Option Exercise Request/Response Types
// =============================================================================

type (
	ExerciseOptionRequest        = tradingv1.ExerciseOptionRequest
	ExerciseOptionResponse       = tradingv1.ExerciseOptionResponse
	DoNotExerciseOptionRequest   = tradingv1.DoNotExerciseOptionRequest
	DoNotExerciseOptionResponse  = tradingv1.DoNotExerciseOptionResponse
)

// =============================================================================
// Option Contract Request/Response Types
// =============================================================================

type (
	ListOptionContractsRequest  = tradingv1.ListOptionContractsRequest
	ListOptionContractsResponse = tradingv1.ListOptionContractsResponse
	GetOptionContractRequest    = tradingv1.GetOptionContractRequest
)

// =============================================================================
// Asset Request/Response Types
// =============================================================================

type (
	ListAssetsRequest  = tradingv1.ListAssetsRequest
	ListAssetsResponse = tradingv1.ListAssetsResponse
	GetAssetRequest    = tradingv1.GetAssetRequest
)

// =============================================================================
// Clock/Calendar Request/Response Types
// =============================================================================

type (
	GetClockRequest     = tradingv1.GetClockRequest
	GetCalendarRequest  = tradingv1.GetCalendarRequest
	GetCalendarResponse = tradingv1.GetCalendarResponse
)

// =============================================================================
// Watchlist Request/Response Types
// =============================================================================

type (
	ListWatchlistsRequest        = tradingv1.ListWatchlistsRequest
	ListWatchlistsResponse       = tradingv1.ListWatchlistsResponse
	CreateWatchlistRequest       = tradingv1.CreateWatchlistRequest
	GetWatchlistRequest          = tradingv1.GetWatchlistRequest
	GetWatchlistByNameRequest    = tradingv1.GetWatchlistByNameRequest
	UpdateWatchlistRequest       = tradingv1.UpdateWatchlistRequest
	UpdateWatchlistByNameRequest = tradingv1.UpdateWatchlistByNameRequest
	DeleteWatchlistRequest       = tradingv1.DeleteWatchlistRequest
	DeleteWatchlistByNameRequest = tradingv1.DeleteWatchlistByNameRequest
	DeleteWatchlistResponse      = tradingv1.DeleteWatchlistResponse
	AddWatchlistAssetRequest     = tradingv1.AddWatchlistAssetRequest
	AddWatchlistAssetByNameRequest = tradingv1.AddWatchlistAssetByNameRequest
	RemoveWatchlistAssetRequest  = tradingv1.RemoveWatchlistAssetRequest
	RemoveWatchlistAssetResponse = tradingv1.RemoveWatchlistAssetResponse
)

// =============================================================================
// Crypto Funding Request/Response Types
// =============================================================================

type (
	ListCryptoWalletsRequest         = tradingv1.ListCryptoWalletsRequest
	ListCryptoWalletsResponse        = tradingv1.ListCryptoWalletsResponse
	ListCryptoTransfersRequest       = tradingv1.ListCryptoTransfersRequest
	ListCryptoTransfersResponse      = tradingv1.ListCryptoTransfersResponse
	GetCryptoTransferRequest         = tradingv1.GetCryptoTransferRequest
	CreateCryptoTransferRequest      = tradingv1.CreateCryptoTransferRequest
	ListWhitelistedAddressesRequest  = tradingv1.ListWhitelistedAddressesRequest
	ListWhitelistedAddressesResponse = tradingv1.ListWhitelistedAddressesResponse
	CreateWhitelistedAddressRequest  = tradingv1.CreateWhitelistedAddressRequest
	DeleteWhitelistedAddressRequest  = tradingv1.DeleteWhitelistedAddressRequest
	DeleteWhitelistedAddressResponse = tradingv1.DeleteWhitelistedAddressResponse
	GetCryptoTransferEstimateRequest = tradingv1.GetCryptoTransferEstimateRequest
)

// =============================================================================
// String Constants (formerly enums)
// =============================================================================

// OrderSide string constants.
const (
	OrderSide_BUY  = "buy"
	OrderSide_SELL = "sell"
)

// OrderType string constants.
const (
	OrderType_MARKET        = "market"
	OrderType_LIMIT         = "limit"
	OrderType_STOP          = "stop"
	OrderType_STOP_LIMIT    = "stop_limit"
	OrderType_TRAILING_STOP = "trailing_stop"
)

// TimeInForce string constants.
const (
	TimeInForce_DAY = "day"
	TimeInForce_GTC = "gtc"
	TimeInForce_OPG = "opg"
	TimeInForce_CLS = "cls"
	TimeInForce_IOC = "ioc"
	TimeInForce_FOK = "fok"
)

// OrderClass string constants.
const (
	OrderClass_SIMPLE  = "simple"
	OrderClass_BRACKET = "bracket"
	OrderClass_OCO     = "oco"
	OrderClass_OTO     = "oto"
	OrderClass_MLEG    = "mleg"
)

// OrderStatus string constants.
const (
	OrderStatus_NEW                = "new"
	OrderStatus_PARTIALLY_FILLED   = "partially_filled"
	OrderStatus_FILLED             = "filled"
	OrderStatus_DONE_FOR_DAY       = "done_for_day"
	OrderStatus_CANCELED           = "canceled"
	OrderStatus_EXPIRED            = "expired"
	OrderStatus_REPLACED           = "replaced"
	OrderStatus_PENDING_CANCEL     = "pending_cancel"
	OrderStatus_PENDING_REPLACE    = "pending_replace"
	OrderStatus_PENDING_NEW        = "pending_new"
	OrderStatus_ACCEPTED           = "accepted"
	OrderStatus_ACCEPTED_FOR_BIDDING = "accepted_for_bidding"
	OrderStatus_STOPPED            = "stopped"
	OrderStatus_REJECTED           = "rejected"
	OrderStatus_SUSPENDED          = "suspended"
	OrderStatus_CALCULATED         = "calculated"
	OrderStatus_HELD               = "held"
)

// PositionSide string constants.
const (
	PositionSide_LONG  = "long"
	PositionSide_SHORT = "short"
)

// PositionIntent string constants.
const (
	PositionIntent_BUY_TO_OPEN    = "buy_to_open"
	PositionIntent_BUY_TO_CLOSE   = "buy_to_close"
	PositionIntent_SELL_TO_OPEN   = "sell_to_open"
	PositionIntent_SELL_TO_CLOSE  = "sell_to_close"
)

// AssetClass string constants.
const (
	AssetClass_US_EQUITY = "us_equity"
	AssetClass_CRYPTO    = "crypto"
	AssetClass_US_OPTION = "us_option"
)

// AccountStatus string constants.
const (
	AccountStatus_ONBOARDING        = "ONBOARDING"
	AccountStatus_SUBMISSION_FAILED = "SUBMISSION_FAILED"
	AccountStatus_SUBMITTED         = "SUBMITTED"
	AccountStatus_ACCOUNT_UPDATED   = "ACCOUNT_UPDATED"
	AccountStatus_APPROVAL_PENDING  = "APPROVAL_PENDING"
	AccountStatus_ACTIVE            = "ACTIVE"
	AccountStatus_REJECTED          = "REJECTED"
	AccountStatus_DISABLED          = "DISABLED"
	AccountStatus_ACCOUNT_CLOSED    = "ACCOUNT_CLOSED"
)

// CryptoStatus string constants.
const (
	CryptoStatus_INACTIVE = "INACTIVE"
	CryptoStatus_ACTIVE   = "ACTIVE"
)
