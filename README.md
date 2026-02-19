# alpaca-go

A Go SDK for the [Alpaca Trading API](https://alpaca.markets/docs/api-references/), built using Protocol Buffers and [sebuf](https://github.com/sebastienmelki/sebuf) for type-safe HTTP client generation.

## 30-Second Quick Start

```go
package main

import (
    "context"
    "fmt"
    "github.com/sebastienmelki/alpaca-go/pkg/trading"
)

func main() {
    client := trading.NewPaperClient("YOUR_API_KEY", "YOUR_API_SECRET")
    account, _ := client.GetAccount(context.Background(), &trading.GetAccountRequest{})
    fmt.Printf("Buying Power: %s\n", account.BuyingPower)
}
```

```bash
go get github.com/sebastienmelki/alpaca-go
```

## Features

- **Type-safe API clients** generated from Protocol Buffer definitions
- **Full Alpaca API coverage**: Trading, Market Data, Broker, and Auth APIs
- **Paper trading support** with dedicated client constructors
- **Multi-version support**: v1, v1beta, and v2 APIs where applicable
- **Automatic JSON serialization** with protojson
- **OpenAPI 3.1 documentation** generated from proto files

## Installation

```bash
go get github.com/sebastienmelki/alpaca-go
```

## Quick Start

### Trading API

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/sebastienmelki/alpaca-go/pkg/trading"
)

