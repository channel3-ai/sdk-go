// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"context"
	"net/http"
	"slices"

	"github.com/channel3-ai/sdk-go/internal/apijson"
	shimjson "github.com/channel3-ai/sdk-go/internal/encoding/json"
	"github.com/channel3-ai/sdk-go/internal/requestconfig"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/channel3-ai/sdk-go/packages/param"
	"github.com/channel3-ai/sdk-go/packages/respjson"
)

// SearchService contains methods and other services that help with interacting
// with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
//
// Deprecated: use `products.search` (or `products.search_by_image` /
// `products.find_similar`) instead; this resource will be removed in the next
// major version
type SearchService struct {
	options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.options = opts
	return
}

// Search for products with pagination support.
//
// At least one of `query`, `image_url`, or `base64_image` must be provided;
// requests with none of these will return 422.
//
// Deprecated: use `products.search` instead, which auto-paginates; will be removed
// in the next major version
func (r *SearchService) Perform(ctx context.Context, body SearchPerformParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Search and locale options for a search request.
type SearchConfigParam struct {
	// If True, search will only use keyword search and not vector search. Keyword-only
	// search is not supported with image input.
	KeywordSearchOnly param.Opt[bool] `json:"keyword_search_only,omitzero"`
	// ISO 3166-1 alpha-2 country code. May stay unset for pan-region storefronts (e.g.
	// `currency=EUR` with no specific country).
	//
	// Any of "US", "GB", "EU", "AU", "CA", "IE", "DE", "AT", "FR", "BE", "IT", "ES",
	// "NL", "SE", "FI", "PT", "CZ".
	Country SearchConfigCountry `json:"country,omitzero"`
	// ISO 4217 currency code. When unset, inferred from `country` (e.g. `GB` → `GBP`),
	// defaulting to `USD`.
	//
	// Any of "USD", "CAD", "AUD", "GBP", "EUR", "SEK", "CZK".
	Currency SearchConfigCurrency `json:"currency,omitzero"`
	// ISO 639-1 language code. When unset, inferred from `country` (preferred) then
	// `currency`, defaulting to `en`.
	//
	// Any of "en", "de", "fr", "it", "es", "nl", "sv", "fi", "pt", "cs".
	Language SearchConfigLanguage `json:"language,omitzero"`
	paramObj
}

func (r SearchConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ISO 3166-1 alpha-2 country code. May stay unset for pan-region storefronts (e.g.
// `currency=EUR` with no specific country).
type SearchConfigCountry string

const (
	SearchConfigCountryUs SearchConfigCountry = "US"
	SearchConfigCountryGB SearchConfigCountry = "GB"
	SearchConfigCountryEu SearchConfigCountry = "EU"
	SearchConfigCountryAu SearchConfigCountry = "AU"
	SearchConfigCountryCa SearchConfigCountry = "CA"
	SearchConfigCountryIe SearchConfigCountry = "IE"
	SearchConfigCountryDe SearchConfigCountry = "DE"
	SearchConfigCountryAt SearchConfigCountry = "AT"
	SearchConfigCountryFr SearchConfigCountry = "FR"
	SearchConfigCountryBe SearchConfigCountry = "BE"
	SearchConfigCountryIt SearchConfigCountry = "IT"
	SearchConfigCountryEs SearchConfigCountry = "ES"
	SearchConfigCountryNl SearchConfigCountry = "NL"
	SearchConfigCountrySe SearchConfigCountry = "SE"
	SearchConfigCountryFi SearchConfigCountry = "FI"
	SearchConfigCountryPt SearchConfigCountry = "PT"
	SearchConfigCountryCz SearchConfigCountry = "CZ"
)

// ISO 4217 currency code. When unset, inferred from `country` (e.g. `GB` → `GBP`),
// defaulting to `USD`.
type SearchConfigCurrency string

const (
	SearchConfigCurrencyUsd SearchConfigCurrency = "USD"
	SearchConfigCurrencyCad SearchConfigCurrency = "CAD"
	SearchConfigCurrencyAud SearchConfigCurrency = "AUD"
	SearchConfigCurrencyGbp SearchConfigCurrency = "GBP"
	SearchConfigCurrencyEur SearchConfigCurrency = "EUR"
	SearchConfigCurrencySek SearchConfigCurrency = "SEK"
	SearchConfigCurrencyCzk SearchConfigCurrency = "CZK"
)

// ISO 639-1 language code. When unset, inferred from `country` (preferred) then
// `currency`, defaulting to `en`.
type SearchConfigLanguage string

const (
	SearchConfigLanguageEn SearchConfigLanguage = "en"
	SearchConfigLanguageDe SearchConfigLanguage = "de"
	SearchConfigLanguageFr SearchConfigLanguage = "fr"
	SearchConfigLanguageIt SearchConfigLanguage = "it"
	SearchConfigLanguageEs SearchConfigLanguage = "es"
	SearchConfigLanguageNl SearchConfigLanguage = "nl"
	SearchConfigLanguageSv SearchConfigLanguage = "sv"
	SearchConfigLanguageFi SearchConfigLanguage = "fi"
	SearchConfigLanguagePt SearchConfigLanguage = "pt"
	SearchConfigLanguageCs SearchConfigLanguage = "cs"
)

// Price filter for search. Values are inclusive.
type SearchFilterPriceParam struct {
	// Maximum price, in dollars and cents
	MaxPrice param.Opt[float64] `json:"max_price,omitzero"`
	// Minimum price, in dollars and cents
	MinPrice param.Opt[float64] `json:"min_price,omitzero"`
	paramObj
}

func (r SearchFilterPriceParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFilterPriceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFilterPriceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search filters for the search API.
type SearchFiltersParam struct {
	// Filter by age group. Age-agnostic products are treated as adult products.
	//
	// Any of "newborn", "infant", "toddler", "kids", "adult".
	Age []string `json:"age,omitzero"`
	// If provided, only products with these availability statuses will be returned
	Availability []AvailabilityStatus `json:"availability,omitzero"`
	// If provided, only products from these brands will be returned
	BrandIDs []string `json:"brand_ids,omitzero"`
	// If provided, only products from these categories will be returned. Accepts
	// category slugs.
	CategoryIDs []string `json:"category_ids,omitzero"`
	// Filter by product condition. Incubating: condition data is currently incomplete;
	// products without condition data will be included in all condition filter
	// results.
	//
	// Any of "new", "refurbished", "used".
	Condition SearchFiltersCondition `json:"condition,omitzero"`
	// If provided, products from these brands will be excluded from the results
	ExcludeBrandIDs []string `json:"exclude_brand_ids,omitzero"`
	// If provided, products in these categories (or their descendants) will be
	// excluded from the results. Accepts category slugs.
	ExcludeCategoryIDs []string `json:"exclude_category_ids,omitzero"`
	// If provided, products from these websites will be excluded from the results.
	// Accepts website IDs or domains (e.g. "nike.com").
	ExcludeWebsiteIDs []string `json:"exclude_website_ids,omitzero"`
	// Any of "male", "female".
	Gender SearchFiltersGender `json:"gender,omitzero"`
	// If provided, only products from these websites will be returned. Accepts website
	// IDs or domains (e.g. "nike.com").
	WebsiteIDs []string `json:"website_ids,omitzero"`
	// Price filter for search. Values are inclusive.
	Price SearchFilterPriceParam `json:"price,omitzero"`
	paramObj
}

func (r SearchFiltersParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFiltersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFiltersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by product condition. Incubating: condition data is currently incomplete;
// products without condition data will be included in all condition filter
// results.
type SearchFiltersCondition string

const (
	SearchFiltersConditionNew         SearchFiltersCondition = "new"
	SearchFiltersConditionRefurbished SearchFiltersCondition = "refurbished"
	SearchFiltersConditionUsed        SearchFiltersCondition = "used"
)

type SearchFiltersGender string

const (
	SearchFiltersGenderMale   SearchFiltersGender = "male"
	SearchFiltersGenderFemale SearchFiltersGender = "female"
)

// Search request with pagination support.
type SearchRequestParam struct {
	// Base64 encoded image. At least one of `query`, `image_url`, or `base64_image`
	// must be provided.
	Base64Image param.Opt[string] `json:"base64_image,omitzero"`
	// Image URL. At least one of `query`, `image_url`, or `base64_image` must be
	// provided.
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Optional limit on the number of results. Default is 20, max is 30.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Opaque token from a previous search response to fetch the next page of results.
	PageToken param.Opt[string] `json:"page_token,omitzero"`
	// Search query. At least one of `query`, `image_url`, or `base64_image` must be
	// provided.
	Query param.Opt[string] `json:"query,omitzero"`
	// Optional configuration
	Config SearchConfigParam `json:"config,omitzero"`
	// Optional filters. Search will only consider products that match all of the
	// filters.
	Filters SearchFiltersParam `json:"filters,omitzero"`
	paramObj
}

func (r SearchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// v1 paginated search response.
type SearchResponse struct {
	Products []ProductDetail `json:"products" api:"required"`
	// Token to fetch the next page. Null when no more results.
	NextPageToken string `json:"next_page_token" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Products      respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchPerformParams struct {
	// Search request with pagination support.
	SearchRequest SearchRequestParam
	paramObj
}

func (r SearchPerformParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SearchRequest)
}
func (r *SearchPerformParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
