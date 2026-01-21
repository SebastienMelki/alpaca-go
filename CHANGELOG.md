# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-01-21

### Added

#### Trading API (`pkg/trading`)
- Account management: `GetAccount`, `GetAccountConfigurations`, `UpdateAccountConfigurations`
- Portfolio: `GetPortfolioHistory`, `GetAccountActivities`
- Orders: `CreateOrder`, `ListOrders`, `GetOrder`, `GetOrderByClientId`, `ReplaceOrder`, `CancelOrder`, `CancelAllOrders`
- Positions: `ListPositions`, `GetPosition`, `ClosePosition`, `CloseAllPositions`, `ExerciseOption`
- Assets: `ListAssets`, `GetAsset`
- Market info: `GetClock`, `GetCalendar`
- Watchlists: Full CRUD operations with `ListWatchlists`, `CreateWatchlist`, `GetWatchlist`, `UpdateWatchlist`, `DeleteWatchlist`, `AddWatchlistAsset`, `RemoveWatchlistAsset`
- Paper trading support via `NewPaperClient()`

#### Market Data API (`pkg/marketdata`)
- Stock data: bars, trades, quotes, snapshots, auctions
- Crypto data: bars, trades, quotes, snapshots
- Options data: bars, trades, quotes, snapshots, option chains
- News: `GetNews`
- Screener: `GetMostActives`, `GetMovers`

#### Broker API (`pkg/broker`)
- Account management: `CreateAccount`, `ListAccounts`, `GetAccount`, `UpdateAccount`, `CloseAccount`
- ACH relationships: `CreateACHRelationship`, `ListACHRelationships`, `DeleteACHRelationship`
- Transfers: `CreateTransfer`, `ListTransfers`, `GetTransfer`, `CancelTransfer`
- Trading for managed accounts: orders and positions management
- Sandbox environment support via `NewSandboxClient()`

#### Infrastructure
- Protocol Buffer definitions for all three Alpaca APIs
- Generated HTTP clients using `protoc-gen-go-client`
- Generated OpenAPI 3.1 specifications for each service
- Makefile with `generate`, `build`, `lint`, `lint-fix`, `buf-lint`, `check`, and `release` targets
- golangci-lint configuration for code quality

[1.0.0]: https://github.com/sebastienmelki/alpaca-go/releases/tag/v1.0.0
