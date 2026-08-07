// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/channel3-ai/sdk-go/internal/apijson"
	"github.com/channel3-ai/sdk-go/internal/apiquery"
	"github.com/channel3-ai/sdk-go/internal/requestconfig"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/channel3-ai/sdk-go/packages/pagination"
	"github.com/channel3-ai/sdk-go/packages/param"
	"github.com/channel3-ai/sdk-go/packages/respjson"
)

// ReportingTransactionService contains methods and other services that help with
// interacting with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportingTransactionService] method instead.
type ReportingTransactionService struct {
	options []option.RequestOption
}

// NewReportingTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReportingTransactionService(opts ...option.RequestOption) (r ReportingTransactionService) {
	r = ReportingTransactionService{}
	r.options = opts
	return
}

// List transactions for your account over a datetime window.
//
// Defaults to the last 30 days ending now. Maximum window is 90 days. Pass an
// offset-aware ISO datetime to express local time (e.g. last 6 hours). Returns a
// summary of net commission (after take rate) plus a paginated list of
// transactions (most recent first). Network-approved commissions appear as
// pending.
func (r *ReportingTransactionService) List(ctx context.Context, query ReportingTransactionListParams, opts ...option.RequestOption) (res *pagination.AnalyticsPage[Transaction], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/reporting/transactions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List transactions for your account over a datetime window.
//
// Defaults to the last 30 days ending now. Maximum window is 90 days. Pass an
// offset-aware ISO datetime to express local time (e.g. last 6 hours). Returns a
// summary of net commission (after take rate) plus a paginated list of
// transactions (most recent first). Network-approved commissions appear as
// pending.
func (r *ReportingTransactionService) ListAutoPaging(ctx context.Context, query ReportingTransactionListParams, opts ...option.RequestOption) *pagination.AnalyticsPageAutoPager[Transaction] {
	return pagination.NewAnalyticsPageAutoPager(r.List(ctx, query, opts...))
}

// Vendor-facing transaction status (approved is surfaced as pending).
type PublicTransactionStatus string

const (
	PublicTransactionStatusPending PublicTransactionStatus = "pending"
	PublicTransactionStatusPaid    PublicTransactionStatus = "paid"
)

// A single CPA transaction.
type Transaction struct {
	// Transaction ID.
	ID string `json:"id" api:"required"`
	// Vendor net commission (after Channel3 take rate).
	CommissionAmount float64 `json:"commission_amount" api:"required"`
	// Order amount in the transaction currency.
	OrderAmount float64 `json:"order_amount" api:"required"`
	// Purchase timestamp, returned with a UTC offset (Z).
	PurchasedAt time.Time `json:"purchased_at" api:"required" format:"date-time"`
	// pending (includes network-approved) or paid.
	//
	// Any of "pending", "paid".
	Status PublicTransactionStatus `json:"status" api:"required"`
	// Brand name, if known.
	BrandName string `json:"brand_name" api:"nullable"`
	// Purchase city, if available.
	City string `json:"city" api:"nullable"`
	// Purchase country, if available.
	Country string `json:"country" api:"nullable"`
	// Compact product reference on click/transaction items.
	Product ReportingProduct `json:"product" api:"nullable"`
	// Partner-supplied user identifier from the originating click, if provided.
	UserID string `json:"user_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CommissionAmount respjson.Field
		OrderAmount      respjson.Field
		PurchasedAt      respjson.Field
		Status           respjson.Field
		BrandName        respjson.Field
		City             respjson.Field
		Country          respjson.Field
		Product          respjson.Field
		UserID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Transaction) RawJSON() string { return r.JSON.raw }
func (r *Transaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated transactions for a vendor over a date range.
type TransactionsResponse struct {
	// Inclusive end of the resolved query window. Always returned with a UTC offset
	// (Z); request values with other offsets are converted.
	EndDate time.Time `json:"end_date" api:"required" format:"date-time"`
	// Whether more pages are available.
	HasMore bool          `json:"has_more" api:"required"`
	Items   []Transaction `json:"items" api:"required"`
	// Page size.
	Limit int64 `json:"limit" api:"required"`
	// Current page (1-indexed).
	Page int64 `json:"page" api:"required"`
	// Inclusive start of the resolved query window. Always returned with a UTC offset
	// (Z); request values with other offsets are converted.
	StartDate time.Time `json:"start_date" api:"required" format:"date-time"`
	// Aggregate transaction stats for the requested date range.
	Summary TransactionsSummary `json:"summary" api:"required"`
	// Total matching transactions in the date range.
	TotalCount int64 `json:"total_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndDate     respjson.Field
		HasMore     respjson.Field
		Items       respjson.Field
		Limit       respjson.Field
		Page        respjson.Field
		StartDate   respjson.Field
		Summary     respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionsResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aggregate transaction stats for the requested date range.
type TransactionsSummary struct {
	// Vendor net commission already paid out.
	PaidCommission float64 `json:"paid_commission" api:"required"`
	// Vendor net commission still pending payout.
	PendingCommission float64 `json:"pending_commission" api:"required"`
	// Sum of vendor net commission (pending + paid) for the date range.
	TotalCommission float64 `json:"total_commission" api:"required"`
	// Total transactions in the date range.
	TotalCount int64 `json:"total_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PaidCommission    respjson.Field
		PendingCommission respjson.Field
		TotalCommission   respjson.Field
		TotalCount        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionsSummary) RawJSON() string { return r.JSON.raw }
func (r *TransactionsSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportingTransactionListParams struct {
	// Inclusive end of the window (ISO 8601 datetime with optional offset, e.g.
	// 2026-08-01T23:59:59-04:00). Offset-aware values are converted to UTC; naive
	// values are treated as UTC.
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date-time" json:"-"`
	// Inclusive start of the window (ISO 8601 datetime with optional offset, e.g.
	// 2026-08-01T00:00:00-04:00). Offset-aware values are converted to UTC; naive
	// values are treated as UTC.
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date-time" json:"-"`
	// Filter results to clicks or transactions for this user.
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// Items per page (max 100).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number (1-indexed).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportingTransactionListParams]'s query parameters as
// `url.Values`.
func (r ReportingTransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
