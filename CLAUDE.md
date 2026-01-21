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
- **Trading API**: Orders, positions, account management (`api.alpaca.markets`)
- **Market Data API**: Real-time and historical quotes/bars (`data.alpaca.markets`)
- **Broker API**: Multi-account management for B2B use cases
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
├── core/v1/              # Shared types (identifiers)
│   └── identifiers.proto
├── trading/v1/           # Trading API (orders, positions, account)
│   ├── service.proto     # TradingService definition
│   ├── account.proto     # Account model
│   ├── order.proto       # Order model + enums
│   ├── position.proto    # Position model
│   └── *.proto           # Request/response per operation
├── marketdata/v2/        # Market Data API (stocks, crypto, options)
└── broker/v1/            # Broker API (multi-account management)
```

### Generated Code Structure

```
internal/gen/             # Generated Go clients (private)
pkg/                      # Public client wrappers
├── trading/              # Trading API client
├── marketdata/           # Market Data API client
└── broker/               # Broker API client
docs/                     # Generated OpenAPI 3.1 specs per service
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
