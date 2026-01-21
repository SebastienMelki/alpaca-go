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
	)

	return &Client{client}
}

// NewPaperClient creates a client for paper trading.
func NewPaperClient(apiKey, apiSecret string, opts ...Option) *Client {
	return NewClient(apiKey, apiSecret, append(opts, WithBaseURL(PaperBaseURL))...)
}

// Re-export core model types.
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

// Re-export enums.
type (
	OrderSide         = tradingv1.OrderSide
	OrderType         = tradingv1.OrderType
	OrderStatus       = tradingv1.OrderStatus
	TimeInForce       = tradingv1.TimeInForce
	AssetClass        = tradingv1.AssetClass
	AssetStatus       = tradingv1.AssetStatus
	PositionSide      = tradingv1.PositionSide
	OrderClass        = tradingv1.OrderClass
	ActivityType      = tradingv1.ActivityType
	DtbpCheck         = tradingv1.DtbpCheck
	TradeConfirmEmail = tradingv1.TradeConfirmEmail
	AccountStatus     = tradingv1.AccountStatus
)

// Re-export request types.
type (
	GetAccountRequest                  = tradingv1.GetAccountRequest
	GetAccountConfigurationsRequest    = tradingv1.GetAccountConfigurationsRequest
	UpdateAccountConfigurationsRequest = tradingv1.UpdateAccountConfigurationsRequest
	GetPortfolioHistoryRequest         = tradingv1.GetPortfolioHistoryRequest
	GetAccountActivitiesRequest        = tradingv1.GetAccountActivitiesRequest
	GetAccountActivitiesByTypeRequest  = tradingv1.GetAccountActivitiesByTypeRequest
	CreateOrderRequest                 = tradingv1.CreateOrderRequest
	ListOrdersRequest                  = tradingv1.ListOrdersRequest
	GetOrderRequest                    = tradingv1.GetOrderRequest
	GetOrderByClientIdRequest          = tradingv1.GetOrderByClientIdRequest
	ReplaceOrderRequest                = tradingv1.ReplaceOrderRequest
	CancelOrderRequest                 = tradingv1.CancelOrderRequest
	CancelAllOrdersRequest             = tradingv1.CancelAllOrdersRequest
	ListPositionsRequest               = tradingv1.ListPositionsRequest
	GetPositionRequest                 = tradingv1.GetPositionRequest
	ClosePositionRequest               = tradingv1.ClosePositionRequest
	CloseAllPositionsRequest           = tradingv1.CloseAllPositionsRequest
	ExerciseOptionRequest              = tradingv1.ExerciseOptionRequest
	ListAssetsRequest                  = tradingv1.ListAssetsRequest
	GetAssetRequest                    = tradingv1.GetAssetRequest
	GetClockRequest                    = tradingv1.GetClockRequest
	GetCalendarRequest                 = tradingv1.GetCalendarRequest
	ListWatchlistsRequest              = tradingv1.ListWatchlistsRequest
	CreateWatchlistRequest             = tradingv1.CreateWatchlistRequest
	GetWatchlistRequest                = tradingv1.GetWatchlistRequest
	UpdateWatchlistRequest             = tradingv1.UpdateWatchlistRequest
	DeleteWatchlistRequest             = tradingv1.DeleteWatchlistRequest
	AddWatchlistAssetRequest           = tradingv1.AddWatchlistAssetRequest
	RemoveWatchlistAssetRequest        = tradingv1.RemoveWatchlistAssetRequest
)

// Re-export response types.
type (
	GetAccountActivitiesResponse = tradingv1.GetAccountActivitiesResponse
	ListOrdersResponse           = tradingv1.ListOrdersResponse
	CancelOrderResponse          = tradingv1.CancelOrderResponse
	CancelAllOrdersResponse      = tradingv1.CancelAllOrdersResponse
	ListPositionsResponse        = tradingv1.ListPositionsResponse
	CloseAllPositionsResponse    = tradingv1.CloseAllPositionsResponse
	ExerciseOptionResponse       = tradingv1.ExerciseOptionResponse
	ListAssetsResponse           = tradingv1.ListAssetsResponse
	GetCalendarResponse          = tradingv1.GetCalendarResponse
	ListWatchlistsResponse       = tradingv1.ListWatchlistsResponse
	DeleteWatchlistResponse      = tradingv1.DeleteWatchlistResponse
	RemoveWatchlistAssetResponse = tradingv1.RemoveWatchlistAssetResponse
)

// Re-export nested types.
type (
	TakeProfitSpec = tradingv1.TakeProfitSpec
	StopLossSpec   = tradingv1.StopLossSpec
)

// Enum value constants for OrderSide.
const (
	OrderSide_ORDER_SIDE_UNSPECIFIED = tradingv1.OrderSide_ORDER_SIDE_UNSPECIFIED
	OrderSide_ORDER_SIDE_BUY         = tradingv1.OrderSide_ORDER_SIDE_BUY
	OrderSide_ORDER_SIDE_SELL        = tradingv1.OrderSide_ORDER_SIDE_SELL
)

// Enum value constants for OrderType.
const (
	OrderType_ORDER_TYPE_UNSPECIFIED   = tradingv1.OrderType_ORDER_TYPE_UNSPECIFIED
	OrderType_ORDER_TYPE_MARKET        = tradingv1.OrderType_ORDER_TYPE_MARKET
	OrderType_ORDER_TYPE_LIMIT         = tradingv1.OrderType_ORDER_TYPE_LIMIT
	OrderType_ORDER_TYPE_STOP          = tradingv1.OrderType_ORDER_TYPE_STOP
	OrderType_ORDER_TYPE_STOP_LIMIT    = tradingv1.OrderType_ORDER_TYPE_STOP_LIMIT
	OrderType_ORDER_TYPE_TRAILING_STOP = tradingv1.OrderType_ORDER_TYPE_TRAILING_STOP
)

// Enum value constants for TimeInForce.
const (
	TimeInForce_TIME_IN_FORCE_UNSPECIFIED = tradingv1.TimeInForce_TIME_IN_FORCE_UNSPECIFIED
	TimeInForce_TIME_IN_FORCE_DAY         = tradingv1.TimeInForce_TIME_IN_FORCE_DAY
	TimeInForce_TIME_IN_FORCE_GTC         = tradingv1.TimeInForce_TIME_IN_FORCE_GTC
	TimeInForce_TIME_IN_FORCE_OPG         = tradingv1.TimeInForce_TIME_IN_FORCE_OPG
	TimeInForce_TIME_IN_FORCE_CLS         = tradingv1.TimeInForce_TIME_IN_FORCE_CLS
	TimeInForce_TIME_IN_FORCE_IOC         = tradingv1.TimeInForce_TIME_IN_FORCE_IOC
	TimeInForce_TIME_IN_FORCE_FOK         = tradingv1.TimeInForce_TIME_IN_FORCE_FOK
)

// Enum value constants for OrderClass.
const (
	OrderClass_ORDER_CLASS_UNSPECIFIED = tradingv1.OrderClass_ORDER_CLASS_UNSPECIFIED
	OrderClass_ORDER_CLASS_SIMPLE      = tradingv1.OrderClass_ORDER_CLASS_SIMPLE
	OrderClass_ORDER_CLASS_BRACKET     = tradingv1.OrderClass_ORDER_CLASS_BRACKET
	OrderClass_ORDER_CLASS_OCO         = tradingv1.OrderClass_ORDER_CLASS_OCO
	OrderClass_ORDER_CLASS_OTO         = tradingv1.OrderClass_ORDER_CLASS_OTO
)
