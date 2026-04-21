package marketdata

import (
	"net/http"

	marketdatav1beta "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/marketdata/v1beta"
	marketdatav2 "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/marketdata/v2"
)

const (
	BaseURL        = "https://data.alpaca.markets"
	SandboxBaseURL = "https://data.sandbox.alpaca.markets"
)

// Client wraps the generated MarketDataService v2 and MarketDataBetaService clients with Alpaca-specific defaults.
type Client struct {
	// V2 provides access to stock market data endpoints (v2 API).
	V2 marketdatav2.MarketDataServiceClient

	// V1Beta provides access to crypto, options, news, screener, forex, fixed income, logos, and corporate actions (beta APIs).
	V1Beta marketdatav1beta.MarketDataBetaServiceClient

	// Fields for custom implementations.
	httpClient *http.Client
	baseURL    string
	apiKey     string
	apiSecret  string
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

// WithBaseURL sets a custom base URL (defaults to BaseURL).
func WithBaseURL(url string) Option {
	return func(o *options) { o.baseURL = url }
}

// WithDiscardUnknownFields sets whether to discard unknown fields in JSON responses.
// When true, unknown fields are silently ignored instead of causing unmarshal errors.
func WithDiscardUnknownFields(discard bool) Option {
	return func(o *options) { o.discardUnknownFields = discard }
}

// NewClient creates a new Market Data API client.
func NewClient(apiKey, apiSecret string, opts ...Option) *Client {
	cfg := &options{
		httpClient: http.DefaultClient,
		baseURL:    BaseURL,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	v2Client := marketdatav2.NewMarketDataServiceClient(
		cfg.baseURL,
		marketdatav2.WithMarketDataServiceHTTPClient(cfg.httpClient),
		marketdatav2.WithMarketDataServiceAPCAAPIKEYID(apiKey),
		marketdatav2.WithMarketDataServiceAPCAAPISECRETKEY(apiSecret),
		marketdatav2.WithMarketDataServiceDiscardUnknownFields(cfg.discardUnknownFields),
	)

	v1betaClient := marketdatav1beta.NewMarketDataBetaServiceClient(
		cfg.baseURL,
		marketdatav1beta.WithMarketDataBetaServiceHTTPClient(cfg.httpClient),
		marketdatav1beta.WithMarketDataBetaServiceAPCAAPIKEYID(apiKey),
		marketdatav1beta.WithMarketDataBetaServiceAPCAAPISECRETKEY(apiSecret),
		marketdatav1beta.WithMarketDataBetaServiceDiscardUnknownFields(cfg.discardUnknownFields),
	)

	return &Client{
		V2:         v2Client,
		V1Beta:     v1betaClient,
		httpClient: cfg.httpClient,
		baseURL:    cfg.baseURL,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
	}
}

// NewSandboxClient creates a client for the broker sandbox market data environment.
func NewSandboxClient(apiKey, apiSecret string, opts ...Option) *Client {
	return NewClient(apiKey, apiSecret, append(opts, WithBaseURL(SandboxBaseURL))...)
}

// =============================================================================
// Call Option Types
// =============================================================================

type (
	V2CallOption     = marketdatav2.MarketDataServiceCallOption
	V1BetaCallOption = marketdatav1beta.MarketDataBetaServiceCallOption
)

// Per-call DiscardUnknownFields overrides.
var (
	WithV2CallDiscardUnknownFields     = marketdatav2.WithMarketDataServiceCallDiscardUnknownFields
	WithV1BetaCallDiscardUnknownFields = marketdatav1beta.WithMarketDataBetaServiceCallDiscardUnknownFields
)

// =============================================================================
// Common Stock Types (v2)
// =============================================================================

type (
	Bar           = marketdatav2.Bar
	Trade         = marketdatav2.Trade
	Quote         = marketdatav2.Quote
	Snapshot      = marketdatav2.Snapshot
	Auction       = marketdatav2.Auction
	DailyAuctions = marketdatav2.DailyAuctions
)

// =============================================================================
// News Types (v1beta)
// =============================================================================

type (
	NewsArticle = marketdatav1beta.NewsArticle
	NewsImage   = marketdatav1beta.NewsImage
)

// =============================================================================
// Crypto Types (v1beta)
// =============================================================================

type (
	CryptoBar            = marketdatav1beta.CryptoBar
	CryptoTrade          = marketdatav1beta.CryptoTrade
	CryptoQuote          = marketdatav1beta.CryptoQuote
	CryptoSnapshot       = marketdatav1beta.CryptoSnapshot
	CryptoOrderbook      = marketdatav1beta.CryptoOrderbook
	CryptoOrderbookEntry = marketdatav1beta.CryptoOrderbookEntry
)

// =============================================================================
// Option Types (v1beta)
// =============================================================================

type (
	OptionBar      = marketdatav1beta.OptionBar
	OptionTrade    = marketdatav1beta.OptionTrade
	OptionQuote    = marketdatav1beta.OptionQuote
	OptionSnapshot = marketdatav1beta.OptionSnapshot
	OptionGreeks   = marketdatav1beta.OptionGreeks
)

// =============================================================================
// Forex Types (v1beta)
// =============================================================================

type (
	ForexRate     = marketdatav1beta.ForexRate
	ForexRateList = marketdatav1beta.ForexRateList
)

// =============================================================================
// Fixed Income Types (v1beta)
// =============================================================================

type (
	FixedIncomePrice = marketdatav1beta.FixedIncomePrice
)

// =============================================================================
// Screener Types (v1beta)
// =============================================================================

type (
	MostActive = marketdatav1beta.MostActive
	Mover      = marketdatav1beta.Mover
)

// =============================================================================
// Corporate Actions Types (v1beta)
// =============================================================================

type (
	CorporateActions   = marketdatav1beta.CorporateActions
	ReverseSplit       = marketdatav1beta.ReverseSplit
	ForwardSplit       = marketdatav1beta.ForwardSplit
	UnitSplit          = marketdatav1beta.UnitSplit
	StockDividend      = marketdatav1beta.StockDividend
	CashDividend       = marketdatav1beta.CashDividend
	SpinOff            = marketdatav1beta.SpinOff
	CashMerger         = marketdatav1beta.CashMerger
	StockMerger        = marketdatav1beta.StockMerger
	StockAndCashMerger = marketdatav1beta.StockAndCashMerger
	Redemption         = marketdatav1beta.Redemption
	NameChange         = marketdatav1beta.NameChange
	WorthlessRemoval   = marketdatav1beta.WorthlessRemoval
	RightsDistribution = marketdatav1beta.RightsDistribution
)

// =============================================================================
// Stock Bars Request/Response Types (v2)
// =============================================================================

type (
	// Multi-symbol requests.
	GetStockBarsRequest        = marketdatav2.GetStockBarsRequest
	GetStockBarsResponse       = marketdatav2.GetStockBarsResponse
	GetLatestStockBarsRequest  = marketdatav2.GetLatestStockBarsRequest
	GetLatestStockBarsResponse = marketdatav2.GetLatestStockBarsResponse

	// Single-symbol requests.
	GetStockBarsSingleRequest       = marketdatav2.GetStockBarsSingleRequest
	GetStockBarsSingleResponse      = marketdatav2.GetStockBarsSingleResponse
	GetLatestStockBarSingleRequest  = marketdatav2.GetLatestStockBarSingleRequest
	GetLatestStockBarSingleResponse = marketdatav2.GetLatestStockBarSingleResponse
)

// =============================================================================
// Stock Trades Request/Response Types (v2)
// =============================================================================

type (
	// Multi-symbol requests.
	GetStockTradesRequest        = marketdatav2.GetStockTradesRequest
	GetStockTradesResponse       = marketdatav2.GetStockTradesResponse
	GetLatestStockTradesRequest  = marketdatav2.GetLatestStockTradesRequest
	GetLatestStockTradesResponse = marketdatav2.GetLatestStockTradesResponse

	// Single-symbol requests.
	GetStockTradesSingleRequest       = marketdatav2.GetStockTradesSingleRequest
	GetStockTradesSingleResponse      = marketdatav2.GetStockTradesSingleResponse
	GetLatestStockTradeSingleRequest  = marketdatav2.GetLatestStockTradeSingleRequest
	GetLatestStockTradeSingleResponse = marketdatav2.GetLatestStockTradeSingleResponse
)

// =============================================================================
// Stock Quotes Request/Response Types (v2)
// =============================================================================

type (
	// Multi-symbol requests.
	GetStockQuotesRequest        = marketdatav2.GetStockQuotesRequest
	GetStockQuotesResponse       = marketdatav2.GetStockQuotesResponse
	GetLatestStockQuotesRequest  = marketdatav2.GetLatestStockQuotesRequest
	GetLatestStockQuotesResponse = marketdatav2.GetLatestStockQuotesResponse

	// Single-symbol requests.
	GetStockQuotesSingleRequest       = marketdatav2.GetStockQuotesSingleRequest
	GetStockQuotesSingleResponse      = marketdatav2.GetStockQuotesSingleResponse
	GetLatestStockQuoteSingleRequest  = marketdatav2.GetLatestStockQuoteSingleRequest
	GetLatestStockQuoteSingleResponse = marketdatav2.GetLatestStockQuoteSingleResponse
)

// =============================================================================
// Stock Snapshots Request/Response Types (v2)
// =============================================================================

type (
	GetStockSnapshotsRequest  = marketdatav2.GetStockSnapshotsRequest
	GetStockSnapshotsResponse = marketdatav2.GetStockSnapshotsResponse
	GetStockSnapshotRequest   = marketdatav2.GetStockSnapshotRequest
)

// =============================================================================
// Stock Auctions Request/Response Types (v2)
// =============================================================================

type (
	// Multi-symbol requests.
	GetStockAuctionsRequest  = marketdatav2.GetStockAuctionsRequest
	GetStockAuctionsResponse = marketdatav2.GetStockAuctionsResponse

	// Single-symbol requests.
	GetStockAuctionsSingleRequest  = marketdatav2.GetStockAuctionsSingleRequest
	GetStockAuctionsSingleResponse = marketdatav2.GetStockAuctionsSingleResponse
)

// =============================================================================
// Stock Meta Request/Response Types (v2)
// =============================================================================

type (
	GetStockMetaConditionsRequest  = marketdatav2.GetStockMetaConditionsRequest
	GetStockMetaConditionsResponse = marketdatav2.GetStockMetaConditionsResponse
	GetStockMetaExchangesRequest   = marketdatav2.GetStockMetaExchangesRequest
	GetStockMetaExchangesResponse  = marketdatav2.GetStockMetaExchangesResponse
)

// =============================================================================
// Crypto Bars Request/Response Types (v1beta)
// =============================================================================

type (
	GetCryptoBarsRequest        = marketdatav1beta.GetCryptoBarsRequest
	GetCryptoBarsResponse       = marketdatav1beta.GetCryptoBarsResponse
	GetLatestCryptoBarsRequest  = marketdatav1beta.GetLatestCryptoBarsRequest
	GetLatestCryptoBarsResponse = marketdatav1beta.GetLatestCryptoBarsResponse
)

// =============================================================================
// Crypto Trades Request/Response Types (v1beta)
// =============================================================================

type (
	GetCryptoTradesRequest        = marketdatav1beta.GetCryptoTradesRequest
	GetCryptoTradesResponse       = marketdatav1beta.GetCryptoTradesResponse
	GetLatestCryptoTradesRequest  = marketdatav1beta.GetLatestCryptoTradesRequest
	GetLatestCryptoTradesResponse = marketdatav1beta.GetLatestCryptoTradesResponse
)

// =============================================================================
// Crypto Quotes Request/Response Types (v1beta)
// =============================================================================

type (
	GetCryptoQuotesRequest        = marketdatav1beta.GetCryptoQuotesRequest
	GetCryptoQuotesResponse       = marketdatav1beta.GetCryptoQuotesResponse
	GetLatestCryptoQuotesRequest  = marketdatav1beta.GetLatestCryptoQuotesRequest
	GetLatestCryptoQuotesResponse = marketdatav1beta.GetLatestCryptoQuotesResponse
)

// =============================================================================
// Crypto Snapshots Request/Response Types (v1beta)
// =============================================================================

type (
	GetCryptoSnapshotsRequest  = marketdatav1beta.GetCryptoSnapshotsRequest
	GetCryptoSnapshotsResponse = marketdatav1beta.GetCryptoSnapshotsResponse
)

// =============================================================================
// Crypto Orderbooks Request/Response Types (v1beta)
// =============================================================================

type (
	GetCryptoOrderbooksRequest  = marketdatav1beta.GetCryptoOrderbooksRequest
	GetCryptoOrderbooksResponse = marketdatav1beta.GetCryptoOrderbooksResponse
)

// =============================================================================
// Option Bars Request/Response Types (v1beta)
// =============================================================================

type (
	GetOptionBarsRequest  = marketdatav1beta.GetOptionBarsRequest
	GetOptionBarsResponse = marketdatav1beta.GetOptionBarsResponse
)

// =============================================================================
// Option Trades Request/Response Types (v1beta)
// =============================================================================

type (
	GetOptionTradesRequest        = marketdatav1beta.GetOptionTradesRequest
	GetOptionTradesResponse       = marketdatav1beta.GetOptionTradesResponse
	GetLatestOptionTradesRequest  = marketdatav1beta.GetLatestOptionTradesRequest
	GetLatestOptionTradesResponse = marketdatav1beta.GetLatestOptionTradesResponse
)

// =============================================================================
// Option Quotes Request/Response Types (v1beta)
// =============================================================================

type (
	GetLatestOptionQuotesRequest  = marketdatav1beta.GetLatestOptionQuotesRequest
	GetLatestOptionQuotesResponse = marketdatav1beta.GetLatestOptionQuotesResponse
)

// =============================================================================
// Option Snapshots Request/Response Types (v1beta)
// =============================================================================

type (
	GetOptionSnapshotsRequest  = marketdatav1beta.GetOptionSnapshotsRequest
	GetOptionSnapshotsResponse = marketdatav1beta.GetOptionSnapshotsResponse
	GetOptionChainRequest      = marketdatav1beta.GetOptionChainRequest
	GetOptionChainResponse     = marketdatav1beta.GetOptionChainResponse
)

// =============================================================================
// Option Meta Request/Response Types (v1beta)
// =============================================================================

type (
	GetOptionMetaConditionsRequest  = marketdatav1beta.GetOptionMetaConditionsRequest
	GetOptionMetaConditionsResponse = marketdatav1beta.GetOptionMetaConditionsResponse
	GetOptionMetaExchangesRequest   = marketdatav1beta.GetOptionMetaExchangesRequest
	GetOptionMetaExchangesResponse  = marketdatav1beta.GetOptionMetaExchangesResponse
)

// =============================================================================
// News Request/Response Types (v1beta)
// =============================================================================

type (
	GetNewsRequest  = marketdatav1beta.GetNewsRequest
	GetNewsResponse = marketdatav1beta.GetNewsResponse
)

// =============================================================================
// Screener Request/Response Types (v1beta)
// =============================================================================

type (
	GetMostActivesRequest  = marketdatav1beta.GetMostActivesRequest
	GetMostActivesResponse = marketdatav1beta.GetMostActivesResponse
	GetMoversRequest       = marketdatav1beta.GetMoversRequest
	GetMoversResponse      = marketdatav1beta.GetMoversResponse
)

// =============================================================================
// Corporate Actions Request/Response Types (v1beta)
// =============================================================================

type (
	GetCorporateActionsRequest  = marketdatav1beta.GetCorporateActionsRequest
	GetCorporateActionsResponse = marketdatav1beta.GetCorporateActionsResponse
)

// =============================================================================
// Forex Request/Response Types (v1beta)
// =============================================================================

type (
	GetForexRatesRequest        = marketdatav1beta.GetForexRatesRequest
	GetForexRatesResponse       = marketdatav1beta.GetForexRatesResponse
	GetLatestForexRatesRequest  = marketdatav1beta.GetLatestForexRatesRequest
	GetLatestForexRatesResponse = marketdatav1beta.GetLatestForexRatesResponse
)

// =============================================================================
// Fixed Income Request/Response Types (v1beta)
// =============================================================================

type (
	GetFixedIncomeLatestPricesRequest  = marketdatav1beta.GetFixedIncomeLatestPricesRequest
	GetFixedIncomeLatestPricesResponse = marketdatav1beta.GetFixedIncomeLatestPricesResponse
)

// =============================================================================
// Logos Request/Response Types (v1beta)
// =============================================================================

type (
	GetLogoRequest  = marketdatav1beta.GetLogoRequest
	GetLogoResponse = marketdatav1beta.GetLogoResponse
)

// =============================================================================
// String Constants (formerly enums)
// =============================================================================

// Timeframe string constants.
const (
	Timeframe_1MIN   = "1Min"
	Timeframe_5MIN   = "5Min"
	Timeframe_15MIN  = "15Min"
	Timeframe_30MIN  = "30Min"
	Timeframe_1HOUR  = "1Hour"
	Timeframe_4HOUR  = "4Hour"
	Timeframe_1DAY   = "1Day"
	Timeframe_1WEEK  = "1Week"
	Timeframe_1MONTH = "1Month"
)

// Feed string constants.
const (
	Feed_IEX = "iex"
	Feed_SIP = "sip"
	Feed_OTC = "otc"
)

// Adjustment string constants.
const (
	Adjustment_RAW      = "raw"
	Adjustment_SPLIT    = "split"
	Adjustment_DIVIDEND = "dividend"
	Adjustment_ALL      = "all"
)

// Sort string constants.
const (
	Sort_ASC  = "asc"
	Sort_DESC = "desc"
)

// CryptoLoc string constants.
const (
	CryptoLoc_US = "us"
)
