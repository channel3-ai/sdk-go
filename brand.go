// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/channel3-ai/sdk-go/internal/apijson"
	"github.com/channel3-ai/sdk-go/internal/apiquery"
	"github.com/channel3-ai/sdk-go/internal/requestconfig"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/channel3-ai/sdk-go/packages/pagination"
	"github.com/channel3-ai/sdk-go/packages/param"
	"github.com/channel3-ai/sdk-go/packages/respjson"
)

// BrandService contains methods and other services that help with interacting with
// the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandService] method instead.
type BrandService struct {
	options []option.RequestOption
}

// NewBrandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrandService(opts ...option.RequestOption) (r BrandService) {
	r = BrandService{}
	r.options = opts
	return
}

// Get detailed information about a specific brand by its ID.
func (r *BrandService) Get(ctx context.Context, brandID string, opts ...option.RequestOption) (res *Brand, err error) {
	opts = slices.Concat(r.options, opts)
	if brandID == "" {
		err = errors.New("missing required brand_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/brands/%s", url.PathEscape(brandID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Paginated list of brands.
func (r *BrandService) List(ctx context.Context, query BrandListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Brand], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/brands"
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

// Paginated list of brands.
func (r *BrandService) ListAutoPaging(ctx context.Context, query BrandListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Brand] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Find a brand by name.
//
// Deprecated: use `search` (returns a list) instead; will be removed in the next
// major version
func (r *BrandService) Find(ctx context.Context, query BrandFindParams, opts ...option.RequestOption) (res *Brand, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v0/brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search brands by free-text query.
func (r *BrandService) Search(ctx context.Context, query BrandSearchParams, opts ...option.RequestOption) (res *SearchBrandsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/brands/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Brand struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// The maximum commission rate for the brand, as a percentage
	BestCommissionRate float64 `json:"best_commission_rate"`
	Description        string  `json:"description" api:"nullable"`
	LogoURL            string  `json:"logo_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Name               respjson.Field
		BestCommissionRate respjson.Field
		Description        respjson.Field
		LogoURL            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Brand) RawJSON() string { return r.JSON.raw }
func (r *Brand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchBrandsResponse struct {
	// Brands matching the query, ordered by relevance.
	Brands []Brand `json:"brands" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Brands      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchBrandsResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchBrandsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandListParams struct {
	// Pagination cursor returned by a prior call. Omit for the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Max items per page (1-100).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandListParams]'s query parameters as `url.Values`.
func (r BrandListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrandFindParams struct {
	Query string `query:"query" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [BrandFindParams]'s query parameters as `url.Values`.
func (r BrandFindParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrandSearchParams struct {
	// Free-text query (e.g. 'Nike', 'lululemon').
	Query string `query:"query" api:"required" json:"-"`
	// Maximum number of brands to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BrandSearchParams]'s query parameters as `url.Values`.
func (r BrandSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