func main() {
    // Create a client for live trading
    client := trading.NewClient("YOUR_API_KEY", "YOUR_API_SECRET")

    // Or use paper trading
    // client := trading.NewPaperClient("YOUR_API_KEY", "YOUR_API_SECRET")

    ctx := context.Background()

    // Get account information
    account, err := client.GetAccount(ctx, &trading.GetAccountRequest{})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Account: %s, Buying Power: %s\n", account.AccountNumber, account.BuyingPower)

    // Place a market order
    order, err := client.CreateOrder(ctx, &trading.CreateOrderRequest{
        Symbol:      "AAPL",
        Qty:         "1",
        Side:        trading.OrderSide_ORDER_SIDE_BUY,
        Type:        trading.OrderType_ORDER_TYPE_MARKET,
        TimeInForce: trading.TimeInForce_TIME_IN_FORCE_DAY,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Order placed: %s\n", order.Id)

    // List positions
    positions, err := client.ListPositions(ctx, &trading.ListPositionsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    for _, pos := range positions.Positions {
        fmt.Printf("Position: %s, Qty: %s\n", pos.Symbol, pos.Qty)
    }
}
```

### Market Data API

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/sebastienmelki/alpaca-go/pkg/marketdata"
)

func main() {
    client := marketdata.NewClient("YOUR_API_KEY", "YOUR_API_SECRET")

    // Or use the sandbox environment (for broker sandbox)
    // client := marketdata.NewSandboxClient("YOUR_API_KEY", "YOUR_API_SECRET")

    ctx := context.Background()

    // Get latest stock bars (V2 API - stable)
    bars, err := client.V2.GetLatestStockBars(ctx, &marketdata.GetLatestStockBarsRequest{
        Symbols: "AAPL,MSFT,GOOGL",
    })
    if err != nil {
        log.Fatal(err)
    }
    for symbol, bar := range bars.Bars {
        fmt.Printf("%s: Close=%s, Volume=%d\n", symbol, bar.Close, bar.Volume)
    }

    // Get crypto quotes (V1Beta API - beta features)
    cryptoQuotes, err := client.V1Beta.GetLatestCryptoQuotes(ctx, &marketdata.GetLatestCryptoQuotesRequest{
        Symbols: "BTC/USD",
    })
    if err != nil {
        log.Fatal(err)
    }
    for symbol, quote := range cryptoQuotes.Quotes {
        fmt.Printf("%s: Bid=%s, Ask=%s\n", symbol, quote.BidPrice, quote.AskPrice)
    }
}
```

### Broker API

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/sebastienmelki/alpaca-go/pkg/broker"
)

func main() {
    // Broker API uses Basic Auth
    client := broker.NewClient("YOUR_API_KEY", "YOUR_API_SECRET")

    // Or use sandbox environment
    // client := broker.NewSandboxClient("YOUR_API_KEY", "YOUR_API_SECRET")

    ctx := context.Background()

    // List accounts (V1 API)
    accounts, err := client.V1.ListAccounts(ctx, &broker.ListAccountsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    for _, acc := range accounts.Accounts {
        fmt.Printf("Account: %s, Status: %s\n", acc.Id, acc.Status)
    }

    // Access V1Beta features (funding wallets)
    // wallet, err := client.V1Beta.GetFundingWallet(ctx, &broker.GetFundingWalletRequest{...})

    // Access V2 features (SSE events)
    // events, err := client.V2.SubscribeTradeEventsV2(ctx, &broker.SubscribeTradeEventsV2Request{...})
}
```

## API Coverage

### Trading API (`pkg/trading`)

| Category | Endpoints |
|----------|-----------|
| Account | GetAccount, GetAccountConfigurations, UpdateAccountConfigurations, GetPortfolioHistory, GetAccountActivities |
| Orders | CreateOrder, ListOrders, GetOrder, GetOrderByClientId, ReplaceOrder, CancelOrder, CancelAllOrders |
| Positions | ListPositions, GetPosition, ClosePosition, CloseAllPositions, ExerciseOption |
| Assets | ListAssets, GetAsset |
| Options | ListOptionContracts, GetOptionContract |
| Market | GetClock, GetCalendar |
| Watchlists | ListWatchlists, CreateWatchlist, GetWatchlist, UpdateWatchlist, DeleteWatchlist, AddWatchlistAsset, RemoveWatchlistAsset |
| Crypto Funding | ListCryptoWallets, GetCryptoWallet, ListCryptoTransfers, CreateCryptoWithdrawal |

### Market Data API (`pkg/marketdata`)

**V2 (Stable)** - Stock data endpoints:

| Category | Endpoints |
|----------|-----------|
| Stocks | GetStockBars, GetLatestStockBars, GetStockTrades, GetLatestStockTrades, GetStockQuotes, GetLatestStockQuotes, GetStockSnapshots, GetStockSnapshot, GetStockAuctions |
| Metadata | GetStockConditions, GetExchangeCodes |

**V1Beta (Beta)** - Crypto, options, and additional features:

| Category | Endpoints |
|----------|-----------|
| Crypto | GetCryptoBars, GetLatestCryptoBars, GetCryptoTrades, GetLatestCryptoTrades, GetCryptoQuotes, GetLatestCryptoQuotes, GetCryptoSnapshots, GetCryptoOrderbooks |
| Options | GetOptionBars, GetOptionTrades, GetOptionQuotes, GetOptionSnapshots, GetOptionChain, GetOptionMeta |
| News | GetNews |
| Screener | GetMostActives, GetMovers |
| Corporate Actions | GetCorporateActions |
| Forex | GetForexRates, GetLatestForexRates |
| Logos | GetLogo |

### Broker API (`pkg/broker`)

**V1 (Core)** - Main broker operations:

| Category | Endpoints |
|----------|-----------|
| Accounts | CreateAccount, ListAccounts, GetAccount, UpdateAccount, CloseAccount |
| ACH | CreateACHRelationship, ListACHRelationships, DeleteACHRelationship |
| Transfers | CreateTransfer, ListTransfers, GetTransfer, CancelTransfer |
| Trading | CreateTradingOrder, ListTradingOrders, GetTradingOrder, CancelTradingOrder |
| Positions | ListTradingPositions, GetTradingPosition, CloseTradingPosition, CloseAllTradingPositions |
| Documents | ListAccountDocuments, DownloadAccountDocument, DownloadW8BenDocument |
| Watchlists | ListBrokerWatchlists, CreateBrokerWatchlist, GetBrokerWatchlist, UpdateBrokerWatchlist, DeleteBrokerWatchlist |
| Journals | CreateJournal, ListJournals, GetJournal, DeleteJournal, CreateBatchJournal, ReverseBatchJournal |
| Rebalancing | ListPortfolios, CreatePortfolio, GetPortfolio, UpdatePortfolio, DeletePortfolio, ListRuns, CreateRun |
| CIP/KYC | GetCIPInfo, UpdateCIPInfo |
| Onfido | CreateOnfidoApplicant, CreateOnfidoCheck, GenerateOnfidoSDKToken |
| OAuth | CreateOAuthToken, AuthorizeOAuth, CreateOAuthClient, UpdateOAuthClient |
| Options | GetOptionsApproval, RequestOptionsApproval, UpdateOptionsApproval, ListBrokerOptionContracts |
| Crypto | ListBrokerCryptoWallets, ListBrokerCryptoTransfers, CreateBrokerCryptoTransfer |
| FPSL | ListFPSLLoans, ListFPSLTiers, ListAPRTiers |
| JIT | ListJITLedgers, GetJITLedgerBalances, GetJITLimits, CreateJITSettlement |
| IRA | ListIRAExcessContributions |
| Calendar | GetMarketCalendar, GetMarketClock |

**V1Beta** - Beta funding features:

| Category | Endpoints |
|----------|-----------|
| Funding Wallets | GetFundingWallet, CreateFundingWallet, ListFundingWalletTransfers, CreateFundingWalletWithdrawal |
| Recipient Banks | GetRecipientBank, CreateRecipientBank, DeleteRecipientBank |

**V2** - SSE event streaming:

| Category | Endpoints |
|----------|-----------|
| Events | SubscribeTradeEventsV2, SubscribeJournalEventsV2, SubscribeSystemEventsV2, SubscribeAdminActionsV2, SubscribeFundingStatusV2 |

### Auth API (`pkg/auth`)

| Category | Endpoints |
|----------|-----------|
| OAuth2 | IssueToken |

### Pagination

Paginated endpoints return a `NextPageToken` field. Pass it back as `PageToken` to fetch the next page. An empty `NextPageToken` means you've reached the last page.

```go
ctx := context.Background()
client := marketdata.NewClient("YOUR_API_KEY", "YOUR_API_SECRET")

var pageToken string
for {
    resp, err := client.V2.GetStockBars(ctx, &marketdata.GetStockBarsRequest{
        Symbols:   "AAPL",
        Timeframe: "1Day",
        Start:     "2024-01-01",
        End:       "2024-12-31",
        PageToken: pageToken,
    })
    if err != nil {
        log.Fatal(err)
    }

    // process resp.Bars ...

    if resp.NextPageToken == "" {
        break
    }
    pageToken = resp.NextPageToken
}
```

The same pattern applies to all paginated endpoints across the Market Data, Broker, and Trading APIs.

## Client Options

All clients support functional options for customization:

```go
import (
    "net/http"
    "time"

    "github.com/sebastienmelki/alpaca-go/pkg/trading"
)

// Custom HTTP client with timeout
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

client := trading.NewClient(
    "YOUR_API_KEY",
    "YOUR_API_SECRET",
    trading.WithHTTPClient(httpClient),
    trading.WithBaseURL("https://custom-proxy.example.com"),
)
```

## Base URLs

| API | Live | Paper/Sandbox |
|-----|------|---------------|
| Trading | `https://api.alpaca.markets` | `https://paper-api.alpaca.markets` |
| Market Data | `https://data.alpaca.markets` | `https://data.sandbox.alpaca.markets` |
| Broker | `https://broker-api.alpaca.markets` | `https://broker-api.sandbox.alpaca.markets` |
| Auth | `https://api.alpaca.markets` | - |

## Project Structure

```
alpaca-go/
├── alpaca/                  # Protocol Buffer definitions
│   ├── auth/v1/            # Auth/OAuth API protos
│   ├── trading/v1/         # Trading API protos
│   ├── marketdata/
│   │   ├── v2/             # Market Data v2 (stocks - stable)
│   │   └── v1beta/         # Market Data v1beta (crypto, options, news - beta)
│   └── broker/
│       ├── v1/             # Broker v1 (core operations)
│       ├── v1beta/         # Broker v1beta (funding wallets)
│       └── v2/             # Broker v2 (SSE events)
├── internal/gen/           # Generated code (private)
├── pkg/                    # Public client wrappers
│   ├── auth/              # Auth API client
│   ├── trading/           # Trading API client
│   ├── marketdata/        # Market Data API client (V2 + V1Beta)
│   └── broker/            # Broker API client (V1 + V1Beta + V2)
├── cmd/apitest/           # API test harness
└── docs/                   # Generated OpenAPI specs
```

## API Testing Tool

The `cmd/apitest` directory contains a test harness for validating generated API clients and showcasing SDK usage. This is primarily used to verify that proto files generate proper Go code.

```bash
# Test Market Data APIs
make test-marketdata

# Test Broker APIs
make test-broker

# Test Trading APIs
make test-trading
```

Set your credentials in a `.env` file or environment variables:

```bash
ALPACA_API_KEY=your_api_key
ALPACA_API_SECRET=your_api_secret
TEST_ACCOUNT_ID=your_test_account_id  # For broker tests
```

## Development

### Prerequisites

- Go 1.21+
- [buf](https://buf.build/docs/installation)
- Protocol Buffer compiler

### Build Commands

```bash
# Install code generators
make install-tools

# Generate code from proto files
make generate

# Build
make build

# Run linter
make lint

# Run linter with auto-fix
make lint-fix

# Run tests
go test ./...

# Run all checks
make check
```

### Regenerating Code

After modifying `.proto` files:

```bash
buf generate
```

This regenerates:
- Go types and HTTP clients in `internal/gen/`
- OpenAPI 3.1 specs in `docs/`

## Authentication

### Trading & Market Data APIs

Use API key headers:
- `APCA-API-KEY-ID`: Your API key ID
- `APCA-API-SECRET-KEY`: Your API secret key

The SDK handles this automatically when you create a client.

### Broker API

Uses HTTP Basic Auth with base64-encoded credentials. The SDK handles this automatically.

## Error Handling

The SDK returns structured errors from the Alpaca API:

```go
order, err := client.CreateOrder(ctx, &trading.CreateOrderRequest{...})
if err != nil {
    // Error includes status code and message from Alpaca
    log.Printf("Failed to create order: %v", err)
}
```

## License

MIT License - see LICENSE file for details.

## Links

- [Alpaca API Documentation](https://alpaca.markets/docs/api-references/)
- [sebuf - Protocol Buffer HTTP generator](https://github.com/sebastienmelki/sebuf)
