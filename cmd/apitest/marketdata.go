package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sebastienmelki/alpaca-go/pkg/marketdata"
)

// Test constants for market data.
const (
	stockSymbols  = "AAPL,TSLA"
	stockSymbol   = "AAPL"
	cryptoSymbols = "BTC/USD,ETH/USD"
	cryptoSymbol  = "BTC/USD"
	optionSymbol  = "AAPL250117C00150000"
	cryptoLoc     = "us"
	defaultTF     = "1Day"
)

func testMarketData() {
	apiKey := getEnv("ALPACA_API_KEY", "")
	apiSecret := getEnv("ALPACA_API_SECRET", "")

	client := marketdata.NewClient(
		apiKey,
		apiSecret,
		marketdata.WithBaseURL("https://data.sandbox.alpaca.markets"),
	)

	ctx := context.Background()
	var results []TestResult

	// Calculate date range (last 5 days).
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -5)
	startStr := startDate.Format(time.RFC3339)
	endStr := endDate.Format(time.RFC3339)

	fmt.Println("Testing Market Data APIs...")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println()

	// =========================================================================
	// Stock APIs
	// =========================================================================
	fmt.Println("Stock APIs:")
	results = append(results,
		testGetStockBars(ctx, client, startStr, endStr),
		testGetLatestStockBars(ctx, client),
		testGetStockBarsSingle(ctx, client, startStr, endStr),
		testGetLatestStockBarSingle(ctx, client),
		testGetStockTrades(ctx, client, startStr),
		testGetLatestStockTrades(ctx, client),
		testGetStockTradesSingle(ctx, client, startStr),
		testGetLatestStockTradeSingle(ctx, client),
		testGetStockQuotes(ctx, client, startStr),
		testGetLatestStockQuotes(ctx, client),
		testGetStockQuotesSingle(ctx, client, startStr),
		testGetLatestStockQuoteSingle(ctx, client),
		testGetStockSnapshots(ctx, client),
		testGetStockSnapshot(ctx, client),
		testGetStockAuctions(ctx, client, startStr),
		testGetStockAuctionsSingle(ctx, client, startStr),
		testGetStockMetaConditions(ctx, client),
		testGetStockMetaExchanges(ctx, client),
	)
	fmt.Println()

	// =========================================================================
	// Crypto APIs
	// =========================================================================
	fmt.Println("Crypto APIs:")
	results = append(results,
		testGetCryptoBars(ctx, client, startStr),
		testGetLatestCryptoBars(ctx, client),
		testGetCryptoTrades(ctx, client, startStr),
		testGetLatestCryptoTrades(ctx, client),
		testGetCryptoQuotes(ctx, client, startStr),
		testGetLatestCryptoQuotes(ctx, client),
		testGetCryptoSnapshots(ctx, client),
		testGetCryptoOrderbooks(ctx, client),
	)
	fmt.Println()

	// =========================================================================
	// Option APIs
	// =========================================================================
	fmt.Println("Option APIs:")
	results = append(results,
		testGetOptionBars(ctx, client, startStr),
		testGetOptionTrades(ctx, client, startStr),
		testGetLatestOptionTrades(ctx, client),
		testGetLatestOptionQuotes(ctx, client),
		testGetOptionSnapshots(ctx, client),
		testGetOptionChain(ctx, client),
		testGetOptionMetaConditions(ctx, client),
		testGetOptionMetaExchanges(ctx, client),
	)
	fmt.Println()

	// =========================================================================
	// News API
	// =========================================================================
	fmt.Println("News API:")
	results = append(results,
		testGetNews(ctx, client),
	)
	fmt.Println()

	// =========================================================================
	// Screener APIs
	// =========================================================================
	fmt.Println("Screener APIs:")
	results = append(results,
		testGetMostActives(ctx, client),
		testGetMovers(ctx, client),
	)
	fmt.Println()

	// =========================================================================
	// Corporate Actions API
	// =========================================================================
	fmt.Println("Corporate Actions API:")
	results = append(results,
		testGetCorporateActions(ctx, client),
	)
	fmt.Println()

	// =========================================================================
	// Forex APIs
	// =========================================================================
	fmt.Println("Forex APIs:")
	results = append(results,
		testGetForexRates(ctx, client, startStr),
		testGetLatestForexRates(ctx, client),
	)
	fmt.Println()

	// =========================================================================
	// Fixed Income API
	// =========================================================================
	fmt.Println("Fixed Income API:")
	results = append(results,
		testGetFixedIncomeLatestPrices(ctx, client),
	)
	fmt.Println()

	// =========================================================================
	// Logo API
	// =========================================================================
	fmt.Println("Logo API:")
	results = append(results,
		testGetLogo(ctx, client),
	)
	fmt.Println()

	printSummary(results)
}

