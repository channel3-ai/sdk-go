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
// At least one of `query`, `image_url`, `base64_image`, or `page_token` must be
// provided; requests with none of these will return 422.
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
	// Deprecated: use `mode`. `true` is equivalent to `mode=keyword`.
	//
	// Deprecated: deprecated
	KeywordSearchOnly param.Opt[bool] `json:"keyword_search_only,omitzero"`
	// ISO 3166-1 alpha-2 country code. May stay unset for pan-region storefronts (e.g.
	// `currency=EUR` with no specific country).
	//
	// Any of "US", "GB", "EU", "AU", "CA", "IE", "DE", "AT", "FR", "BE", "IT", "ES",
	// "NL", "SE", "FI", "PT", "CZ", "GR", "RO".
	Country SearchConfigCountry `json:"country,omitzero"`
	// ISO 4217 currency code. When unset, inferred from `country` (e.g. `GB` → `GBP`),
	// defaulting to `USD`.
	//
	// Any of "USD", "CAD", "AUD", "GBP", "EUR", "SEK", "CZK", "RON".
	Currency SearchConfigCurrency `json:"currency,omitzero"`
	// ISO 639-1 language code. When unset, inferred from `country` (preferred) then
	// `currency`, defaulting to `en`.
	//
	// Any of "en", "de", "fr", "it", "es", "nl", "sv", "fi", "pt", "cs", "el", "ro".
	Language SearchConfigLanguage `json:"language,omitzero"`
	// Preferred unit for length dimensions (length/width/height) in responses. A
	// request dimension filter's unit for the field takes precedence; when neither is
	// set, the merchant's stated unit is returned.
	//
	// Any of "mm", "cm", "m", "in", "ft".
	LengthUnit SearchConfigLengthUnit `json:"length_unit,omitzero"`
	// Preferred unit for weight dimensions in responses. A request dimension filter's
	// weight unit takes precedence; when neither is set, the merchant's stated unit is
	// returned.
	//
	// Any of "mg", "g", "kg", "oz", "lb".
	WeightUnit SearchConfigWeightUnit `json:"weight_unit,omitzero"`
	// Search strategy. `default` (recommended) combines lexical + semantic search and
	// is right for most use cases. `keyword` is lexical only — use it for real-time,
	// low-latency needs like ad targeting. `agentic` uses an LLM to plan multiple
	// structured sub-searches for complex queries, with higher latency than the other
	// modes.
	//
	// Any of "keyword", "default", "agentic".
	Mode SearchConfigMode `json:"mode,omitzero"`
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
	SearchConfigCountryGr SearchConfigCountry = "GR"
	SearchConfigCountryRo SearchConfigCountry = "RO"
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
	SearchConfigCurrencyRon SearchConfigCurrency = "RON"
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
	SearchConfigLanguageEl SearchConfigLanguage = "el"
	SearchConfigLanguageRo SearchConfigLanguage = "ro"
)

// Preferred unit for length dimensions (length/width/height) in responses. A
// request dimension filter's unit for the field takes precedence; when neither is
// set, the merchant's stated unit is returned.
type SearchConfigLengthUnit string

const (
	SearchConfigLengthUnitMm SearchConfigLengthUnit = "mm"
	SearchConfigLengthUnitCm SearchConfigLengthUnit = "cm"
	SearchConfigLengthUnitM  SearchConfigLengthUnit = "m"
	SearchConfigLengthUnitIn SearchConfigLengthUnit = "in"
	SearchConfigLengthUnitFt SearchConfigLengthUnit = "ft"
)

// Search strategy. `default` (recommended) combines lexical + semantic search and
// is right for most use cases. `keyword` is lexical only — use it for real-time,
// low-latency needs like ad targeting. `agentic` uses an LLM to plan multiple
// structured sub-searches for complex queries, with higher latency than the other
// modes.
type SearchConfigMode string

const (
	SearchConfigModeKeyword SearchConfigMode = "keyword"
	SearchConfigModeDefault SearchConfigMode = "default"
	SearchConfigModeAgentic SearchConfigMode = "agentic"
)

// Preferred unit for weight dimensions in responses. A request dimension filter's
// weight unit takes precedence; when neither is set, the merchant's stated unit is
// returned.
type SearchConfigWeightUnit string

