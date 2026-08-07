package broker

import (
	"encoding/base64"
	"net/http"

	"google.golang.org/protobuf/proto"

	brokerv1 "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/broker/v1"
	brokerv1beta "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/broker/v1beta"
	brokerv2 "github.com/sebastienmelki/alpaca-go/internal/gen/alpaca/broker/v2"
)

const (
	LiveBaseURL    = "https://broker-api.alpaca.markets"
	SandboxBaseURL = "https://broker-api.sandbox.alpaca.markets"
)

// Client wraps the generated BrokerService v1, v1beta, and v2 clients with Alpaca-specific defaults.
type Client struct {
	V1 brokerv1.BrokerServiceClient

	// V1Beta provides access to v1beta funding wallet and recipient bank endpoints.
	V1Beta brokerv1beta.BrokerV1BetaServiceClient

	// V2 provides access to v2 SSE event endpoints.
	V2 brokerv2.BrokerV2ServiceClient
}

// Option configures a Client.
type Option func(*options)

type options struct {
	httpClient           *http.Client
	baseURL              string
	discardUnknownFields bool
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(c *http.Client) Option {
	return func(o *options) { o.httpClient = c }
}

// WithBaseURL sets a custom base URL (defaults to LiveBaseURL).
func WithBaseURL(url string) Option {
	return func(o *options) { o.baseURL = url }
}

// WithDiscardUnknownFields sets whether to discard unknown fields in JSON responses.
// When true, unknown fields are silently ignored instead of causing unmarshal errors.
func WithDiscardUnknownFields(discard bool) Option {
	return func(o *options) { o.discardUnknownFields = discard }
}

// NewClient creates a new Broker API client.
// The Broker API uses HTTP Basic Auth with API key and secret.
func NewClient(apiKey, apiSecret string, opts ...Option) *Client {
	cfg := &options{
		httpClient: http.DefaultClient,
		baseURL:    LiveBaseURL,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	// Broker API uses Basic Auth
	auth := base64.StdEncoding.EncodeToString([]byte(apiKey + ":" + apiSecret))

	v1Client := brokerv1.NewBrokerServiceClient(
		cfg.baseURL,
		brokerv1.WithBrokerServiceHTTPClient(cfg.httpClient),
		brokerv1.WithBrokerServiceAuthorization("Basic "+auth),
		brokerv1.WithBrokerServiceDiscardUnknownFields(cfg.discardUnknownFields),
	)

	v1betaClient := brokerv1beta.NewBrokerV1BetaServiceClient(
		cfg.baseURL,
		brokerv1beta.WithBrokerV1BetaServiceHTTPClient(cfg.httpClient),
		brokerv1beta.WithBrokerV1BetaServiceAuthorization("Basic "+auth),
		brokerv1beta.WithBrokerV1BetaServiceDiscardUnknownFields(cfg.discardUnknownFields),
	)

	v2Client := brokerv2.NewBrokerV2ServiceClient(
		cfg.baseURL,
		brokerv2.WithBrokerV2ServiceHTTPClient(cfg.httpClient),
		brokerv2.WithBrokerV2ServiceAuthorization("Basic "+auth),
		brokerv2.WithBrokerV2ServiceDiscardUnknownFields(cfg.discardUnknownFields),
	)

	return &Client{
		V1:     v1Client,
		V1Beta: v1betaClient,
		V2:     v2Client,
	}
}

// NewSandboxClient creates a client for the broker sandbox environment.
func NewSandboxClient(apiKey, apiSecret string, opts ...Option) *Client {
	return NewClient(apiKey, apiSecret, append(opts, WithBaseURL(SandboxBaseURL))...)
}

// =============================================================================
// Core Account Types
// =============================================================================

type (
	BrokerAccount   = brokerv1.BrokerAccount
	BrokerOrder     = brokerv1.BrokerOrder
	BrokerPosition  = brokerv1.BrokerPosition
	ACHRelationship = brokerv1.ACHRelationship
	Transfer        = brokerv1.Transfer
)

// Account-related types.
type (
	Contact        = brokerv1.Contact
	Identity       = brokerv1.Identity
	Disclosures    = brokerv1.Disclosures
	Agreement      = brokerv1.Agreement
	TrustedContact = brokerv1.TrustedContact
)

// =============================================================================
// Account Request/Response Types
// =============================================================================

type (
	CreateAccountRequest       = brokerv1.CreateAccountRequest
	ListAccountsRequest        = brokerv1.ListAccountsRequest
	ListAccountsResponse       = brokerv1.ListAccountsResponse
	GetBrokerAccountRequest    = brokerv1.GetBrokerAccountRequest
	UpdateBrokerAccountRequest = brokerv1.UpdateBrokerAccountRequest
	CloseBrokerAccountRequest  = brokerv1.CloseBrokerAccountRequest
	CloseBrokerAccountResponse = brokerv1.CloseBrokerAccountResponse
)

// =============================================================================
// Trading Account Details Types
// =============================================================================

type (
	TradeAccount                = brokerv1.TradeAccount
	AdminConfigurations         = brokerv1.AdminConfigurations
	AccountConfigurations       = brokerv1.AccountConfigurations
	RestrictToLiquidationReasons = brokerv1.RestrictToLiquidationReasons
	GetTradingAccountRequest    = brokerv1.GetTradingAccountRequest
)

// =============================================================================
// Account Activity Types
// =============================================================================

type (
	AccountActivity                    = brokerv1.AccountActivity
	ListAccountActivitiesRequest       = brokerv1.ListAccountActivitiesRequest
	ListAccountActivitiesResponse      = brokerv1.ListAccountActivitiesResponse
	ListAccountActivitiesByTypeRequest = brokerv1.ListAccountActivitiesByTypeRequest
)

// =============================================================================
// Broker Portfolio History Types
// =============================================================================

type (
	BrokerPortfolioHistory            = brokerv1.BrokerPortfolioHistory
	CashflowValues                    = brokerv1.CashflowValues
	GetBrokerPortfolioHistoryRequest  = brokerv1.GetBrokerPortfolioHistoryRequest
)

// =============================================================================
// ACH Relationship Types
// =============================================================================

type (
	CreateACHRelationshipRequest  = brokerv1.CreateACHRelationshipRequest
	ListACHRelationshipsRequest   = brokerv1.ListACHRelationshipsRequest
	ListACHRelationshipsResponse  = brokerv1.ListACHRelationshipsResponse
	DeleteACHRelationshipRequest  = brokerv1.DeleteACHRelationshipRequest
	DeleteACHRelationshipResponse = brokerv1.DeleteACHRelationshipResponse
)

// =============================================================================
// Transfer Types
// =============================================================================

type (
	CreateTransferRequest  = brokerv1.CreateTransferRequest
	ListTransfersRequest   = brokerv1.ListTransfersRequest
	ListTransfersResponse  = brokerv1.ListTransfersResponse
	GetTransferRequest     = brokerv1.GetTransferRequest
	CancelTransferRequest  = brokerv1.CancelTransferRequest
	CancelTransferResponse = brokerv1.CancelTransferResponse
)

// =============================================================================
// Trading Order Types (for accounts)
// =============================================================================

type (
	CreateTradingOrderRequest  = brokerv1.CreateTradingOrderRequest
	ListTradingOrdersRequest   = brokerv1.ListTradingOrdersRequest
	ListTradingOrdersResponse  = brokerv1.ListTradingOrdersResponse
	GetTradingOrderRequest     = brokerv1.GetTradingOrderRequest
	CancelTradingOrderRequest  = brokerv1.CancelTradingOrderRequest
	CancelTradingOrderResponse = brokerv1.CancelTradingOrderResponse
)

// =============================================================================
// Trading Position Types (for accounts)
// =============================================================================

type (
	ListTradingPositionsRequest      = brokerv1.ListTradingPositionsRequest
	ListTradingPositionsResponse     = brokerv1.ListTradingPositionsResponse
	GetTradingPositionRequest        = brokerv1.GetTradingPositionRequest
	CloseTradingPositionRequest      = brokerv1.CloseTradingPositionRequest
	CloseAllTradingPositionsRequest  = brokerv1.CloseAllTradingPositionsRequest
	CloseAllTradingPositionsResponse = brokerv1.CloseAllTradingPositionsResponse
	ClosedTradingPosition            = brokerv1.ClosedTradingPosition
)

// =============================================================================
// Journal Types
// =============================================================================

type (
	Journal                     = brokerv1.Journal
	CreateJournalRequest        = brokerv1.CreateJournalRequest
	ListJournalsRequest         = brokerv1.ListJournalsRequest
	ListJournalsResponse        = brokerv1.ListJournalsResponse
	GetJournalRequest           = brokerv1.GetJournalRequest
	DeleteJournalRequest        = brokerv1.DeleteJournalRequest
	DeleteJournalResponse       = brokerv1.DeleteJournalResponse
	CreateBatchJournalRequest   = brokerv1.CreateBatchJournalRequest
	CreateBatchJournalResponse  = brokerv1.CreateBatchJournalResponse
	ReverseBatchJournalRequest  = brokerv1.ReverseBatchJournalRequest
	ReverseBatchJournalResponse = brokerv1.ReverseBatchJournalResponse
)

// =============================================================================
// CIP (Customer Identification Program) Types
// =============================================================================

type (
	CIPInfo              = brokerv1.CIPInfo
	CIPKYCResult         = brokerv1.CIPKYCResult
	CIPDocumentResult    = brokerv1.CIPDocumentResult
	CIPPhotoResult       = brokerv1.CIPPhotoResult
	CIPIdentityResult    = brokerv1.CIPIdentityResult
	CIPWatchlistResult   = brokerv1.CIPWatchlistResult
	GetCIPInfoRequest    = brokerv1.GetCIPInfoRequest
	UpdateCIPInfoRequest = brokerv1.UpdateCIPInfoRequest
)

// =============================================================================
// Document Types
// =============================================================================

type (
	AccountDocument                 = brokerv1.AccountDocument
	ListAccountDocumentsRequest     = brokerv1.ListAccountDocumentsRequest
	ListAccountDocumentsResponse    = brokerv1.ListAccountDocumentsResponse
	DownloadAccountDocumentRequest  = brokerv1.DownloadAccountDocumentRequest
	DownloadAccountDocumentResponse = brokerv1.DownloadAccountDocumentResponse
	UploadAccountDocumentRequest    = brokerv1.UploadAccountDocumentRequest
	DownloadW8BenDocumentRequest    = brokerv1.DownloadW8BenDocumentRequest
)

// =============================================================================
// Watchlist Types
// =============================================================================

type (
	BrokerWatchlist                    = brokerv1.BrokerWatchlist
	BrokerWatchlistAsset               = brokerv1.BrokerWatchlistAsset
	ListBrokerWatchlistsRequest        = brokerv1.ListBrokerWatchlistsRequest
	ListBrokerWatchlistsResponse       = brokerv1.ListBrokerWatchlistsResponse
	CreateBrokerWatchlistRequest       = brokerv1.CreateBrokerWatchlistRequest
	GetBrokerWatchlistRequest          = brokerv1.GetBrokerWatchlistRequest
	UpdateBrokerWatchlistRequest       = brokerv1.UpdateBrokerWatchlistRequest
	DeleteBrokerWatchlistRequest       = brokerv1.DeleteBrokerWatchlistRequest
	DeleteBrokerWatchlistResponse      = brokerv1.DeleteBrokerWatchlistResponse
	AddBrokerWatchlistAssetRequest     = brokerv1.AddBrokerWatchlistAssetRequest
	RemoveBrokerWatchlistAssetRequest  = brokerv1.RemoveBrokerWatchlistAssetRequest
	RemoveBrokerWatchlistAssetResponse = brokerv1.RemoveBrokerWatchlistAssetResponse
)

// =============================================================================
// Calendar Types
// =============================================================================

type (
	MarketCalendar           = brokerv1.MarketDay
	MarketClock              = brokerv1.MarketClock
	GetMarketCalendarRequest = brokerv1.GetMarketCalendarRequest
	GetMarketClockRequest    = brokerv1.GetMarketClockRequest
)

// =============================================================================
// OAuth Types
// =============================================================================

type (
	OAuthToken                       = brokerv1.OAuthToken
	OAuthClient                      = brokerv1.OAuthClient
	OAuthAuthorization               = brokerv1.OAuthAuthorization
	CreateOAuthTokenRequest          = brokerv1.CreateOAuthTokenRequest
	AuthorizeOAuthRequest            = brokerv1.AuthorizeOAuthRequest
	AuthorizeOAuthResponse           = brokerv1.AuthorizeOAuthResponse
	GetOAuthClientRequest            = brokerv1.GetOAuthClientRequest
	CreateOAuthClientRequest         = brokerv1.CreateOAuthClientRequest
	UpdateOAuthClientRequest         = brokerv1.UpdateOAuthClientRequest
	DeleteOAuthClientRequest         = brokerv1.DeleteOAuthClientRequest
	DeleteOAuthClientResponse        = brokerv1.DeleteOAuthClientResponse
	RevokeOAuthAuthorizationRequest  = brokerv1.RevokeOAuthAuthorizationRequest
	RevokeOAuthAuthorizationResponse = brokerv1.RevokeOAuthAuthorizationResponse
)

// =============================================================================
// Onfido (KYC) Types
// =============================================================================

type (
	OnfidoApplicant               = brokerv1.OnfidoApplicant
	OnfidoCheck                   = brokerv1.OnfidoCheck
	OnfidoReport                  = brokerv1.OnfidoReport
	OnfidoSDKToken                = brokerv1.OnfidoSDKToken
	OnfidoDocument                = brokerv1.OnfidoDocument
	OnfidoPhoto                   = brokerv1.OnfidoPhoto
	CreateOnfidoApplicantRequest  = brokerv1.CreateOnfidoApplicantRequest
	GetOnfidoApplicantRequest     = brokerv1.GetOnfidoApplicantRequest
	GenerateOnfidoSDKTokenRequest = brokerv1.GenerateOnfidoSDKTokenRequest
	CreateOnfidoCheckRequest      = brokerv1.CreateOnfidoCheckRequest
	GetOnfidoCheckRequest         = brokerv1.GetOnfidoCheckRequest
	ListOnfidoChecksRequest       = brokerv1.ListOnfidoChecksRequest
	ListOnfidoChecksResponse      = brokerv1.ListOnfidoChecksResponse
	UploadOnfidoDocumentRequest   = brokerv1.UploadOnfidoDocumentRequest
	UploadOnfidoPhotoRequest      = brokerv1.UploadOnfidoPhotoRequest
)

// =============================================================================
// Instant Funding Types
// =============================================================================

type (
	InstantFunding                        = brokerv1.InstantFunding
	InstantFundingInterest                = brokerv1.InstantFundingInterest
	InstantFundingFee                     = brokerv1.InstantFundingFee
	InstantFundingSettlement              = brokerv1.InstantFundingSettlement
	ListInstantFundingRequest             = brokerv1.ListInstantFundingRequest
	ListInstantFundingResponse            = brokerv1.ListInstantFundingResponse
	DeleteInstantFundingRequest           = brokerv1.DeleteInstantFundingRequest
	DeleteInstantFundingResponse          = brokerv1.DeleteInstantFundingResponse
	ListInstantFundingSettlementsRequest  = brokerv1.ListInstantFundingSettlementsRequest
	ListInstantFundingSettlementsResponse = brokerv1.ListInstantFundingSettlementsResponse
	CreateInstantFundingSettlementRequest = brokerv1.CreateInstantFundingSettlementRequest
)

// =============================================================================
// Broker Crypto Funding Types
// =============================================================================

type (
	BrokerCryptoWallet                     = brokerv1.BrokerCryptoWallet
	BrokerCryptoTransfer                   = brokerv1.BrokerCryptoTransfer
	BrokerWhitelistedAddress               = brokerv1.BrokerWhitelistedAddress
	ListBrokerCryptoWalletsRequest         = brokerv1.ListBrokerCryptoWalletsRequest
	ListBrokerCryptoWalletsResponse        = brokerv1.ListBrokerCryptoWalletsResponse
	ListBrokerCryptoTransfersRequest       = brokerv1.ListBrokerCryptoTransfersRequest
	ListBrokerCryptoTransfersResponse      = brokerv1.ListBrokerCryptoTransfersResponse
	GetBrokerCryptoTransferRequest         = brokerv1.GetBrokerCryptoTransferRequest
	CreateBrokerCryptoTransferRequest      = brokerv1.CreateBrokerCryptoTransferRequest
	ListBrokerWhitelistedAddressesRequest  = brokerv1.ListBrokerWhitelistedAddressesRequest
	ListBrokerWhitelistedAddressesResponse = brokerv1.ListBrokerWhitelistedAddressesResponse
	CreateBrokerWhitelistedAddressRequest  = brokerv1.CreateBrokerWhitelistedAddressRequest
	DeleteBrokerWhitelistedAddressRequest  = brokerv1.DeleteBrokerWhitelistedAddressRequest
	DeleteBrokerWhitelistedAddressResponse = brokerv1.DeleteBrokerWhitelistedAddressResponse
)

// =============================================================================
// Funding Wallet Types (v1beta)
// =============================================================================

type (
	FundingWallet                        = brokerv1beta.FundingWallet
	FundingWalletTransfer                = brokerv1beta.FundingWalletTransfer
	GetFundingWalletRequest              = brokerv1beta.GetFundingWalletRequest
	CreateFundingWalletRequest           = brokerv1beta.CreateFundingWalletRequest
	BatchCreateFundingWalletsRequest     = brokerv1beta.BatchCreateFundingWalletsRequest
	ListFundingWalletTransfersRequest    = brokerv1beta.ListFundingWalletTransfersRequest
	ListFundingWalletTransfersResponse   = brokerv1beta.ListFundingWalletTransfersResponse
	GetFundingWalletTransferRequest      = brokerv1beta.GetFundingWalletTransferRequest
	CreateFundingWalletWithdrawalRequest = brokerv1beta.CreateFundingWalletWithdrawalRequest
	ListFundingDetailsRequest            = brokerv1beta.ListFundingDetailsRequest
	DemoDepositRequest                   = brokerv1beta.DemoDepositRequest
)

// =============================================================================
// Recipient Bank Types (v1beta)
// =============================================================================

type (
	RecipientBank               = brokerv1beta.RecipientBank
	GetRecipientBankRequest     = brokerv1beta.GetRecipientBankRequest
	CreateRecipientBankRequest  = brokerv1beta.CreateRecipientBankRequest
	DeleteRecipientBankRequest  = brokerv1beta.DeleteRecipientBankRequest
	DeleteRecipientBankResponse = brokerv1beta.DeleteRecipientBankResponse
)

// =============================================================================
// FPSL (Fully Paid Securities Lending) Types
// =============================================================================

type (
	FPSLLoan              = brokerv1.FPSLLoan
	FPSLTier              = brokerv1.FPSLTier
	APRTier               = brokerv1.APRTier
	ListFPSLLoansRequest  = brokerv1.ListFPSLLoansRequest
	ListFPSLLoansResponse = brokerv1.ListFPSLLoansResponse
	ListFPSLTiersRequest  = brokerv1.ListFPSLTiersRequest
	ListFPSLTiersResponse = brokerv1.ListFPSLTiersResponse
	ListAPRTiersRequest   = brokerv1.ListAPRTiersRequest
	ListAPRTiersResponse  = brokerv1.ListAPRTiersResponse
)

// =============================================================================
// JIT (Just-In-Time) Types
// =============================================================================

type (
	JITLedger                    = brokerv1.JITLedger
	JITLedgerBalance             = brokerv1.JITLedgerBalance
	JITLimits                    = brokerv1.JITLimits
	JITSettlement                = brokerv1.JITSettlement
	ListJITLedgersRequest        = brokerv1.ListJITLedgersRequest
	ListJITLedgersResponse       = brokerv1.ListJITLedgersResponse
	GetJITLedgerBalancesRequest  = brokerv1.GetJITLedgerBalancesRequest
	GetJITLedgerBalancesResponse = brokerv1.GetJITLedgerBalancesResponse
	GetJITLimitsRequest          = brokerv1.GetJITLimitsRequest
	ListJITSettlementsRequest    = brokerv1.ListJITSettlementsRequest
	ListJITSettlementsResponse   = brokerv1.ListJITSettlementsResponse
	CreateJITSettlementRequest   = brokerv1.CreateJITSettlementRequest
	GetJITSettlementRequest      = brokerv1.GetJITSettlementRequest
)

// =============================================================================
// Country Types
// =============================================================================

type (
	CountryInfo           = brokerv1.CountryInfo
	ListCountriesRequest  = brokerv1.ListCountriesRequest
	ListCountriesResponse = brokerv1.ListCountriesResponse
)

// =============================================================================
// Reporting Types
// Note: Generic EOD report endpoints have been removed.
// Use specific reporting endpoints like GetEODPositions, GetEODCashInterest, etc.
// =============================================================================

// =============================================================================
// IRA Types
// =============================================================================

type (
	IRAExcessContribution              = brokerv1.IRAExcessContribution
	ListIRAExcessContributionsRequest  = brokerv1.ListIRAExcessContributionsRequest
	ListIRAExcessContributionsResponse = brokerv1.ListIRAExcessContributionsResponse
)

// =============================================================================
// Options Types
// =============================================================================

type (
	OptionsApproval                   = brokerv1.OptionsApproval
	GetOptionsApprovalRequest         = brokerv1.GetOptionsApprovalRequest
	RequestOptionsApprovalRequest     = brokerv1.RequestOptionsApprovalRequest
	UpdateOptionsApprovalRequest      = brokerv1.UpdateOptionsApprovalRequest
	BrokerOptionContract              = brokerv1.BrokerOptionContract
	ListBrokerOptionContractsRequest  = brokerv1.ListBrokerOptionContractsRequest
	ListBrokerOptionContractsResponse = brokerv1.ListBrokerOptionContractsResponse
	GetBrokerOptionContractRequest    = brokerv1.GetBrokerOptionContractRequest
)

// =============================================================================
// Rebalancing Types
// =============================================================================

type (
	RebalancingPortfolio       = brokerv1.RebalancingPortfolio
	PortfolioWeight            = brokerv1.PortfolioWeight
	RebalancingRun             = brokerv1.RebalancingRun
	RebalancingSubscription    = brokerv1.RebalancingSubscription
	ListPortfoliosRequest      = brokerv1.ListPortfoliosRequest
	ListPortfoliosResponse     = brokerv1.ListPortfoliosResponse
	GetPortfolioRequest        = brokerv1.GetPortfolioRequest
	CreatePortfolioRequest     = brokerv1.CreatePortfolioRequest
	UpdatePortfolioRequest     = brokerv1.UpdatePortfolioRequest
	DeletePortfolioRequest     = brokerv1.DeletePortfolioRequest
	DeletePortfolioResponse    = brokerv1.DeletePortfolioResponse
	ListRunsRequest            = brokerv1.ListRunsRequest
	ListRunsResponse           = brokerv1.ListRunsResponse
	GetRunRequest              = brokerv1.GetRunRequest
	CreateRunRequest           = brokerv1.CreateRunRequest
	CancelRunRequest           = brokerv1.CancelRunRequest
	CancelRunResponse          = brokerv1.CancelRunResponse
	ListSubscriptionsRequest   = brokerv1.ListSubscriptionsRequest
	ListSubscriptionsResponse  = brokerv1.ListSubscriptionsResponse
	GetSubscriptionRequest     = brokerv1.GetSubscriptionRequest
	CreateSubscriptionRequest  = brokerv1.CreateSubscriptionRequest
	DeleteSubscriptionRequest  = brokerv1.DeleteSubscriptionRequest
	DeleteSubscriptionResponse = brokerv1.DeleteSubscriptionResponse
)

// =============================================================================
// SSE Event Types (v1 - deprecated, use V2)
// =============================================================================

type (
	SSEEvent                       = brokerv1.SSEEvent
	AccountUpdatedEvent            = brokerv1.AccountUpdatedEvent
	TradeUpdatedEvent              = brokerv1.TradeUpdatedEvent
	TransferUpdatedEvent           = brokerv1.TransferUpdatedEvent
	JournalUpdatedEvent            = brokerv1.JournalUpdatedEvent
	NonTradingActivityEvent        = brokerv1.NonTradingActivityEvent
	SubscribeAccountEventsRequest  = brokerv1.SubscribeAccountEventsRequest
	SubscribeTradeEventsRequest    = brokerv1.SubscribeTradeEventsRequest
	SubscribeTransferEventsRequest = brokerv1.SubscribeTransferEventsRequest
	SubscribeJournalEventsRequest  = brokerv1.SubscribeJournalEventsRequest
	SubscribeNTAEventsRequest      = brokerv1.SubscribeNTAEventsRequest
)

// =============================================================================
// SSE Event Types (v2 - recommended)
// =============================================================================

type (
	TradeUpdateEventV2              = brokerv2.TradeUpdateEventV2
	TradeUpdateEventV2Leg           = brokerv2.TradeUpdateEventV2Leg
	TradeUpdateEventV2Order         = brokerv2.TradeUpdateEventV2Order
	JournalStatusEventV2            = brokerv2.JournalStatusEventV2
	SystemEventV2                   = brokerv2.SystemEventV2
	AdminActionEventV2              = brokerv2.AdminActionEventV2
	AdminActionBelongsTo            = brokerv2.AdminActionBelongsTo
	AdminActionCreatedBy            = brokerv2.AdminActionCreatedBy
	LiquidationContext              = brokerv2.LiquidationContext
	TransactionCancelContext        = brokerv2.TransactionCancelContext
	FundingStatusEventV2            = brokerv2.FundingStatusEventV2
	SubscribeTradeEventsV2Request   = brokerv2.SubscribeTradeEventsV2Request
	SubscribeJournalEventsV2Request = brokerv2.SubscribeJournalEventsV2Request
	SubscribeSystemEventsV2Request  = brokerv2.SubscribeSystemEventsV2Request
	SubscribeAdminActionsV2Request  = brokerv2.SubscribeAdminActionsV2Request
	SubscribeFundingStatusV2Request = brokerv2.SubscribeFundingStatusV2Request
)

// =============================================================================
// IPO Types
// =============================================================================

type (
	IPOOffering               = brokerv1.IPOOffering
	ListIPOOfferingsRequest   = brokerv1.ListIPOOfferingsRequest
	ListIPOOfferingsResponse  = brokerv1.ListIPOOfferingsResponse
	GetIPOOfferingRequest     = brokerv1.GetIPOOfferingRequest
	IPOEvent                  = brokerv2.IPOEvent
	IPOEventPayload           = brokerv2.IPOEventPayload
	SubscribeIPOEventsRequest = brokerv2.SubscribeIPOEventsRequest
)

// =============================================================================
// Call Option Types
// =============================================================================

type (
	V1CallOption     = brokerv1.BrokerServiceCallOption
	V1BetaCallOption = brokerv1beta.BrokerV1BetaServiceCallOption
	V2CallOption     = brokerv2.BrokerV2ServiceCallOption
)

// Per-call DiscardUnknownFields overrides.
var (
	WithV1CallDiscardUnknownFields     = brokerv1.WithBrokerServiceCallDiscardUnknownFields
	WithV1BetaCallDiscardUnknownFields = brokerv1beta.WithBrokerV1BetaServiceCallDiscardUnknownFields
	WithV2CallDiscardUnknownFields     = brokerv2.WithBrokerV2ServiceCallDiscardUnknownFields
)

// EventStream is a generic type alias for reading Server-Sent Events from v2 streaming endpoints.
// Use Next() to read events, Err() to check for errors, and Close() to release the connection.
type EventStream[T proto.Message] = brokerv2.BrokerV2ServiceEventStream[T]
