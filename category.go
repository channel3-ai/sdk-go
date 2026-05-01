// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/public-sdk-go/internal/apijson"
	"github.com/stainless-sdks/public-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/public-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/public-sdk-go/option"
	"github.com/stainless-sdks/public-sdk-go/packages/pagination"
	"github.com/stainless-sdks/public-sdk-go/packages/param"
	"github.com/stainless-sdks/public-sdk-go/packages/respjson"
)

// CategoryService contains methods and other services that help with interacting
// with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCategoryService] method instead.
type CategoryService struct {
	options []option.RequestOption
}

// NewCategoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCategoryService(opts ...option.RequestOption) (r CategoryService) {
	r = CategoryService{}
	r.options = opts
	return
}

// Look up a category by slug.
func (r *CategoryService) Get(ctx context.Context, slug string, opts ...option.RequestOption) (res *Category, err error) {
	opts = slices.Concat(r.options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/categories/%s", url.PathEscape(slug))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Paginated list of all categories.
func (r *CategoryService) List(ctx context.Context, query CategoryListParams, opts ...option.RequestOption) (res *pagination.CategoryPage[CategorySummary], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/categories"
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

// Paginated list of all categories.
func (r *CategoryService) ListAutoPaging(ctx context.Context, query CategoryListParams, opts ...option.RequestOption) *pagination.CategoryPageAutoPager[CategorySummary] {
	return pagination.NewCategoryPageAutoPager(r.List(ctx, query, opts...))
}

// Search categories by free-text query.
func (r *CategoryService) Search(ctx context.Context, query CategorySearchParams, opts ...option.RequestOption) (res *SearchCategoriesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/categories/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Full category representation used in detail responses.
type Category struct {
	// Whether this category has subcategories
	HasChildren bool `json:"has_children" api:"required"`
	// URL-friendly slug (e.g. 'sofas')
	Slug string `json:"slug" api:"required"`
	// Human-readable category title
	Title string `json:"title" api:"required"`
	// Structured attributes applicable to this category
	Attributes []CategoryAttribute `json:"attributes"`
	// Direct subcategories only (one level)
	Children []CategoryRef `json:"children"`
	// Natural-language description of products in this category
	Description string `json:"description" api:"nullable"`
	// Hierarchical path as a structured list, root first; the last entry is this
	// category itself
	Path []CategoryRef `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasChildren respjson.Field
		Slug        respjson.Field
		Title       respjson.Field
		Attributes  respjson.Field
		Children    respjson.Field
		Description respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Category) RawJSON() string { return r.JSON.raw }
func (r *Category) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A structured attribute (e.g. Color, Size) applicable to a category.
type CategoryAttribute struct {
	// Human-readable attribute name (e.g. 'Color')
	Name string `json:"name" api:"required"`
	// URL-friendly identifier (e.g. 'color', 'frame-color')
	Slug string `json:"slug" api:"required"`
	// Allowed values for this attribute. May be empty when no enumerated value set is
	// defined.
	Values []string `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Slug        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CategoryAttribute) RawJSON() string { return r.JSON.raw }
func (r *CategoryAttribute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lean reference to a category, used in path and children arrays.
type CategoryRef struct {
	// URL-friendly slug (e.g. 'sofas')
	Slug string `json:"slug" api:"required"`
	// Human-readable category title
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Slug        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CategoryRef) RawJSON() string { return r.JSON.raw }
func (r *CategoryRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lean category representation used in search hits and list rows.
type CategorySummary struct {
	// Whether this category has subcategories
	HasChildren bool `json:"has_children" api:"required"`
	// URL-friendly slug (e.g. 'sofas')
	Slug string `json:"slug" api:"required"`
	// Human-readable category title
	Title string `json:"title" api:"required"`
	// Hierarchical path as a structured list, root first; the last entry is this
	// category itself
	Path []CategoryRef `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasChildren respjson.Field
		Slug        respjson.Field
		Title       respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CategorySummary) RawJSON() string { return r.JSON.raw }
func (r *CategorySummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaginatedListCategoriesResponse struct {
	// Categories in this page
	Items []CategorySummary `json:"items" api:"required"`
	// 1-indexed page number returned
	Page int64 `json:"page" api:"required"`
	// Number of items per page
	PageSize int64 `json:"page_size" api:"required"`
	// Total number of categories matching the query
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedListCategoriesResponse) RawJSON() string { return r.JSON.raw }
func (r *PaginatedListCategoriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchCategoriesResponse struct {
	// Categories matching the query, ordered by relevance.
	Categories []CategorySummary `json:"categories" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchCategoriesResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchCategoriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CategoryListParams struct {
	// 1-indexed page number.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Items per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// If true, return only top-level (root) categories.
	RootsOnly param.Opt[bool] `query:"roots_only,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CategoryListParams]'s query parameters as `url.Values`.
func (r CategoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CategorySearchParams struct {
	// Free-text query (e.g. 'sofas', 'yoga mats').
	Query string `query:"query" api:"required" json:"-"`
	// Maximum number of categories to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CategorySearchParams]'s query parameters as `url.Values`.
func (r CategorySearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