// =============================================================================
// Stock API Tests
// =============================================================================

func testGetStockBars(ctx context.Context, client *marketdata.Client, start, end string) TestResult {
	result := TestResult{Name: "GetStockBars"}
	resp, err := client.V2.GetStockBars(ctx, &marketdata.GetStockBarsRequest{
		Symbols:   stockSymbols,
		Timeframe: defaultTF,
		Start:     start,
		End:       end,
		Limit:     10,
		Feed:      "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got bars for %d symbols", len(resp.Bars))
	printResult(result)
	return result
}

func testGetLatestStockBars(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestStockBars"}
	resp, err := client.V2.GetLatestStockBars(ctx, &marketdata.GetLatestStockBarsRequest{
		Symbols: stockSymbols,
		Feed:    "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest bars for %d symbols", len(resp.Bars))
	printResult(result)
	return result
}

func testGetStockBarsSingle(ctx context.Context, client *marketdata.Client, start, end string) TestResult {
	result := TestResult{Name: "GetStockBarsSingle"}
	resp, err := client.V2.GetStockBarsSingle(ctx, &marketdata.GetStockBarsSingleRequest{
		Symbol:    stockSymbol,
		Timeframe: defaultTF,
		Start:     start,
		End:       end,
		Limit:     10,
		Feed:      "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d bars for %s", len(resp.Bars), stockSymbol)
	printResult(result)
	return result
}

func testGetLatestStockBarSingle(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestStockBarSingle"}
	resp, err := client.V2.GetLatestStockBarSingle(ctx, &marketdata.GetLatestStockBarSingleRequest{
		Symbol: stockSymbol,
		Feed:   "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest bar for %s (t=%s)", resp.Symbol, resp.Bar.GetT())
	printResult(result)
	return result
}

func testGetStockTrades(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetStockTrades"}
	resp, err := client.V2.GetStockTrades(ctx, &marketdata.GetStockTradesRequest{
		Symbols: stockSymbols,
		Start:   start,
		Limit:   10,
		Feed:    "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	totalTrades := 0
	for _, tradesList := range resp.Trades {
		totalTrades += len(tradesList.Trades)
	}
	result.Details = fmt.Sprintf("Got trades for %d symbols (%d total)", len(resp.Trades), totalTrades)
	printResult(result)
	return result
}

func testGetLatestStockTrades(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestStockTrades"}
	resp, err := client.V2.GetLatestStockTrades(ctx, &marketdata.GetLatestStockTradesRequest{
		Symbols: stockSymbols,
		Feed:    "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest trades for %d symbols", len(resp.Trades))
	printResult(result)
	return result
}

func testGetStockTradesSingle(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetStockTradesSingle"}
	resp, err := client.V2.GetStockTradesSingle(ctx, &marketdata.GetStockTradesSingleRequest{
		Symbol: stockSymbol,
		Start:  start,
		Limit:  10,
		Feed:   "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d trades for %s", len(resp.Trades), stockSymbol)
	printResult(result)
	return result
}

func testGetLatestStockTradeSingle(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestStockTradeSingle"}
	resp, err := client.V2.GetLatestStockTradeSingle(ctx, &marketdata.GetLatestStockTradeSingleRequest{
		Symbol: stockSymbol,
		Feed:   "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest trade for %s (p=%.2f)", resp.Symbol, resp.Trade.GetP())
	printResult(result)
	return result
}

func testGetStockQuotes(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetStockQuotes"}
	resp, err := client.V2.GetStockQuotes(ctx, &marketdata.GetStockQuotesRequest{
		Symbols: stockSymbols,
		Start:   start,
		Limit:   10,
		Feed:    "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	totalQuotes := 0
	for _, quotesList := range resp.Quotes {
		totalQuotes += len(quotesList.Quotes)
	}
	result.Details = fmt.Sprintf("Got quotes for %d symbols (%d total)", len(resp.Quotes), totalQuotes)
	printResult(result)
	return result
}

func testGetLatestStockQuotes(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestStockQuotes"}
	resp, err := client.V2.GetLatestStockQuotes(ctx, &marketdata.GetLatestStockQuotesRequest{
		Symbols: stockSymbols,
		Feed:    "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest quotes for %d symbols", len(resp.Quotes))
	printResult(result)
	return result
}

func testGetStockQuotesSingle(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetStockQuotesSingle"}
	resp, err := client.V2.GetStockQuotesSingle(ctx, &marketdata.GetStockQuotesSingleRequest{
		Symbol: stockSymbol,
		Start:  start,
		Limit:  10,
		Feed:   "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d quotes for %s", len(resp.Quotes), stockSymbol)
	printResult(result)
	return result
}

func testGetLatestStockQuoteSingle(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestStockQuoteSingle"}
	resp, err := client.V2.GetLatestStockQuoteSingle(ctx, &marketdata.GetLatestStockQuoteSingleRequest{
		Symbol: stockSymbol,
		Feed:   "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest quote for %s (bid=%.2f, ask=%.2f)", resp.Symbol, resp.Quote.GetBp(), resp.Quote.GetAp())
	printResult(result)
	return result
}

func testGetStockSnapshots(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetStockSnapshots"}
	resp, err := client.V2.GetStockSnapshots(ctx, &marketdata.GetStockSnapshotsRequest{
		Symbols: stockSymbols,
		Feed:    "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got snapshots for %d symbols", len(resp.Snapshots))
	printResult(result)
	return result
}

func testGetStockSnapshot(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetStockSnapshot"}
	resp, err := client.V2.GetStockSnapshot(ctx, &marketdata.GetStockSnapshotRequest{
		Symbol: stockSymbol,
		Feed:   "iex",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = "Got snapshot for " + resp.Symbol
	printResult(result)
	return result
}

func testGetStockAuctions(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetStockAuctions"}
	resp, err := client.V2.GetStockAuctions(ctx, &marketdata.GetStockAuctionsRequest{
		Symbols: stockSymbols,
		Start:   start,
		Limit:   10,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	totalAuctions := 0
	for _, auctionsList := range resp.Auctions {
		totalAuctions += len(auctionsList.Auctions)
	}
	result.Details = fmt.Sprintf("Got auctions for %d symbols (%d total)", len(resp.Auctions), totalAuctions)
	printResult(result)
	return result
}

func testGetStockAuctionsSingle(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetStockAuctionsSingle"}
	resp, err := client.V2.GetStockAuctionsSingle(ctx, &marketdata.GetStockAuctionsSingleRequest{
		Symbol: stockSymbol,
		Start:  start,
		Limit:  10,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d auctions for %s", len(resp.Auctions), stockSymbol)
	printResult(result)
	return result
}

func testGetStockMetaConditions(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetStockMetaConditions"}
	resp, err := client.V2.GetStockMetaConditions(ctx, &marketdata.GetStockMetaConditionsRequest{
		Ticktype: "trade",
		Tape:     "A",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d conditions", len(resp.Conditions))
	printResult(result)
	return result
}

func testGetStockMetaExchanges(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetStockMetaExchanges"}
	resp, err := client.V2.GetStockMetaExchanges(ctx, &marketdata.GetStockMetaExchangesRequest{})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d exchanges", len(resp.Exchanges))
	printResult(result)
	return result
}

// =============================================================================
// Crypto API Tests
// =============================================================================

func testGetCryptoBars(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetCryptoBars"}
	resp, err := client.V1Beta.GetCryptoBars(ctx, &marketdata.GetCryptoBarsRequest{
		Loc:       cryptoLoc,
		Symbols:   cryptoSymbols,
		Timeframe: "1Hour",
		Start:     start,
		Limit:     10,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	totalBars := 0
	for _, barsList := range resp.Bars {
		totalBars += len(barsList.Bars)
	}
	result.Details = fmt.Sprintf("Got bars for %d symbols (%d total)", len(resp.Bars), totalBars)
	printResult(result)
	return result
}

func testGetLatestCryptoBars(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestCryptoBars"}
	resp, err := client.V1Beta.GetLatestCryptoBars(ctx, &marketdata.GetLatestCryptoBarsRequest{
		Loc:     cryptoLoc,
		Symbols: cryptoSymbols,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest bars for %d symbols", len(resp.Bars))
	printResult(result)
	return result
}

func testGetCryptoTrades(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetCryptoTrades"}
	resp, err := client.V1Beta.GetCryptoTrades(ctx, &marketdata.GetCryptoTradesRequest{
		Loc:     cryptoLoc,
		Symbols: cryptoSymbols,
		Start:   start,
		Limit:   10,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	totalTrades := 0
	for _, tradesList := range resp.Trades {
		totalTrades += len(tradesList.Trades)
	}
	result.Details = fmt.Sprintf("Got trades for %d symbols (%d total)", len(resp.Trades), totalTrades)
	printResult(result)
	return result
}

func testGetLatestCryptoTrades(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestCryptoTrades"}
	resp, err := client.V1Beta.GetLatestCryptoTrades(ctx, &marketdata.GetLatestCryptoTradesRequest{
		Loc:     cryptoLoc,
		Symbols: cryptoSymbols,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest trades for %d symbols", len(resp.Trades))
	printResult(result)
	return result
}

func testGetCryptoQuotes(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetCryptoQuotes"}
	resp, err := client.V1Beta.GetCryptoQuotes(ctx, &marketdata.GetCryptoQuotesRequest{
		Loc:     cryptoLoc,
		Symbols: cryptoSymbols,
		Start:   start,
		Limit:   10,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	totalQuotes := 0
	for _, quotesList := range resp.Quotes {
		totalQuotes += len(quotesList.Quotes)
	}
	result.Details = fmt.Sprintf("Got quotes for %d symbols (%d total)", len(resp.Quotes), totalQuotes)
	printResult(result)
	return result
}

func testGetLatestCryptoQuotes(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestCryptoQuotes"}
	resp, err := client.V1Beta.GetLatestCryptoQuotes(ctx, &marketdata.GetLatestCryptoQuotesRequest{
		Loc:     cryptoLoc,
		Symbols: cryptoSymbols,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest quotes for %d symbols", len(resp.Quotes))
	printResult(result)
	return result
}

func testGetCryptoSnapshots(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetCryptoSnapshots"}
	resp, err := client.V1Beta.GetCryptoSnapshots(ctx, &marketdata.GetCryptoSnapshotsRequest{
		Loc:     cryptoLoc,
		Symbols: cryptoSymbols,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got snapshots for %d symbols", len(resp.Snapshots))
	printResult(result)
	return result
}

func testGetCryptoOrderbooks(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetCryptoOrderbooks"}
	resp, err := client.V1Beta.GetCryptoOrderbooks(ctx, &marketdata.GetCryptoOrderbooksRequest{
		Loc:     cryptoLoc,
		Symbols: cryptoSymbols,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got orderbooks for %d symbols", len(resp.Orderbooks))
	printResult(result)
	return result
}

// =============================================================================
// Option API Tests
// =============================================================================

func testGetOptionBars(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetOptionBars"}
	resp, err := client.V1Beta.GetOptionBars(ctx, &marketdata.GetOptionBarsRequest{
		Symbols:   optionSymbol,
		Timeframe: defaultTF,
		Start:     start,
		Limit:     10,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	totalBars := 0
	for _, barsList := range resp.Bars {
		totalBars += len(barsList.Bars)
	}
	result.Details = fmt.Sprintf("Got bars for %d options (%d total)", len(resp.Bars), totalBars)
	printResult(result)
	return result
}

func testGetOptionTrades(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetOptionTrades"}
	resp, err := client.V1Beta.GetOptionTrades(ctx, &marketdata.GetOptionTradesRequest{
		Symbols: optionSymbol,
		Start:   start,
		Limit:   10,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	totalTrades := 0
	for _, tradesList := range resp.Trades {
		totalTrades += len(tradesList.Trades)
	}
	result.Details = fmt.Sprintf("Got trades for %d options (%d total)", len(resp.Trades), totalTrades)
	printResult(result)
	return result
}

func testGetLatestOptionTrades(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestOptionTrades"}
	resp, err := client.V1Beta.GetLatestOptionTrades(ctx, &marketdata.GetLatestOptionTradesRequest{
		Symbols: optionSymbol,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest trades for %d options", len(resp.Trades))
	printResult(result)
	return result
}

func testGetLatestOptionQuotes(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestOptionQuotes"}
	resp, err := client.V1Beta.GetLatestOptionQuotes(ctx, &marketdata.GetLatestOptionQuotesRequest{
		Symbols: optionSymbol,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest quotes for %d options", len(resp.Quotes))
	printResult(result)
	return result
}

func testGetOptionSnapshots(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetOptionSnapshots"}
	resp, err := client.V1Beta.GetOptionSnapshots(ctx, &marketdata.GetOptionSnapshotsRequest{
		Symbols: optionSymbol,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got snapshots for %d options", len(resp.Snapshots))
	printResult(result)
	return result
}

func testGetOptionChain(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetOptionChain"}
	resp, err := client.V1Beta.GetOptionChain(ctx, &marketdata.GetOptionChainRequest{
		UnderlyingSymbol: stockSymbol,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got option chain with %d options", len(resp.Snapshots))
	printResult(result)
	return result
}

func testGetOptionMetaConditions(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetOptionMetaConditions"}
	resp, err := client.V1Beta.GetOptionMetaConditions(ctx, &marketdata.GetOptionMetaConditionsRequest{
		Ticktype: "trade",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d conditions", len(resp.Conditions))
	printResult(result)
	return result
}

func testGetOptionMetaExchanges(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetOptionMetaExchanges"}
	resp, err := client.V1Beta.GetOptionMetaExchanges(ctx, &marketdata.GetOptionMetaExchangesRequest{})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d exchanges", len(resp.Exchanges))
	printResult(result)
	return result
}

// =============================================================================
// News API Tests
// =============================================================================

func testGetNews(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetNews"}
	resp, err := client.V1Beta.GetNews(ctx, &marketdata.GetNewsRequest{
		Symbols: stockSymbol,
		Limit:   10,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d news articles", len(resp.News))
	printResult(result)
	return result
}

// =============================================================================
// Screener API Tests
// =============================================================================

func testGetMostActives(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetMostActives"}
	resp, err := client.V1Beta.GetMostActives(ctx, &marketdata.GetMostActivesRequest{})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d most active stocks", len(resp.MostActives))
	printResult(result)
	return result
}

func testGetMovers(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetMovers"}
	resp, err := client.V1Beta.GetMovers(ctx, &marketdata.GetMoversRequest{
		MarketType: "stocks",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got %d gainers and %d losers", len(resp.Gainers), len(resp.Losers))
	printResult(result)
	return result
}

// =============================================================================
// Corporate Actions API Tests
// =============================================================================

func testGetCorporateActions(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetCorporateActions"}
	resp, err := client.V1Beta.GetCorporateActions(ctx, &marketdata.GetCorporateActionsRequest{
		Symbols: stockSymbol,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	ca := resp.CorporateActions
	total := len(ca.GetCashDividends()) + len(ca.GetStockDividends()) + len(ca.GetForwardSplits()) + len(ca.GetReverseSplits())
	result.Details = fmt.Sprintf("Got %d corporate actions", total)
	printResult(result)
	return result
}

// =============================================================================
// Forex API Tests
// =============================================================================

func testGetForexRates(ctx context.Context, client *marketdata.Client, start string) TestResult {
	result := TestResult{Name: "GetForexRates"}
	resp, err := client.V1Beta.GetForexRates(ctx, &marketdata.GetForexRatesRequest{
		CurrencyPairs: "EUR/USD",
		Timeframe:     "1Day",
		Start:         start,
		Limit:         10,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got rates for %d pairs", len(resp.Rates))
	printResult(result)
	return result
}

func testGetLatestForexRates(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLatestForexRates"}
	resp, err := client.V1Beta.GetLatestForexRates(ctx, &marketdata.GetLatestForexRatesRequest{
		CurrencyPairs: "EUR/USD",
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got latest rates for %d pairs", len(resp.Rates))
	printResult(result)
	return result
}

// =============================================================================
// Fixed Income API Tests
// =============================================================================

func testGetFixedIncomeLatestPrices(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetFixedIncomeLatestPrices"}
	resp, err := client.V1Beta.GetFixedIncomeLatestPrices(ctx, &marketdata.GetFixedIncomeLatestPricesRequest{
		Isins: "912797HE8", // US Treasury Bill ISIN
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got prices for %d instruments", len(resp.Prices))
	printResult(result)
	return result
}

// =============================================================================
// Logo API Tests
// =============================================================================

func testGetLogo(ctx context.Context, client *marketdata.Client) TestResult {
	result := TestResult{Name: "GetLogo"}
	resp, err := client.GetLogo(ctx, &marketdata.GetLogoRequest{
		Symbol:      stockSymbol,
		Placeholder: true,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got logo (%d bytes, %s)", len(resp.Image), resp.ContentType)
	printResult(result)
	return result
}
