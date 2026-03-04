# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go SDK for the Alpaca Trading API, built using [sebuf](https://github.com/sebastienmelki/sebuf) to generate HTTP clients and OpenAPI documentation from Protocol Buffer definitions.

## Build Commands

```bash
# Install all required tools (code generators + linter)
make install-tools

# Generate Go code from proto files
make generate

# Build
make build

# Run tests
go test ./...

# Run a single test
go test -run TestName ./path/to/package

# Run linter
make lint

# Run linter with auto-fix
make lint-fix

# Lint proto files
make buf-lint

# Run all checks (buf-lint, lint, generate, build, test)
make check
```

## Architecture

### Protobuf-First Design

All API definitions live in `.proto` files. The sebuf generators produce:
- HTTP client code with automatic JSON serialization
- Request validation using `buf.validate` annotations
- OpenAPI 3.1 documentation per service

### Alpaca API Coverage

The SDK covers these Alpaca APIs:
- **Trading API**: Orders, positions, account management, crypto funding, option contracts (`api.alpaca.markets`)
- **Market Data API**: Real-time and historical quotes/bars for stocks (v2), crypto/options/news (v1beta) (`data.alpaca.markets`)
- **Broker API**: Multi-account management with v1 (core), v1beta (funding wallets), v2 (SSE events)
- **Auth API**: OAuth2 token issuance (`api.alpaca.markets`)
- **Paper Trading**: Same endpoints at `paper-api.alpaca.markets`

### Authentication Pattern

Alpaca uses header-based API key authentication:
```
APCA-API-KEY-ID: <key_id>
APCA-API-SECRET-KEY: <secret_key>
```

Define these as required service headers in proto files using `sebuf.http.service_headers`.

### Proto File Organization

```
alpaca/
├── auth/v1/              # Auth/OAuth API
│   ├── service.proto     # AuthService definition
│   └── token.proto       # Token request/response
├── trading/v1/           # Trading API (orders, positions, account)
│   ├── service.proto     # TradingService definition
│   ├── account.proto     # Account model
│   ├── order.proto       # Order model + enums
│   ├── position.proto    # Position model
│   ├── option_contract.proto  # Option contracts
│   ├── crypto_funding.proto   # Crypto wallets/transfers
│   └── *.proto           # Request/response per operation
├── marketdata/
│   ├── v2/               # Market Data v2 (stable - stocks only)
│   │   ├── service.proto
│   │   └── stock_*.proto
│   └── v1beta/           # Market Data v1beta (beta - crypto, options, news)
│       ├── service.proto
│       ├── crypto_*.proto
│       ├── option_*.proto
│       └── news.proto
└── broker/
    ├── v1/               # Broker v1 (core operations)
    │   ├── service.proto # 80+ RPC methods
    │   ├── broker_account.proto
    │   ├── journal.proto
    │   ├── rebalancing.proto
    │   └── *.proto
    ├── v1beta/           # Broker v1beta (funding wallets)
    │   └── service.proto
    └── v2/               # Broker v2 (SSE events)
        └── service.proto
```

### Generated Code Structure

```
internal/gen/             # Generated Go clients (private)
pkg/                      # Public client wrappers
├── auth/                 # Auth API client
├── trading/              # Trading API client
├── marketdata/           # Market Data API client (V2 + V1Beta)
└── broker/               # Broker API client (V1 + V1Beta + V2)
cmd/apitest/              # API test harness for validation
docs/                     # Generated OpenAPI 3.1 specs per service
curl-verified.md          # SDK methods verified against raw curl responses
```

## Key Patterns

### Validation Annotations

Use `buf.validate` for request validation:
```protobuf
message CreateOrderRequest {
  string symbol = 1 [(buf.validate.field).string = {min_len: 1, max_len: 10}];
  string qty = 2 [(buf.validate.field).string.pattern = "^[0-9]+(\\.[0-9]+)?$"];
}
```

### Error Messages

Name error response messages with `Error` suffix to auto-implement Go's `error` interface:
```protobuf
message OrderRejectedError {
  string code = 1;
  string message = 2;
}
```

### HTTP Methods and Paths

Define REST endpoints with HTTP methods:
```protobuf
rpc GetAccount(GetAccountRequest) returns (Account) {
  option (sebuf.http.config) = {
    path: "/{account_id}"
    method: HTTP_METHOD_GET
  };
}

rpc CreateOrder(CreateOrderRequest) returns (Order) {
  option (sebuf.http.config) = {
    path: "/{account_id}/orders"
    method: HTTP_METHOD_POST
  };
}

rpc CancelOrder(CancelOrderRequest) returns (CancelOrderResponse) {
  option (sebuf.http.config) = {
    path: "/{account_id}/orders/{order_id}"
    method: HTTP_METHOD_DELETE
  };
}
```

### Query Parameters

Define query parameters with `sebuf.http.query`:
```protobuf
message ListOrdersRequest {
  string status = 2 [(sebuf.http.query) = { name: "status" }];
  int32 limit = 3 [(sebuf.http.query) = { name: "limit" }];
}
```

### Service Headers

Define API authentication at the service level:
```protobuf
service TradingService {
  option (sebuf.http.service_headers) = {
    required_headers: [
      {name: "APCA-API-KEY-ID" type: "string"},
      {name: "APCA-API-SECRET-KEY" type: "string"}
    ]
  };
}
```

## Alpaca API Base URLs

- Live Trading: `https://api.alpaca.markets`
- Paper Trading: `https://paper-api.alpaca.markets`
- Market Data: `https://data.alpaca.markets`
- Broker API: `https://broker-api.alpaca.markets`
- Broker Sandbox: `https://broker-api.sandbox.alpaca.markets`
- Auth API: `https://api.alpaca.markets`

## API Test Harness

The `cmd/apitest` directory contains a test harness for validating generated clients:

```bash
# Test against live APIs (requires .env with credentials)
make test-marketdata
make test-broker
make test-trading
```

Environment variables: `ALPACA_API_KEY`, `ALPACA_API_SECRET`, `TEST_ACCOUNT_ID`
