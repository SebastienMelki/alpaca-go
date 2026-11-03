# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- `TradeUpdateEventV2.reason` — Alpaca's top-level cause on trade-update events (`CORPORATE_ACTION` on `canceled`, `TOO_LATE_TO_CANCEL` on `order_cancel_rejected`/`order_replace_rejected`, `TRADE_BUST` on `trade_bust`, or free-form venue text). Empty when Alpaca sends none. Previously dropped on decode.
- `NTAEvent` — typed non-trading-activity payload for `/v1/events/nta`, re-exported from `pkg/broker`. It carries every scalar field the feed sends, including `group_id` and `event_ulid`.

### Changed
- **Breaking:** `SubscribeNTAEvents` now streams: it returns `(*BrokerServiceEventStream[*NTAEvent], error)` rather than `(*SubscribeSSEResponse, error)`. `/v1/events/nta` is a live SSE feed, so the unary form could never return. Callers should range over the stream with `Next` and `Close` it.

## [1.12.1] - 2026-09-17

### Added
- Broker Assets API: `ListAssets` (`GET /v1/assets`) and `GetAsset` (`GET /v1/assets/{symbol_or_id}`), with broker-specific fields (`attributes`, `borrow_status`, `margin_requirement_long`/`short`, etc.)
- `buf.validate` constraints on the assets query params `status`, `asset_class` and `attributes`, limited to Alpaca's documented values

## [1.12.0] - 2026-08-07

### Fixed
- `TradeUpdateEventV2.order` was typed `string`, but `/v2/events/trades` sends a nested JSON object. `protojson` rejected it with `invalid value for string field order`, so **every** trade-update event failed to unmarshal and the stream yielded nothing.
- OpenAPI specs: examples on string-typed fields were emitted as untagged YAML scalars (e.g. `"50000.00"` parsed as a float), fixed by the sebuf v0.23.2 bump

### Added
- `TradeUpdateEventV2Order` — typed order payload for trade-update events, re-exported from `pkg/broker`.

### Changed
- **Breaking:** `TradeUpdateEventV2.Order` is now `*TradeUpdateEventV2Order` rather than `string`. Callers that unmarshalled the JSON string themselves should read the typed fields instead.
- Bumped sebuf to v0.23.2; `install-tools` now pins the sebuf plugins to the `go.mod` version and installs the missing `protoc-gen-go-http`

## [1.11.2] - 2026-07-07

### Fixed
- Added `json_name` annotations to `CreateTradingOrderRequest`, `BrokerTakeProfitSpec` and `BrokerStopLossSpec` so the client sends snake_case keys; previously `POST /trading/accounts/{id}/orders` failed with `422 invalid time_in_force`

## [1.11.1] - 2026-06-11

### Fixed
- Regenerated code with sebuf v0.18.0; v1.11.0 had been generated with an older sebuf, reverting the v1.10.0 generated-file updates

## [1.11.0] - 2026-06-11

### Changed
- **Breaking:** `IPOEvent.account_number` (field 7) replaced by `account_id` on allocation events

## [1.10.0] - 2026-06-10

### Fixed
- Bumped sebuf to v0.18.0, which generates `UnmarshalJSONSebuf` for messages with `int64_encoding=NUMBER` fields, so `DiscardUnknownFields` is honoured on those types (e.g. `IPOOffering`) instead of being silently dropped

## [1.9.0] - 2026-06-09

### Changed
- Regenerated all code and OpenAPI specs (`make all`)

## [1.8.0] - 2026-06-09

### Fixed
- **Breaking:** `GetIPOOffering` now returns `GetIPOOfferingResponse` (wrapping the offering in `data`) to match the API's response envelope

## [1.7.0] - 2026-05-22

### Added
- `GetIPOOffering` RPC for `GET /v1/ipos/{ipo_reference}`; `GetIPOOfferingRequest` re-exported from `pkg/broker`
- `ListIPOOfferingsRequest`: `availability`, `ticker`, `limit` and `page_token` query params
- `IPOOffering.sixty_minute_expiration_time`

## [1.6.0] - 2026-05-12

### Added
- Broker APIs: `GetTradingAccount`, `ListAccountActivities`, `ListAccountActivitiesByType` and `GetBrokerPortfolioHistory`, with public type aliases, OpenAPI docs and test harness coverage

