// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/public-sdk-go/internal/apijson"
	"github.com/stainless-sdks/public-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/public-sdk-go/option"
	"github.com/stainless-sdks/public-sdk-go/packages/param"
	"github.com/stainless-sdks/public-sdk-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CursorPage[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPage[T]) GetNextPage() (res *CursorPage[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorPageAutoPager[T any] struct {
	page *CursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorPageAutoPager[T any](page *CursorPage[T], err error) *CursorPageAutoPager[T] {
	return &CursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPageAutoPager[T]) Index() int {
	return r.run
}

type SearchPage[T any] struct {
	Products      []T    `json:"products"`
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Products      respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r SearchPage[T]) RawJSON() string { return r.JSON.raw }
func (r *SearchPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SearchPage[T]) GetNextPage() (res *SearchPage[T], err error) {
	if len(r.Products) == 0 {
		return nil, nil
	}
	next := r.NextPageToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("page_token", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SearchPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SearchPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SearchPageAutoPager[T any] struct {
	page *SearchPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewSearchPageAutoPager[T any](page *SearchPage[T], err error) *SearchPageAutoPager[T] {
	return &SearchPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SearchPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Products) == 0 {
		return false
	}
	if r.idx >= len(r.page.Products) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Products) == 0 {
			return false
		}
	}
	r.cur = r.page.Products[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SearchPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *SearchPageAutoPager[T]) Err() error {
	return r.err
}

func (r *SearchPageAutoPager[T]) Index() int {
	return r.run
}

type CategoryPage[T any] struct {
	Items    []T   `json:"items"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
	Total    int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CategoryPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CategoryPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CategoryPage[T]) GetNextPage() (res *CategoryPage[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	currentPage := r.Page
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CategoryPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CategoryPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CategoryPageAutoPager[T any] struct {
	page *CategoryPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCategoryPageAutoPager[T any](page *CategoryPage[T], err error) *CategoryPageAutoPager[T] {
	return &CategoryPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CategoryPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CategoryPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CategoryPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CategoryPageAutoPager[T]) Index() int {
	return r.run
}
