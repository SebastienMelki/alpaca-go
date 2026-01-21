package marketdata

import (
	"net/http"

	marketdatav2 "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/marketdata/v2"
)

const (
	BaseURL = "https://data.alpaca.markets"
)

// Client wraps the generated MarketDataService client with Alpaca-specific defaults.
type Client struct {
	marketdatav2.MarketDataServiceClient
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

// NewClient creates a new Market Data API client.
func NewClient(apiKey, apiSecret string, opts ...Option) *Client {
	cfg := &options{
		httpClient: http.DefaultClient,
		baseURL:    BaseURL,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	client := marketdatav2.NewMarketDataServiceClient(
		cfg.baseURL,
		marketdatav2.WithMarketDataServiceHTTPClient(cfg.httpClient),
		marketdatav2.WithMarketDataServiceAPCAAPIKEYID(apiKey),
		marketdatav2.WithMarketDataServiceAPCAAPISECRETKEY(apiSecret),
	)

	return &Client{client}
}

// Re-export common types.
type (
	Bar         = marketdatav2.Bar
	Trade       = marketdatav2.Trade
	Quote       = marketdatav2.Quote
	Snapshot    = marketdatav2.Snapshot
	NewsArticle = marketdatav2.NewsArticle
	Auction     = marketdatav2.Auction
)

// Re-export crypto types.
type (
	CryptoBar      = marketdatav2.CryptoBar
	CryptoTrade    = marketdatav2.CryptoTrade
	CryptoQuote    = marketdatav2.CryptoQuote
	CryptoSnapshot = marketdatav2.CryptoSnapshot
)

// Re-export option types.
type (
	OptionBar      = marketdatav2.OptionBar
	OptionTrade    = marketdatav2.OptionTrade
	OptionQuote    = marketdatav2.OptionQuote
	OptionSnapshot = marketdatav2.OptionSnapshot
)

// Re-export screener types.
type (
	MostActive = marketdatav2.MostActive
	Mover      = marketdatav2.Mover
)

// Re-export stock request types.
type (
	GetStockBarsRequest         = marketdatav2.GetStockBarsRequest
	GetLatestStockBarsRequest   = marketdatav2.GetLatestStockBarsRequest
	GetStockTradesRequest       = marketdatav2.GetStockTradesRequest
	GetLatestStockTradesRequest = marketdatav2.GetLatestStockTradesRequest
	GetStockQuotesRequest       = marketdatav2.GetStockQuotesRequest
	GetLatestStockQuotesRequest = marketdatav2.GetLatestStockQuotesRequest
	GetStockSnapshotsRequest    = marketdatav2.GetStockSnapshotsRequest
	GetStockSnapshotRequest     = marketdatav2.GetStockSnapshotRequest
	GetStockAuctionsRequest     = marketdatav2.GetStockAuctionsRequest
)

// Re-export stock response types.
type (
	GetStockBarsResponse         = marketdatav2.GetStockBarsResponse
	GetLatestStockBarsResponse   = marketdatav2.GetLatestStockBarsResponse
	GetStockTradesResponse       = marketdatav2.GetStockTradesResponse
	GetLatestStockTradesResponse = marketdatav2.GetLatestStockTradesResponse
	GetStockQuotesResponse       = marketdatav2.GetStockQuotesResponse
	GetLatestStockQuotesResponse = marketdatav2.GetLatestStockQuotesResponse
	GetStockSnapshotsResponse    = marketdatav2.GetStockSnapshotsResponse
	GetStockAuctionsResponse     = marketdatav2.GetStockAuctionsResponse
)

// Re-export crypto request types.
type (
	GetCryptoBarsRequest         = marketdatav2.GetCryptoBarsRequest
	GetLatestCryptoBarsRequest   = marketdatav2.GetLatestCryptoBarsRequest
	GetCryptoTradesRequest       = marketdatav2.GetCryptoTradesRequest
	GetLatestCryptoTradesRequest = marketdatav2.GetLatestCryptoTradesRequest
	GetCryptoQuotesRequest       = marketdatav2.GetCryptoQuotesRequest
	GetLatestCryptoQuotesRequest = marketdatav2.GetLatestCryptoQuotesRequest
	GetCryptoSnapshotsRequest    = marketdatav2.GetCryptoSnapshotsRequest
)

// Re-export crypto response types.
type (
	GetCryptoBarsResponse         = marketdatav2.GetCryptoBarsResponse
	GetLatestCryptoBarsResponse   = marketdatav2.GetLatestCryptoBarsResponse
	GetCryptoTradesResponse       = marketdatav2.GetCryptoTradesResponse
	GetLatestCryptoTradesResponse = marketdatav2.GetLatestCryptoTradesResponse
	GetCryptoQuotesResponse       = marketdatav2.GetCryptoQuotesResponse
	GetLatestCryptoQuotesResponse = marketdatav2.GetLatestCryptoQuotesResponse
	GetCryptoSnapshotsResponse    = marketdatav2.GetCryptoSnapshotsResponse
)

// Re-export option request types.
type (
	GetOptionBarsRequest         = marketdatav2.GetOptionBarsRequest
	GetLatestOptionBarsRequest   = marketdatav2.GetLatestOptionBarsRequest
	GetOptionTradesRequest       = marketdatav2.GetOptionTradesRequest
	GetLatestOptionTradesRequest = marketdatav2.GetLatestOptionTradesRequest
	GetOptionQuotesRequest       = marketdatav2.GetOptionQuotesRequest
	GetLatestOptionQuotesRequest = marketdatav2.GetLatestOptionQuotesRequest
	GetOptionSnapshotsRequest    = marketdatav2.GetOptionSnapshotsRequest
	GetOptionChainRequest        = marketdatav2.GetOptionChainRequest
)

// Re-export option response types.
type (
	GetOptionBarsResponse         = marketdatav2.GetOptionBarsResponse
	GetLatestOptionBarsResponse   = marketdatav2.GetLatestOptionBarsResponse
	GetOptionTradesResponse       = marketdatav2.GetOptionTradesResponse
	GetLatestOptionTradesResponse = marketdatav2.GetLatestOptionTradesResponse
	GetOptionQuotesResponse       = marketdatav2.GetOptionQuotesResponse
	GetLatestOptionQuotesResponse = marketdatav2.GetLatestOptionQuotesResponse
	GetOptionSnapshotsResponse    = marketdatav2.GetOptionSnapshotsResponse
	GetOptionChainResponse        = marketdatav2.GetOptionChainResponse
)

// Re-export news request/response types.
type (
	GetNewsRequest  = marketdatav2.GetNewsRequest
	GetNewsResponse = marketdatav2.GetNewsResponse
)

// Re-export screener request/response types.
type (
	GetMostActivesRequest  = marketdatav2.GetMostActivesRequest
	GetMostActivesResponse = marketdatav2.GetMostActivesResponse
	GetMoversRequest       = marketdatav2.GetMoversRequest
	GetMoversResponse      = marketdatav2.GetMoversResponse
)

// Re-export enums.
type (
	Timeframe  = marketdatav2.Timeframe
	Adjustment = marketdatav2.Adjustment
	Feed       = marketdatav2.Feed
	Sort       = marketdatav2.Sort
	CryptoLoc  = marketdatav2.CryptoLoc
)

// Enum value constants for Timeframe.
const (
	Timeframe_TIMEFRAME_UNSPECIFIED = marketdatav2.Timeframe_TIMEFRAME_UNSPECIFIED
	Timeframe_TIMEFRAME_1MIN        = marketdatav2.Timeframe_TIMEFRAME_1MIN
	Timeframe_TIMEFRAME_5MIN        = marketdatav2.Timeframe_TIMEFRAME_5MIN
	Timeframe_TIMEFRAME_15MIN       = marketdatav2.Timeframe_TIMEFRAME_15MIN
	Timeframe_TIMEFRAME_30MIN       = marketdatav2.Timeframe_TIMEFRAME_30MIN
	Timeframe_TIMEFRAME_1HOUR       = marketdatav2.Timeframe_TIMEFRAME_1HOUR
	Timeframe_TIMEFRAME_4HOUR       = marketdatav2.Timeframe_TIMEFRAME_4HOUR
	Timeframe_TIMEFRAME_1DAY        = marketdatav2.Timeframe_TIMEFRAME_1DAY
	Timeframe_TIMEFRAME_1WEEK       = marketdatav2.Timeframe_TIMEFRAME_1WEEK
	Timeframe_TIMEFRAME_1MONTH      = marketdatav2.Timeframe_TIMEFRAME_1MONTH
)

// Enum value constants for Feed.
const (
	Feed_FEED_UNSPECIFIED = marketdatav2.Feed_FEED_UNSPECIFIED
	Feed_FEED_IEX         = marketdatav2.Feed_FEED_IEX
	Feed_FEED_SIP         = marketdatav2.Feed_FEED_SIP
)

// Enum value constants for Adjustment.
const (
	Adjustment_ADJUSTMENT_UNSPECIFIED = marketdatav2.Adjustment_ADJUSTMENT_UNSPECIFIED
	Adjustment_ADJUSTMENT_RAW         = marketdatav2.Adjustment_ADJUSTMENT_RAW
	Adjustment_ADJUSTMENT_SPLIT       = marketdatav2.Adjustment_ADJUSTMENT_SPLIT
	Adjustment_ADJUSTMENT_DIVIDEND    = marketdatav2.Adjustment_ADJUSTMENT_DIVIDEND
	Adjustment_ADJUSTMENT_ALL         = marketdatav2.Adjustment_ADJUSTMENT_ALL
)
