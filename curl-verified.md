# Curl-Verified SDK Methods

This document lists SDK methods that have been manually verified against raw `curl` responses to confirm the SDK correctly maps API responses.

**Last updated:** 2026-05-12
**Environment:** Sandbox (`data.sandbox.alpaca.markets`, `broker-api.sandbox.alpaca.markets`)

---

## General Notes

- **Key casing:** The SDK uses `camelCase` JSON keys (protobuf default). The raw API returns `snake_case`. This is expected protobuf Go behavior and not a bug — field values are identical.
- **`omitempty` behavior:** The SDK omits zero-value fields (e.g., `false` booleans, `0` integers, `null` messages, empty arrays) from its JSON output. The raw API returns these explicitly. This is standard protobuf Go JSON serialization — no data is lost.
- **Compact market data keys:** Stock bar and trade fields use Alpaca's single-letter compact format (`c`, `h`, `l`, `n`, `o`, `t`, `v`, `vw`) — both SDK and curl use the same keys.
- **S3 presigned URLs:** Document download URLs include time-limited signatures (`X-Amz-Date`, `X-Amz-Signature`) that differ between SDK and curl calls due to request timing. This is expected.

---

## Marketdata V2 (`data.sandbox.alpaca.markets/v2`)

Client: `marketdata.NewSandboxClient(apiKey, apiSecret).V2`

| Method | Endpoint | Result | Notes |
|---|---|---|---|
| `GetStockBars` | `GET /v2/stocks/bars` | ✅ Match | Values and structure identical. Tested with `AAPL,TSLA`, `1Day` timeframe, `iex` feed. |
| `GetLatestStockTrades` | `GET /v2/stocks/trades/latest` | ✅ Match | Values and structure identical. Tested with `AAPL,TSLA`, `iex` feed. |

---

## Marketdata V1Beta (`data.sandbox.alpaca.markets/v1beta1`)

Client: `marketdata.NewSandboxClient(apiKey, apiSecret).V1Beta`

| Method | Endpoint | Result | Notes |
|---|---|---|---|
| `GetOptionBars` | `GET /v1beta1/options/bars` | ✅ Match | Both SDK and curl returned empty `bars` for tested symbol (expired contract). SDK omits `next_page_token: null` due to `omitempty`. |
| `GetOptionSnapshots` | `GET /v1beta1/options/snapshots` | ✅ Match | Both returned empty snapshots for tested symbol (expired contract). SDK omits `next_page_token: null` due to `omitempty`. |
| `GetOptionChain` | `GET /v1beta1/options/snapshots/{underlying_symbol}` | ✅ Match | Returned 100 snapshots for `AAPL`. Structure matches. |

---

## Broker (`broker-api.sandbox.alpaca.markets`)

Client: `broker.NewSandboxClient(brokerKey, brokerSecret).V1`
Auth: HTTP Basic Auth (`Authorization: Basic base64(key:secret)`)

| Method | Endpoint | Result | Notes |
|---|---|---|---|
| `ListAccounts` | `GET /v1/accounts` | ✅ Match | All account data identical. Only key casing and `omitempty` differences (see General Notes). |
| `GetAccount` | `GET /v1/accounts/{account_id}` | ✅ Match | All data identical. Notable `omitempty` cases: `disclosures` shows `{}` in SDK vs explicit `false` booleans in curl; `number_of_dependents: 0` and `trading_configurations: null` omitted by SDK. |
| `GetTradingAccount` | `GET /v1/trading/accounts/{account_id}/account` | ✅ Match | All fields match. SDK omits zero-value booleans (`trading_blocked`, `transfers_blocked`, `account_blocked`, `trade_suspended_by_user`) and ints (`daytrade_count: 0`, `last_daytrade_count: 0`), and omits `user_configurations: null`. Tiny sub-cent drift on real-time balance fields (e.g. `regt_buying_power`, `non_marginable_buying_power`) between back-to-back calls is expected — the sandbox keeps re-marking the portfolio. |
| `ListAccountActivities` | `GET /v1/accounts/activities` | ✅ Match | All values identical. Only key casing differs (SDK `camelCase` vs curl `snake_case`). Tested with `page_size=5` filtered by `account_id`. |
| `ListAccountActivitiesByType` | `GET /v1/accounts/activities/{activity_type}` | ✅ Match | All values identical for `activity_type=FILL`. Same key-casing diff as above. |
| `GetBrokerPortfolioHistory` | `GET /v1/trading/accounts/{account_id}/account/portfolio/history` | ✅ Match | All time-series fields and scalars populate identically to curl. Requires `[json_name = "..."]` on `profit_loss`, `profit_loss_pct`, `base_value`, `base_value_asof` in `broker_portfolio_history.proto` — without those, `protoc-gen-go-http`'s unwrap codegen emits lowerCamelCase keys (`profitLoss`, `baseValue`, …) which don't match Alpaca's snake_case response and silently drop those fields. With `json_name` applied, both marshal and unmarshal use snake_case and round-trip cleanly. Trailing-zero formatting on floats (curl `3742.440000` vs SDK `3742.44`) is `encoding/json` behavior — values are bit-identical. |
