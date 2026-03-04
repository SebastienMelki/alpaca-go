# Curl-Verified SDK Methods

This document lists SDK methods that have been manually verified against raw `curl` responses to confirm the SDK correctly maps API responses.

**Last updated:** 2026-03-04
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
