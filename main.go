package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sebastienmelki/alpaca-go/pkg/marketdata"
)

func main() {
	client := marketdata.NewClient(
		"CKM3GFJ6Q8CQ7ZGEWTNN",
		"NQ1Bv080gPZi1uGZTgZaB7DQodlLwWTcfQTtW8w6",
		marketdata.WithBaseURL("https://data.sandbox.alpaca.markets"),
	)
	ctx := context.Background()

	// Get latest stock bars
	bars, err := client.GetOptionBars(ctx, &marketdata.GetOptionBarsRequest{
		Symbols:   "TSLA260123C00335000",
		Start:     "2025-12-13T00:00:00Z",
		End:       "2026-01-23T00:00:00Z",
		Timeframe: "5Min",
	})
	if err != nil {
		log.Fatal(err)
	}
	for symbol, _bars := range bars.Bars {
		for _, bar := range _bars.Bars {
			fmt.Printf("%s: Close=%s, Volume=%d\n", symbol, bar.C, bar.V)
		}
	}

	// Get stock quotes
	quotes, err := client.GetLatestStockQuotes(ctx, &marketdata.GetLatestStockQuotesRequest{
		Symbols: "AAPL",
	})
	if err != nil {
		log.Fatal(err)
	}
	for symbol, quote := range quotes.Quotes {
		fmt.Printf("%s: Bid=%s, Ask=%s\n", symbol, quote.Bp, quote.Ap)
	}
}
