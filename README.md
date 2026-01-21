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
- **Full Alpaca API coverage**: Trading, Market Data, and Broker APIs
- **Paper trading support** with dedicated client constructors
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
    ctx := context.Background()

    // Get latest stock bars
    bars, err := client.GetLatestStockBars(ctx, &marketdata.GetLatestStockBarsRequest{
        Symbols: "AAPL,MSFT,GOOGL",
    })
    if err != nil {
        log.Fatal(err)
    }
    for symbol, bar := range bars.Bars {
        fmt.Printf("%s: Close=%s, Volume=%d\n", symbol, bar.Close, bar.Volume)
    }

    // Get stock quotes
    quotes, err := client.GetLatestStockQuotes(ctx, &marketdata.GetLatestStockQuotesRequest{
        Symbols: "AAPL",
    })
    if err != nil {
        log.Fatal(err)
    }
    for symbol, quote := range quotes.Quotes {
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

    // List accounts
    accounts, err := client.ListAccounts(ctx, &broker.ListAccountsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    for _, acc := range accounts.Accounts {
        fmt.Printf("Account: %s, Status: %s\n", acc.Id, acc.Status)
    }
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
| Market | GetClock, GetCalendar |
| Watchlists | ListWatchlists, CreateWatchlist, GetWatchlist, UpdateWatchlist, DeleteWatchlist, AddWatchlistAsset, RemoveWatchlistAsset |

### Market Data API (`pkg/marketdata`)

| Category | Endpoints |
|----------|-----------|
| Stocks | GetStockBars, GetLatestStockBars, GetStockTrades, GetLatestStockTrades, GetStockQuotes, GetLatestStockQuotes, GetStockSnapshots, GetStockSnapshot, GetStockAuctions |
| Crypto | GetCryptoBars, GetLatestCryptoBars, GetCryptoTrades, GetLatestCryptoTrades, GetCryptoQuotes, GetLatestCryptoQuotes, GetCryptoSnapshots |
| Options | GetOptionBars, GetLatestOptionBars, GetOptionTrades, GetLatestOptionTrades, GetOptionQuotes, GetLatestOptionQuotes, GetOptionSnapshots, GetOptionChain |
| News | GetNews |
| Screener | GetMostActives, GetMovers |

### Broker API (`pkg/broker`)

| Category | Endpoints |
|----------|-----------|
| Accounts | CreateAccount, ListAccounts, GetAccount, UpdateAccount, CloseAccount |
| ACH | CreateACHRelationship, ListACHRelationships, DeleteACHRelationship |
| Transfers | CreateTransfer, ListTransfers, GetTransfer, CancelTransfer |
| Trading | CreateTradingOrder, ListTradingOrders, GetTradingOrder, CancelTradingOrder |
| Positions | ListTradingPositions, GetTradingPosition, CloseTradingPosition, CloseAllTradingPositions |

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
| Market Data | `https://data.alpaca.markets` | - |
| Broker | `https://broker-api.alpaca.markets` | `https://broker-api.sandbox.alpaca.markets` |

## Project Structure

```
alpaca-go/
├── alpaca/                  # Protocol Buffer definitions
│   ├── core/v1/            # Shared types
│   ├── trading/v1/         # Trading API protos
│   ├── marketdata/v2/      # Market Data API protos
│   └── broker/v1/          # Broker API protos
├── internal/gen/           # Generated code (private)
├── pkg/                    # Public client wrappers
│   ├── trading/           # Trading API client
│   ├── marketdata/        # Market Data API client
│   └── broker/            # Broker API client
└── docs/                   # Generated OpenAPI specs
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
