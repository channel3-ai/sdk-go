// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/channel3-ai/sdk-go/internal/apijson"
	"github.com/channel3-ai/sdk-go/internal/apiquery"
	shimjson "github.com/channel3-ai/sdk-go/internal/encoding/json"
	"github.com/channel3-ai/sdk-go/internal/requestconfig"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/channel3-ai/sdk-go/packages/pagination"
	"github.com/channel3-ai/sdk-go/packages/param"
	"github.com/channel3-ai/sdk-go/packages/respjson"
)

// PriceTrackingService contains methods and other services that help with
// interacting with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPriceTrackingService] method instead.
type PriceTrackingService struct {
	options []option.RequestOption
}

// NewPriceTrackingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPriceTrackingService(opts ...option.RequestOption) (r PriceTrackingService) {
	r = PriceTrackingService{}
	r.options = opts
	return
}

// List your active price tracking subscriptions.
func (r *PriceTrackingService) ListSubscriptions(ctx context.Context, query PriceTrackingListSubscriptionsParams, opts ...option.RequestOption) (res *pagination.CursorPage[Subscription], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v0/price-tracking/subscriptions"
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

// List your active price tracking subscriptions.
func (r *PriceTrackingService) ListSubscriptionsAutoPaging(ctx context.Context, query PriceTrackingListSubscriptionsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Subscription] {
	return pagination.NewCursorPageAutoPager(r.ListSubscriptions(ctx, query, opts...))
}

// Get price history for a canonical product.
func (r *PriceTrackingService) GetHistory(ctx context.Context, canonicalProductID string, query PriceTrackingGetHistoryParams, opts ...option.RequestOption) (res *PriceHistory, err error) {
	opts = slices.Concat(r.options, opts)
	if canonicalProductID == "" {
		err = errors.New("missing required canonical_product_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/price-tracking/history/%s", url.PathEscape(canonicalProductID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Start tracking prices for a canonical product.
func (r *PriceTrackingService) Start(ctx context.Context, body PriceTrackingStartParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v0/price-tracking/start"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop tracking prices for a canonical product.
func (r *PriceTrackingService) Stop(ctx context.Context, body PriceTrackingStopParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v0/price-tracking/stop"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PaginatedSubscriptionsResponse struct {
	Items      []Subscription `json:"items" api:"required"`
	NextCursor string         `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedSubscriptionsResponse) RawJSON() string { return r.JSON.raw }
func (r *PaginatedSubscriptionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceHistory struct {
	CanonicalProductID string              `json:"canonical_product_id" api:"required"`
	History            []PriceHistoryPoint `json:"history"`
	ProductTitle       string              `json:"product_title" api:"nullable"`
	Statistics         PriceStatistics     `json:"statistics" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanonicalProductID respjson.Field
		History            respjson.Field
		ProductTitle       respjson.Field
		Statistics         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceHistory) RawJSON() string { return r.JSON.raw }
func (r *PriceHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceHistoryPoint struct {
	Currency  string    `json:"currency" api:"required"`
	Price     float64   `json:"price" api:"required"`
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency    respjson.Field
		Price       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceHistoryPoint) RawJSON() string { return r.JSON.raw }
func (r *PriceHistoryPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceStatistics struct {
	Currency     string  `json:"currency" api:"required"`
	CurrentPrice float64 `json:"current_price" api:"required"`
	// Any of "low", "typical", "high".
	CurrentStatus PriceStatisticsCurrentStatus `json:"current_status" api:"required"`
	MaxPrice      float64                      `json:"max_price" api:"required"`
	Mean          float64                      `json:"mean" api:"required"`
	MinPrice      float64                      `json:"min_price" api:"required"`
	StdDev        float64                      `json:"std_dev" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency      respjson.Field
		CurrentPrice  respjson.Field
		CurrentStatus respjson.Field
		MaxPrice      respjson.Field
		Mean          respjson.Field
		MinPrice      respjson.Field
		StdDev        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PriceStatistics) RawJSON() string { return r.JSON.raw }
func (r *PriceStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceStatisticsCurrentStatus string

const (
	PriceStatisticsCurrentStatusLow     PriceStatisticsCurrentStatus = "low"
	PriceStatisticsCurrentStatusTypical PriceStatisticsCurrentStatus = "typical"
	PriceStatisticsCurrentStatusHigh    PriceStatisticsCurrentStatus = "high"
)

// The property CanonicalProductID is required.
type StartTrackingRequestParam struct {
	CanonicalProductID string `json:"canonical_product_id" api:"required"`
	paramObj
}

func (r StartTrackingRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow StartTrackingRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StartTrackingRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CanonicalProductID is required.
type StopTrackingRequestParam struct {
	CanonicalProductID string `json:"canonical_product_id" api:"required"`
	paramObj
}

func (r StopTrackingRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow StopTrackingRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StopTrackingRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Subscription struct {
	CanonicalProductID string    `json:"canonical_product_id" api:"required"`
	CreatedAt          time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "active", "cancelled".
	SubscriptionStatus SubscriptionSubscriptionStatus `json:"subscription_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanonicalProductID respjson.Field
		CreatedAt          respjson.Field
		SubscriptionStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subscription) RawJSON() string { return r.JSON.raw }
func (r *Subscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubscriptionSubscriptionStatus string

const (
	SubscriptionSubscriptionStatusActive    SubscriptionSubscriptionStatus = "active"
	SubscriptionSubscriptionStatusCancelled SubscriptionSubscriptionStatus = "cancelled"
)

type PriceTrackingListSubscriptionsParams struct {
	// Pagination cursor
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Max results (1-100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PriceTrackingListSubscriptionsParams]'s query parameters as
// `url.Values`.
func (r PriceTrackingListSubscriptionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PriceTrackingGetHistoryParams struct {
	// Number of days of history to fetch (max 30)
	Days param.Opt[int64] `query:"days,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PriceTrackingGetHistoryParams]'s query parameters as
// `url.Values`.
func (r PriceTrackingGetHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PriceTrackingStartParams struct {
	StartTrackingRequest StartTrackingRequestParam
	paramObj
}

func (r PriceTrackingStartParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.StartTrackingRequest)
}
func (r *PriceTrackingStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PriceTrackingStopParams struct {
	StopTrackingRequest StopTrackingRequestParam
	paramObj
}

func (r PriceTrackingStopParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.StopTrackingRequest)
}
func (r *PriceTrackingStopParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
