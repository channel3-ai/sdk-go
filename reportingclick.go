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

// ReportingClickService contains methods and other services that help with
// interacting with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportingClickService] method instead.
type ReportingClickService struct {
	options []option.RequestOption
}

// NewReportingClickService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReportingClickService(opts ...option.RequestOption) (r ReportingClickService) {
	r = ReportingClickService{}
	r.options = opts
	return
}

// List clicks for your account over a datetime window.
//
// Defaults to the last 30 days ending now. Maximum window is 90 days. Pass an
// offset-aware ISO datetime to express local time (e.g. last 6 hours). Returns a
// summary plus a paginated list of click events (most recent first).
func (r *ReportingClickService) List(ctx context.Context, query ReportingClickListParams, opts ...option.RequestOption) (res *pagination.AnalyticsPage[Click], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/reporting/clicks"
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

// List clicks for your account over a datetime window.
//
// Defaults to the last 30 days ending now. Maximum window is 90 days. Pass an
// offset-aware ISO datetime to express local time (e.g. last 6 hours). Returns a
// summary plus a paginated list of click events (most recent first).
func (r *ReportingClickService) ListAutoPaging(ctx context.Context, query ReportingClickListParams, opts ...option.RequestOption) *pagination.AnalyticsPageAutoPager[Click] {
	return pagination.NewAnalyticsPageAutoPager(r.List(ctx, query, opts...))
}

// A single click event.
type Click struct {
	// Click event ID.
	ID string `json:"id" api:"required"`
	// When the click occurred, returned with a UTC offset (Z).
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Click city, if available.
	City string `json:"city" api:"nullable"`
	// Click country, if available.
	Country string `json:"country" api:"nullable"`
	// Compact product reference on click/transaction items.
	Product ReportingProduct `json:"product" api:"nullable"`
	// Partner-supplied user identifier from the buy URL, if provided.
	UserID string `json:"user_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Timestamp   respjson.Field
		City        respjson.Field
		Country     respjson.Field
		Product     respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Click) RawJSON() string { return r.JSON.raw }
func (r *Click) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated clicks for a vendor over a date range.
type ClicksResponse struct {
	// Inclusive end of the resolved query window.
	EndDate time.Time `json:"end_date" api:"required" format:"date-time"`
	// Whether more pages are available.
	HasMore bool    `json:"has_more" api:"required"`
	Items   []Click `json:"items" api:"required"`
	// Page size.
	Limit int64 `json:"limit" api:"required"`
	// Current page (1-indexed).
	Page int64 `json:"page" api:"required"`
	// Inclusive start of the resolved query window.
	StartDate time.Time `json:"start_date" api:"required" format:"date-time"`
	// Aggregate click stats for the requested date range.
	Summary ClicksSummary `json:"summary" api:"required"`
	// Total matching clicks in the date range.
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
func (r ClicksResponse) RawJSON() string { return r.JSON.raw }
func (r *ClicksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aggregate click stats for the requested date range.
type ClicksSummary struct {
	// Total clicks in the date range.
	TotalClicks int64 `json:"total_clicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalClicks respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClicksSummary) RawJSON() string { return r.JSON.raw }
func (r *ClicksSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportingClickListParams struct {
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

// URLQuery serializes [ReportingClickListParams]'s query parameters as
// `url.Values`.
func (r ReportingClickListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
