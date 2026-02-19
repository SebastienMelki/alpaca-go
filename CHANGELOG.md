# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.2.1] - 2026-02-19

### Fixed
- Pagination broken across all paginated endpoints: added `json_name = "next_page_token"` to all 21 `next_page_token` fields in proto definitions so the generated `UnmarshalJSON` methods look up the correct snake_case key (`"next_page_token"`) instead of the camelCase key (`"nextPageToken"`) that Alpaca's API never returns

## [1.2.0] - 2026-02-09

### Added

#### Market Data API (`pkg/marketdata`)
- Sandbox environment support via `NewSandboxClient()` (`https://data.sandbox.alpaca.markets`)

## [1.1.0] - 2026-02-05

### Added

#### Auth API (`pkg/auth`)
- New `IssueToken` endpoint for OAuth2 client credentials flow

#### Trading API (`pkg/trading`)
- Option contract endpoints: `ListOptionContracts`, `GetOptionContract`
- Crypto funding endpoints: `ListCryptoWallets`, `GetCryptoWallet`, `ListCryptoTransfers`, `CreateCryptoWithdrawal`

#### Market Data API (`pkg/marketdata`)
- **Reorganized into V2 (stable) and V1Beta (beta) clients**
- V2 (stocks): Stock bars, trades, quotes, snapshots, auctions, metadata
- V1Beta (beta features):
  - Crypto: bars, trades, quotes, snapshots, orderbooks
  - Options: bars, trades, quotes, snapshots, chains, metadata
  - News and screener endpoints
  - Corporate actions
  - Forex rates
  - Company logos

#### Broker API (`pkg/broker`)
- **Reorganized into V1 (core), V1Beta (beta), and V2 (events) clients**
- V1 new endpoints:
  - Documents: `ListAccountDocuments`, `DownloadAccountDocument`, `DownloadW8BenDocument`
  - Watchlists: Full CRUD for broker-managed watchlists
  - Journals: `CreateJournal`, `ListJournals`, `GetJournal`, `DeleteJournal`, `CreateBatchJournal`, `ReverseBatchJournal`
  - Rebalancing: Portfolios, subscriptions, and runs management
  - CIP/KYC: `GetCIPInfo`, `UpdateCIPInfo`
  - Onfido integration: Applicants, checks, SDK tokens, documents
  - OAuth: Token issuance, authorization, client management
  - Options: Approval requests and contract listing
  - Crypto: Wallet management, transfers, whitelisted addresses
  - FPSL (Fully Paid Securities Lending): Loans, tiers, APR tiers
  - JIT (Just-In-Time): Ledgers, balances, limits, settlements
  - Instant Funding: Listings, deletions, settlements
  - IRA: Excess contribution tracking
  - Calendar: Market calendar and clock endpoints
- V1Beta: Funding wallets and recipient bank management
- V2: SSE event streaming for trades, journals, system events, admin actions, funding status

### Changed
- Market Data client now exposes `V2` and `V1Beta` sub-clients for clear API versioning
- Broker client now exposes `V1`, `V1Beta`, and `V2` sub-clients for clear API versioning
- Moved crypto, options, news, and screener endpoints from Market Data v2 to v1beta
- Enhanced proto definitions with improved validation annotations

### Infrastructure
- Added `cmd/apitest` test harness for validating generated API clients
- Added Make targets: `test-marketdata`, `test-broker`, `test-trading`
- Updated golangci-lint configuration
- Added goimports to code generation pipeline

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

[unreleased]: https://github.com/sebastienmelki/alpaca-go/compare/v1.2.1...HEAD
[1.2.1]: https://github.com/sebastienmelki/alpaca-go/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/sebastienmelki/alpaca-go/releases/tag/v1.0.0