### Fixed
- Pinned `json_name` on portfolio history fields so `profit_loss`, `profit_loss_pct`, `base_value` and `base_value_asof` are no longer dropped on unmarshal

## [1.5.0] - 2026-04-30

### Added
- CI workflow running `go build` and `go vet` on push and pull request

### Fixed
- Removed duplicate `V2CallOption` declarations in `pkg/broker` that broke compilation after two concurrent merges

## [1.4.2] - 2026-04-30

### Added
- IPO listing (v1) and IPO event streaming (v2) broker endpoints
- `DiscardUnknownFields` exposed at client level and per call across all public client packages (auth, broker, marketdata, trading)

## [1.4.1] - 2026-04-29

### Added
- `BrokerV2ServiceCallOption` and `WithBrokerV2ServiceCallDiscardUnknownFields` re-exported from `pkg/broker` for per-call discarding of unknown fields

## [1.4.0] - 2026-04-24

### Added
- `TradeUpdateEventV2.group_id`, previously dropped on decode
- Upgraded sebuf to v0.12.0 for `WithXxxDiscardUnknownFields` support in generated clients and SSE streams

## [1.3.2] - 2026-04-21

### Added
- Native SSE streaming for v2 broker event endpoints (sebuf v0.10.0): events are read one at a time through `EventStream` with `Next()`/`Err()`/`Close()` instead of buffering the whole body
- Create Tag workflow for releases, with optional version input and minor auto-bump

### Removed
- **Breaking:** wrapper response messages such as `SubscribeTradeEventsV2Response`; streaming RPCs return the event type directly

## [1.3.1] - 2026-03-04

### Fixed
- Added `json_name = "latestTrade"` to `OptionSnapshot.latest_trade` field in `marketdata/v1beta` to ensure correct JSON key mapping during deserialization

### Added
- `curl-verified.md` — reference document listing SDK methods verified against raw curl responses

## [1.3.0] - 2026-03-03

### Changed
- **Breaking**: Field types in `marketdata/v2` `Trade` and `Quote` messages corrected — trade size (`s`) and quote sizes (`as`, `bs`) changed from `int64` to `uint32`; trade ID (`i`) changed from `int64` to `uint64` to match the Alpaca API's actual types
- Upgraded sebuf from v0.2.1 to v0.7.0
- Updated buf.build dependencies (`sebuf` and `protovalidate` to latest commits)

### Fixed
- `int64` fields (volume, trade count, trade/quote sizes, trade ID, portfolio history timestamps, SSE pagination cursors, news article ID, screener fields) now serialize as JSON numbers instead of quoted strings, matching the Alpaca API wire format
- Added `INT64_ENCODING_NUMBER` annotation to all numeric `int64` fields across broker, market data, and trading proto definitions

### Added
- 12 new `*_encoding.pb.go` generated files implementing custom `MarshalJSON` for messages with `INT64_ENCODING_NUMBER` fields (broker v1, marketdata v1beta and v2, trading v1)

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

[unreleased]: https://github.com/sebastienmelki/alpaca-go/compare/v1.12.1...HEAD
[1.12.1]: https://github.com/sebastienmelki/alpaca-go/compare/v1.12.0...v1.12.1
[1.12.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.11.2...v1.12.0
[1.11.2]: https://github.com/sebastienmelki/alpaca-go/compare/v1.11.1...v1.11.2
[1.11.1]: https://github.com/sebastienmelki/alpaca-go/compare/v1.11.0...v1.11.1
[1.11.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.10.0...v1.11.0
[1.10.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.9.0...v1.10.0
[1.9.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.8.0...v1.9.0
[1.8.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.7.0...v1.8.0
[1.7.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.6.0...v1.7.0
[1.6.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.5.0...v1.6.0
[1.5.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.4.2...v1.5.0
[1.4.2]: https://github.com/sebastienmelki/alpaca-go/compare/v1.4.1...v1.4.2
[1.4.1]: https://github.com/sebastienmelki/alpaca-go/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.3.2...v1.4.0
[1.3.2]: https://github.com/sebastienmelki/alpaca-go/compare/v1.3.1...v1.3.2
[1.3.1]: https://github.com/sebastienmelki/alpaca-go/compare/v1.3.0...v1.3.1
[1.3.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.2.1...v1.3.0
[1.2.1]: https://github.com/sebastienmelki/alpaca-go/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/sebastienmelki/alpaca-go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/sebastienmelki/alpaca-go/releases/tag/v1.0.0