const (
	SearchConfigWeightUnitMg SearchConfigWeightUnit = "mg"
	SearchConfigWeightUnitG  SearchConfigWeightUnit = "g"
	SearchConfigWeightUnitKg SearchConfigWeightUnit = "kg"
	SearchConfigWeightUnitOz SearchConfigWeightUnit = "oz"
	SearchConfigWeightUnitLb SearchConfigWeightUnit = "lb"
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
	// If provided, only products whose extracted attributes match these key/value
	// constraints will be returned. Keys are attribute handles (e.g. 'color',
	// 'material') and values are lists of allowed values (OR within a key, AND across
	// keys). When a category filter is also supplied, all keys must be valid
	// attributes of at least one of the requested categories. See
	// `Category.attributes` for the valid keys/values per category.
	Attributes map[string][]string `json:"attributes,omitzero"`
	// If provided, only products with these availability statuses will be returned
	Availability []AvailabilityStatus `json:"availability,omitzero"`
	// If provided, only products from these brands will be returned
	BrandIDs []string `json:"brand_ids,omitzero"`
	// If provided, only products from these categories will be returned. Accepts
	// category slugs.
	CategoryIDs []string `json:"category_ids,omitzero"`
	// [Beta] Color filter wrapper. Holds required colors and optional match mode.
	Colors SearchFiltersColorsParam `json:"colors,omitzero"`
	// Filter by offer condition. Requires at least one offer matching the requested
	// condition, locale, and any price filter. Offers without condition data are
	// indexed as new.
	//
	// Any of "new", "refurbished", "used".
	Condition SearchFiltersCondition `json:"condition,omitzero"`
	// Physical-dimension range filters, matched against the same offer.
	//
	// Matching products have at least one offer satisfying every provided range
	// (alongside any locale/price/availability filters). Values are compared with a
	// small relative tolerance. An offer with no dimension data for a filtered field
	// does not match; note that when a single merchant on a product reports a
	// dimension it is shared across that product's offers, so a matching offer may not
	// itself surface that dimension in the response.
	Dimensions SearchFiltersDimensionsParam `json:"dimensions,omitzero"`
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
	// If 'on_sale', only products with at least one on-sale offer (priced below its
	// compare-at price) for the requested locale are returned. If omitted, no filter.
	//
	// Any of "on_sale".
	Sale SearchFiltersSale `json:"sale,omitzero"`
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

// [Beta] Color filter wrapper. Holds required colors and optional match mode.
//
// The property Palette is required.
type SearchFiltersColorsParam struct {
	// Colors required in matching products. Treated as an AND condition.
	Palette []SearchFiltersColorsPaletteParam `json:"palette,omitzero" api:"required"`
	// How tightly colors must match: 'strict', 'standard', or 'loose'.
	//
	// Any of "strict", "standard", "loose".
	Match string `json:"match,omitzero"`
	paramObj
}

func (r SearchFiltersColorsParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFiltersColorsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFiltersColorsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SearchFiltersColorsParam](
		"match", "strict", "standard", "loose",
	)
}

// A single color requirement for the color filter.
//
// The property Hex is required.
type SearchFiltersColorsPaletteParam struct {
	// sRGB hex string, e.g. '#a1b2c3'
	Hex string `json:"hex" api:"required"`
	// Percentage of color, where 1.0 is 100%
	Percentage param.Opt[float64] `json:"percentage,omitzero"`
	paramObj
}

func (r SearchFiltersColorsPaletteParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFiltersColorsPaletteParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFiltersColorsPaletteParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by offer condition. Requires at least one offer matching the requested
// condition, locale, and any price filter. Offers without condition data are
// indexed as new.
type SearchFiltersCondition string

const (
	SearchFiltersConditionNew         SearchFiltersCondition = "new"
	SearchFiltersConditionRefurbished SearchFiltersCondition = "refurbished"
	SearchFiltersConditionUsed        SearchFiltersCondition = "used"
)

// Physical-dimension range filters, matched against the same offer.
//
// Matching products have at least one offer satisfying every provided range
// (alongside any locale/price/availability filters). Values are compared with a
// small relative tolerance. An offer with no dimension data for a filtered field
// does not match; note that when a single merchant on a product reports a
// dimension it is shared across that product's offers, so a matching offer may not
// itself surface that dimension in the response.
type SearchFiltersDimensionsParam struct {
	Height SearchFiltersDimensionsHeightParam `json:"height,omitzero"`
	Length SearchFiltersDimensionsLengthParam `json:"length,omitzero"`
	Weight SearchFiltersDimensionsWeightParam `json:"weight,omitzero"`
	Width  SearchFiltersDimensionsWidthParam  `json:"width,omitzero"`
	paramObj
}

func (r SearchFiltersDimensionsParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFiltersDimensionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFiltersDimensionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Unit is required.
type SearchFiltersDimensionsHeightParam struct {
	// Unit that min/max are expressed in
	//
	// Any of "mm", "cm", "m", "in", "ft".
	Unit string `json:"unit,omitzero" api:"required"`
	// Maximum value, in `unit`. Inclusive.
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum value, in `unit`. Inclusive.
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SearchFiltersDimensionsHeightParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFiltersDimensionsHeightParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFiltersDimensionsHeightParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SearchFiltersDimensionsHeightParam](
		"unit", "mm", "cm", "m", "in", "ft",
	)
}

// The property Unit is required.
type SearchFiltersDimensionsLengthParam struct {
	// Unit that min/max are expressed in
	//
	// Any of "mm", "cm", "m", "in", "ft".
	Unit string `json:"unit,omitzero" api:"required"`
	// Maximum value, in `unit`. Inclusive.
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum value, in `unit`. Inclusive.
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SearchFiltersDimensionsLengthParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFiltersDimensionsLengthParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFiltersDimensionsLengthParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SearchFiltersDimensionsLengthParam](
		"unit", "mm", "cm", "m", "in", "ft",
	)
}

// The property Unit is required.
type SearchFiltersDimensionsWeightParam struct {
	// Unit that min/max are expressed in
	//
	// Any of "mg", "g", "kg", "oz", "lb".
	Unit string `json:"unit,omitzero" api:"required"`
	// Maximum value, in `unit`. Inclusive.
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum value, in `unit`. Inclusive.
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SearchFiltersDimensionsWeightParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFiltersDimensionsWeightParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFiltersDimensionsWeightParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SearchFiltersDimensionsWeightParam](
		"unit", "mg", "g", "kg", "oz", "lb",
	)
}

// The property Unit is required.
type SearchFiltersDimensionsWidthParam struct {
	// Unit that min/max are expressed in
	//
	// Any of "mm", "cm", "m", "in", "ft".
	Unit string `json:"unit,omitzero" api:"required"`
	// Maximum value, in `unit`. Inclusive.
	Max param.Opt[float64] `json:"max,omitzero"`
	// Minimum value, in `unit`. Inclusive.
	Min param.Opt[float64] `json:"min,omitzero"`
	paramObj
}

func (r SearchFiltersDimensionsWidthParam) MarshalJSON() (data []byte, err error) {
	type shadow SearchFiltersDimensionsWidthParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFiltersDimensionsWidthParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SearchFiltersDimensionsWidthParam](
		"unit", "mm", "cm", "m", "in", "ft",
	)
}

type SearchFiltersGender string

const (
	SearchFiltersGenderMale   SearchFiltersGender = "male"
	SearchFiltersGenderFemale SearchFiltersGender = "female"
)

// If 'on_sale', only products with at least one on-sale offer (priced below its
// compare-at price) for the requested locale are returned. If omitted, no filter.
type SearchFiltersSale string

const (
	SearchFiltersSaleOnSale SearchFiltersSale = "on_sale"
)

// Search request with pagination support.
type SearchRequestParam struct {
	// Base64 encoded image. At least one of `query`, `image_url`, `base64_image`, or
	// `page_token` must be provided.
	Base64Image param.Opt[string] `json:"base64_image,omitzero"`
	// Image URL. At least one of `query`, `image_url`, `base64_image`, or `page_token`
	// must be provided.
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Optional limit on the number of results. Default is 20, max is 30.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Opaque token from a previous search response to fetch the next page of results.
	PageToken param.Opt[string] `json:"page_token,omitzero"`
	// Search query. At least one of `query`, `image_url`, `base64_image`, or
	// `page_token` must be provided.
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
