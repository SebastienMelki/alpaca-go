package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/sebastienmelki/alpaca-go/pkg/broker"
)

var errTestAccountIDNotSet = errors.New("TEST_ACCOUNT_ID not set")

func testBroker() {
	apiKey := getEnv("ALPACA_API_KEY", "")
	apiSecret := getEnv("ALPACA_API_SECRET", "")
	testAccountID := getEnv("TEST_ACCOUNT_ID", "")

	if apiKey == "" || apiSecret == "" {
		log.Fatal("ALPACA_API_KEY and ALPACA_API_SECRET must be set in environment or .env file")
	}

	client := broker.NewSandboxClient(apiKey, apiSecret)

	ctx := context.Background()
	var results []TestResult

	fmt.Println("Testing Broker APIs...")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println()

	// =========================================================================
	// Account APIs
	// =========================================================================
	fmt.Println("Account APIs:")
	results = append(results,
		testBrokerGetAccount(ctx, client, testAccountID),
		testBrokerListAccounts(ctx, client),
	)
	fmt.Println()

	// =========================================================================
	// ACH Relationship APIs
	// =========================================================================
	fmt.Println("ACH Relationship APIs:")
	results = append(results,
		testBrokerListACHRelationships(ctx, client, testAccountID),
	)
	fmt.Println()

	// =========================================================================
	// Transfer APIs
	// =========================================================================
	fmt.Println("Transfer APIs:")
	results = append(results,
		testBrokerListTransfers(ctx, client, testAccountID),
	)
	fmt.Println()

	// =========================================================================
	// Trading Order APIs
	// =========================================================================
	fmt.Println("Trading Order APIs:")
	results = append(results,
		testBrokerListTradingOrders(ctx, client, testAccountID),
	)
	fmt.Println()

	// =========================================================================
	// Trading Position APIs
	// =========================================================================
	fmt.Println("Trading Position APIs:")
	results = append(results,
		testBrokerListTradingPositions(ctx, client, testAccountID),
	)
	fmt.Println()

	printSummary(results)
}

// =============================================================================
// Account API Tests
// =============================================================================

func testBrokerGetAccount(ctx context.Context, client *broker.Client, accountID string) TestResult {
	result := TestResult{Name: "GetAccount"}
	if accountID == "" {
		result.Error = errTestAccountIDNotSet
		printResult(result)
		return result
	}
	resp, err := client.V1.GetAccount(ctx, &broker.GetBrokerAccountRequest{
		AccountId: accountID,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Got account %s (status: %s)", resp.Id, resp.Status)
	printResult(result)
	return result
}

func testBrokerListAccounts(ctx context.Context, client *broker.Client) TestResult {
	result := TestResult{Name: "ListAccounts"}
	resp, err := client.V1.ListAccounts(ctx, &broker.ListAccountsRequest{})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Found %d accounts", len(resp.Accounts))
	printResult(result)
	return result
}

// =============================================================================
// ACH Relationship API Tests
// =============================================================================

func testBrokerListACHRelationships(ctx context.Context, client *broker.Client, accountID string) TestResult {
	result := TestResult{Name: "ListACHRelationships"}
	if accountID == "" {
		result.Error = errTestAccountIDNotSet
		printResult(result)
		return result
	}
	resp, err := client.V1.ListACHRelationships(ctx, &broker.ListACHRelationshipsRequest{
		AccountId: accountID,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Found %d ACH relationships", len(resp.AchRelationships))
	printResult(result)
	return result
}

// =============================================================================
// Transfer API Tests
// =============================================================================

func testBrokerListTransfers(ctx context.Context, client *broker.Client, accountID string) TestResult {
	result := TestResult{Name: "ListTransfers"}
	if accountID == "" {
		result.Error = errTestAccountIDNotSet
		printResult(result)
		return result
	}
	resp, err := client.V1.ListTransfers(ctx, &broker.ListTransfersRequest{
		AccountId: accountID,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Found %d transfers", len(resp.Transfers))
	printResult(result)
	return result
}

// =============================================================================
// Trading Order API Tests
// =============================================================================

func testBrokerListTradingOrders(ctx context.Context, client *broker.Client, accountID string) TestResult {
	result := TestResult{Name: "ListTradingOrders"}
	if accountID == "" {
		result.Error = errTestAccountIDNotSet
		printResult(result)
		return result
	}
	resp, err := client.V1.ListTradingOrders(ctx, &broker.ListTradingOrdersRequest{
		AccountId: accountID,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Found %d orders", len(resp.Orders))
	printResult(result)
	return result
}

// =============================================================================
// Trading Position API Tests
// =============================================================================

func testBrokerListTradingPositions(ctx context.Context, client *broker.Client, accountID string) TestResult {
	result := TestResult{Name: "ListTradingPositions"}
	if accountID == "" {
		result.Error = errTestAccountIDNotSet
		printResult(result)
		return result
	}
	resp, err := client.V1.ListTradingPositions(ctx, &broker.ListTradingPositionsRequest{
		AccountId: accountID,
	})
	if err != nil {
		result.Error = err
		printResult(result)
		return result
	}
	result.Success = true
	result.Details = fmt.Sprintf("Found %d positions", len(resp.Positions))
	printResult(result)
	return result
}
