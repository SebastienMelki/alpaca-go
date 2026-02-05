// Package main provides a test harness for validating generated API clients.
// This is used to verify proto files generate proper Go code and to showcase SDK usage.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// TestResult tracks the result of a single API test.
type TestResult struct {
	Name    string
	Success bool
	Error   error
	Details string
}

const (
	serviceMarketData = "marketdata"
	serviceBroker     = "broker"
	serviceTrading    = "trading"
)

func main() {
	_ = godotenv.Load() //nolint:errcheck // Optional .env file

	service := flag.String("service", "marketdata", "Service to test: marketdata, broker, trading")
	flag.Parse()

	switch *service {
	case serviceMarketData, "md":
		testMarketData()
	case serviceBroker, "b":
		testBroker()
	case serviceTrading, "t":
		fmt.Println("Trading API tests not yet implemented")
	default:
		log.Fatalf("Unknown service: %s. Use: marketdata (md), broker (b), or trading (t)", *service)
	}
}

// getEnv returns the environment variable value or a default.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func printResult(result TestResult) {
	icon := "+"
	status := "PASS"
	if !result.Success {
		icon = "x"
		status = "FAIL"
	}

	fmt.Printf("  [%s] %-35s %s", icon, result.Name+":", status)
	if result.Success && result.Details != "" {
		fmt.Printf(" - %s", result.Details)
	} else if result.Error != nil {
		fmt.Printf(" - %s", result.Error.Error())
	}
	fmt.Println()
}

func printSummary(results []TestResult) {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("Summary:")
	fmt.Println(strings.Repeat("=", 70))

	passed := 0
	failed := 0
	var failedTests []TestResult

	for _, r := range results {
		if r.Success {
			passed++
		} else {
			failed++
			failedTests = append(failedTests, r)
		}
	}

	fmt.Printf("Total Tests: %d\n", len(results))
	fmt.Printf("Passed: %d\n", passed)
	fmt.Printf("Failed: %d\n", failed)

	if len(failedTests) > 0 {
		fmt.Println("\nFailed Tests:")
		for _, r := range failedTests {
			fmt.Printf("  - %s: %s\n", r.Name, r.Error.Error())
		}
	}

	if failed == 0 {
		fmt.Println("\nAll tests passed!")
	} else {
		log.Printf("\n%d test(s) failed", failed)
	}
}
