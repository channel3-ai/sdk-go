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
	shimjson "github.com/channel3-ai/sdk-go/internal/encoding/json"
	"github.com/channel3-ai/sdk-go/internal/requestconfig"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/channel3-ai/sdk-go/packages/pagination"
	"github.com/channel3-ai/sdk-go/packages/param"
	"github.com/channel3-ai/sdk-go/packages/respjson"
)

// ProductService contains methods and other services that help with interacting
// with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	options []option.RequestOption
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r ProductService) {
	r = ProductService{}
	r.options = opts
	return
}

// Get detailed information about a specific product by its ID.
func (r *ProductService) Get(ctx context.Context, productID string, params ProductGetParams, opts ...option.RequestOption) (res *ProductDetail, err error) {
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/products/%s", url.PathEscape(productID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List and page through products for a set of filters.
//
// Useful for a static, grid view of products for a brand, website, or category.
//
// At least one of `filters.brand_ids`, `filters.category_ids`, or
// `filters.website_ids` must be provided.
func (r *ProductService) Browse(ctx context.Context, params ProductBrowseParams, opts ...option.RequestOption) (res *pagination.SearchPage[ProductDetail], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/browse"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// List and page through products for a set of filters.
//
// Useful for a static, grid view of products for a brand, website, or category.
//
// At least one of `filters.brand_ids`, `filters.category_ids`, or
// `filters.website_ids` must be provided.
func (r *ProductService) BrowseAutoPaging(ctx context.Context, params ProductBrowseParams, opts ...option.RequestOption) *pagination.SearchPageAutoPager[ProductDetail] {
	return pagination.NewSearchPageAutoPager(r.Browse(ctx, params, opts...))
}

// Find products similar to a given product.
//
// Consider setting `filters` to narrow results to the same gender, brand,
// category, price range, etc. when you only want similar items within a specific
// slice of the catalog.
func (r *ProductService) FindSimilar(ctx context.Context, params ProductFindSimilarParams, opts ...option.RequestOption) (res *pagination.SearchPage[ProductDetail], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/similar"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Find products similar to a given product.
//
// Consider setting `filters` to narrow results to the same gender, brand,
// category, price range, etc. when you only want similar items within a specific
// slice of the catalog.
func (r *ProductService) FindSimilarAutoPaging(ctx context.Context, params ProductFindSimilarParams, opts ...option.RequestOption) *pagination.SearchPageAutoPager[ProductDetail] {
	return pagination.NewSearchPageAutoPager(r.FindSimilar(ctx, params, opts...))
}

// Retrieve product information for any supported product URL.
//
// Returns the same Product model as GET /v1/products/{product_id}. The product_id
// in the response can be used with the Product Detail endpoint.
func (r *ProductService) Lookup(ctx context.Context, params ProductLookupParams, opts ...option.RequestOption) (res *LookupResponse, err error) {
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Return monetizable offers (with max commission rate) for a product URL.
//
// Access to this endpoint is restricted. If you think your use-case requires it,
// please contact us. Usually, developers actually want search. This is helpful for
// migrating to Channel3.
func (r *ProductService) Monetize(ctx context.Context, params ProductMonetizeParams, opts ...option.RequestOption) (res *MonetizeResponse, err error) {
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/monetize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Search for products with pagination support.
//
// At least one of `query`, `image_url`, `base64_image`, or `page_token` must be
// provided; requests with none of these will return 422.
func (r *ProductService) Search(ctx context.Context, params ProductSearchParams, opts ...option.RequestOption) (res *pagination.SearchPage[ProductDetail], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/search"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Search for products with pagination support.
//
// At least one of `query`, `image_url`, `base64_image`, or `page_token` must be
// provided; requests with none of these will return 422.
func (r *ProductService) SearchAutoPaging(ctx context.Context, params ProductSearchParams, opts ...option.RequestOption) *pagination.SearchPageAutoPager[ProductDetail] {
	return pagination.NewSearchPageAutoPager(r.Search(ctx, params, opts...))
}

// Search the catalog by image (URL or base64), with pagination support.
//
// Provide exactly one of `image_url` or `base64_image`. For text or text+image
// search, use `POST /v1/search`.
func (r *ProductService) SearchByImage(ctx context.Context, params ProductSearchByImageParams, opts ...option.RequestOption) (res *pagination.SearchPage[ProductDetail], err error) {
	var raw *http.Response
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/image-search"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Search the catalog by image (URL or base64), with pagination support.
//
// Provide exactly one of `image_url` or `base64_image`. For text or text+image
// search, use `POST /v1/search`.
func (r *ProductService) SearchByImageAutoPaging(ctx context.Context, params ProductSearchByImageParams, opts ...option.RequestOption) *pagination.SearchPageAutoPager[ProductDetail] {
	return pagination.NewSearchPageAutoPager(r.SearchByImage(ctx, params, opts...))
}

// Filter-driven product listing with pagination (no free-text query).
type BrowseRequestParam struct {
	// Optional limit on the number of results. Default is 20, max is 30.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Opaque token from a previous browse response to fetch the next page.
	PageToken param.Opt[string] `json:"page_token,omitzero"`
	// Filters to browse by. At least one of `brand_ids`, `category_ids`, or
	// `website_ids` must be provided.
	Filters SearchFiltersParam `json:"filters,omitzero"`
	paramObj
}

func (r BrowseRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BrowseRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrowseRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image-only search request.
type ImageSearchRequestParam struct {
	// Base64 encoded image bytes (no data URI prefix).
	Base64Image param.Opt[string] `json:"base64_image,omitzero"`
	// Publicly accessible URL of the image to search with.
	ImageURL param.Opt[string] `json:"image_url,omitzero"`
	// Optional limit on the number of results. Default is 20, max is 30.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Opaque token from a previous image-search response to fetch the next page of
	// results.
	PageToken param.Opt[string] `json:"page_token,omitzero"`
	// Image segmentation mode. None (default) disables segmentation. "AUTO" segments
	// and crops the main product automatically. A custom string (e.g. "shoe", "mug")
	// segments the specified object.
	Segment param.Opt[string] `json:"segment,omitzero"`
	// Optional locale configuration.
	Config LocaleConfigParam `json:"config,omitzero"`
	// Optional filters. Search will only consider products that match all of the
	// filters.
	Filters SearchFiltersParam `json:"filters,omitzero"`
	paramObj
}

func (r ImageSearchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ImageSearchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImageSearchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Locale options for API requests.
//
// Locale fields are optional; the server infers missing values. Details are on
// `language`, `country`, and `currency` below.
type LocaleConfigParam struct {
	// ISO 3166-1 alpha-2 country code (plus the pan-region `EU`).
	//
	// Any of "US", "GB", "EU", "AU", "CA", "IE", "DE", "AT", "FR", "BE", "IT", "ES",
	// "NL", "SE", "FI", "PT", "CZ", "GR", "RO".
	Country LocaleConfigCountry `json:"country,omitzero"`
	// ISO 4217 currency code.
	//
	// Any of "USD", "CAD", "AUD", "GBP", "EUR", "SEK", "CZK", "RON".
	Currency LocaleConfigCurrency `json:"currency,omitzero"`
	// ISO 639-1 language code.
	//
	// Any of "en", "de", "fr", "it", "es", "nl", "sv", "fi", "pt", "cs", "el", "ro".
	Language LocaleConfigLanguage `json:"language,omitzero"`
	// Preferred unit for length dimensions (length/width/height) in responses. A
	// request dimension filter's unit for the field takes precedence; when neither is
	// set, the merchant's stated unit is returned.
	//
	// Any of "mm", "cm", "m", "in", "ft".
	LengthUnit LocaleConfigLengthUnit `json:"length_unit,omitzero"`
	// Preferred unit for weight dimensions in responses. A request dimension filter's
	// weight unit takes precedence; when neither is set, the merchant's stated unit is
	// returned.
	//
	// Any of "mg", "g", "kg", "oz", "lb".
	WeightUnit LocaleConfigWeightUnit `json:"weight_unit,omitzero"`
	paramObj
}

func (r LocaleConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow LocaleConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocaleConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ISO 3166-1 alpha-2 country code (plus the pan-region `EU`).
type LocaleConfigCountry string

const (
	LocaleConfigCountryUs LocaleConfigCountry = "US"
	LocaleConfigCountryGB LocaleConfigCountry = "GB"
	LocaleConfigCountryEu LocaleConfigCountry = "EU"
	LocaleConfigCountryAu LocaleConfigCountry = "AU"
	LocaleConfigCountryCa LocaleConfigCountry = "CA"
	LocaleConfigCountryIe LocaleConfigCountry = "IE"
	LocaleConfigCountryDe LocaleConfigCountry = "DE"
	LocaleConfigCountryAt LocaleConfigCountry = "AT"
	LocaleConfigCountryFr LocaleConfigCountry = "FR"
	LocaleConfigCountryBe LocaleConfigCountry = "BE"
	LocaleConfigCountryIt LocaleConfigCountry = "IT"
	LocaleConfigCountryEs LocaleConfigCountry = "ES"
	LocaleConfigCountryNl LocaleConfigCountry = "NL"
	LocaleConfigCountrySe LocaleConfigCountry = "SE"
	LocaleConfigCountryFi LocaleConfigCountry = "FI"
	LocaleConfigCountryPt LocaleConfigCountry = "PT"
	LocaleConfigCountryCz LocaleConfigCountry = "CZ"
	LocaleConfigCountryGr LocaleConfigCountry = "GR"
	LocaleConfigCountryRo LocaleConfigCountry = "RO"
)

// ISO 4217 currency code.
type LocaleConfigCurrency string

const (
	LocaleConfigCurrencyUsd LocaleConfigCurrency = "USD"
	LocaleConfigCurrencyCad LocaleConfigCurrency = "CAD"
	LocaleConfigCurrencyAud LocaleConfigCurrency = "AUD"
	LocaleConfigCurrencyGbp LocaleConfigCurrency = "GBP"
	LocaleConfigCurrencyEur LocaleConfigCurrency = "EUR"
	LocaleConfigCurrencySek LocaleConfigCurrency = "SEK"
	LocaleConfigCurrencyCzk LocaleConfigCurrency = "CZK"
	LocaleConfigCurrencyRon LocaleConfigCurrency = "RON"
)

// ISO 639-1 language code.
type LocaleConfigLanguage string

const (
	LocaleConfigLanguageEn LocaleConfigLanguage = "en"
	LocaleConfigLanguageDe LocaleConfigLanguage = "de"
	LocaleConfigLanguageFr LocaleConfigLanguage = "fr"
	LocaleConfigLanguageIt LocaleConfigLanguage = "it"
	LocaleConfigLanguageEs LocaleConfigLanguage = "es"
	LocaleConfigLanguageNl LocaleConfigLanguage = "nl"
	LocaleConfigLanguageSv LocaleConfigLanguage = "sv"
	LocaleConfigLanguageFi LocaleConfigLanguage = "fi"
	LocaleConfigLanguagePt LocaleConfigLanguage = "pt"
	LocaleConfigLanguageCs LocaleConfigLanguage = "cs"
	LocaleConfigLanguageEl LocaleConfigLanguage = "el"
	LocaleConfigLanguageRo LocaleConfigLanguage = "ro"
)

// Preferred unit for length dimensions (length/width/height) in responses. A
// request dimension filter's unit for the field takes precedence; when neither is
// set, the merchant's stated unit is returned.
type LocaleConfigLengthUnit string

const (
	LocaleConfigLengthUnitMm LocaleConfigLengthUnit = "mm"
	LocaleConfigLengthUnitCm LocaleConfigLengthUnit = "cm"
	LocaleConfigLengthUnitM  LocaleConfigLengthUnit = "m"
	LocaleConfigLengthUnitIn LocaleConfigLengthUnit = "in"
	LocaleConfigLengthUnitFt LocaleConfigLengthUnit = "ft"
)

// Preferred unit for weight dimensions in responses. A request dimension filter's
// weight unit takes precedence; when neither is set, the merchant's stated unit is
// returned.
type LocaleConfigWeightUnit string

const (
	LocaleConfigWeightUnitMg LocaleConfigWeightUnit = "mg"
	LocaleConfigWeightUnitG  LocaleConfigWeightUnit = "g"
	LocaleConfigWeightUnitKg LocaleConfigWeightUnit = "kg"
	LocaleConfigWeightUnitOz LocaleConfigWeightUnit = "oz"
	LocaleConfigWeightUnitLb LocaleConfigWeightUnit = "lb"
)

// The property URL is required.
type LookupRequestParam struct {
	// The URL of the product to look up
	URL string `json:"url" api:"required"`
	// Maximum age (in hours) of cached product data before forcing a fresh lookup.
	// Defaults to 3 hours.
	MaxStalenessHours param.Opt[int64] `json:"max_staleness_hours,omitzero"`
	paramObj
}

func (r LookupRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow LookupRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LookupRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from the /v1/lookup endpoint.
type LookupResponse struct {
	// Product with detailed information.
	Product ProductDetail `json:"product" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Product     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LookupResponse) RawJSON() string { return r.JSON.raw }
func (r *LookupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonetizeOffer struct {
	// Merchant domain, e.g. nordstrom.com
	Domain string `json:"domain" api:"required"`
	// buy.trychannel3.com deeplink. Clicks are tracked and routed through the
	// highest-paying affiliate network for the merchant.
	URL string `json:"url" api:"required"`
	// Maximum post-take-rate commission for the merchant, as a decimal (0.05 = 5%).
	// 'Max' because the realized rate may be lower.
	MaxCommissionRate float64 `json:"max_commission_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain            respjson.Field
		URL               respjson.Field
		MaxCommissionRate respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonetizeOffer) RawJSON() string { return r.JSON.raw }
func (r *MonetizeOffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type MonetizeRequestParam struct {
	// The URL of the product to monetize
	URL string `json:"url" api:"required"`
	paramObj
}

func (r MonetizeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MonetizeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonetizeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from the /v1/monetize endpoint — just the list of offers.
type MonetizeResponse struct {
	// Monetizable offers, sorted by max_commission_rate descending.
	Offers []MonetizeOffer `json:"offers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offers      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonetizeResponse) RawJSON() string { return r.JSON.raw }
func (r *MonetizeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Price struct {
	// The currency code of the product, like USD, EUR, GBP, etc.
	Currency string `json:"currency" api:"required"`
	// The current price of the product, including any discounts.
	Price float64 `json:"price" api:"required"`
	// The original price of the product before any discounts.
	CompareAtPrice float64 `json:"compare_at_price" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Currency       respjson.Field
		Price          respjson.Field
		CompareAtPrice respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Price) RawJSON() string { return r.JSON.raw }
func (r *Price) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductBrand struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductBrand) RawJSON() string { return r.JSON.raw }
func (r *ProductBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product with detailed information.
type ProductDetail struct {
	ID    string `json:"id" api:"required"`
	Title string `json:"title" api:"required"`
	// Target age group. Age-agnostic products are typically returned as 'adult'.
	//
	// Any of "newborn", "infant", "toddler", "kids", "adult".
	Age ProductDetailAge `json:"age" api:"nullable"`
	// Ordered list of brands.
	Brands []ProductBrand `json:"brands"`
	// Lean category representation used in search hits and list rows.
	Category    CategorySummary `json:"category" api:"nullable"`
	Description string          `json:"description" api:"nullable"`
	// Product gender. 'unisex' is deprecated: coerced to None on input, never emitted.
	//
	// Any of "male", "female".
	Gender      ProductDetailGender `json:"gender" api:"nullable"`
	Images      []ProductImage      `json:"images"`
	KeyFeatures []string            `json:"key_features" api:"nullable"`
	Materials   []string            `json:"materials" api:"nullable"`
	// All merchant offers for this product in the requested locale.
	Offers []ProductOffer `json:"offers"`
	// Structured attributes extracted for this product, keyed by attribute handle
	// (e.g. 'color', 'material'). Values are the canonical allowed values for that
	// handle.
	StructuredAttributes map[string][]string `json:"structured_attributes"`
	// Wrapper for variant-interaction state on a Product.
	//
	// Holds `options` and `selected`. `options` represent all of the configuration
	// options for the product. `selected` represents the currently selected option
	// values.
	Variants ProductDetailVariants `json:"variants" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Title                respjson.Field
		Age                  respjson.Field
		Brands               respjson.Field
		Category             respjson.Field
		Description          respjson.Field
		Gender               respjson.Field
		Images               respjson.Field
		KeyFeatures          respjson.Field
		Materials            respjson.Field
		Offers               respjson.Field
		StructuredAttributes respjson.Field
		Variants             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductDetail) RawJSON() string { return r.JSON.raw }
func (r *ProductDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Target age group. Age-agnostic products are typically returned as 'adult'.
type ProductDetailAge string

const (
	ProductDetailAgeNewborn ProductDetailAge = "newborn"
	ProductDetailAgeInfant  ProductDetailAge = "infant"
	ProductDetailAgeToddler ProductDetailAge = "toddler"
	ProductDetailAgeKids    ProductDetailAge = "kids"
	ProductDetailAgeAdult   ProductDetailAge = "adult"
)

// Product gender. 'unisex' is deprecated: coerced to None on input, never emitted.
type ProductDetailGender string

const (
	ProductDetailGenderMale   ProductDetailGender = "male"
	ProductDetailGenderFemale ProductDetailGender = "female"
)

// Wrapper for variant-interaction state on a Product.
//
// Holds `options` and `selected`. `options` represent all of the configuration
// options for the product. `selected` represents the currently selected option
// values.
type ProductDetailVariants struct {
	Options  []ProductDetailVariantsOption   `json:"options" api:"required"`
	Selected []ProductDetailVariantsSelected `json:"selected" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		Selected    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductDetailVariants) RawJSON() string { return r.JSON.raw }
func (r *ProductDetailVariants) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One dimension of a product family (e.g. 'Color', 'Size').
type ProductDetailVariantsOption struct {
	// The name of the option (e.g. 'Color', 'Size')
	Name string `json:"name" api:"required"`
	// The values of the option (e.g. ['Blue', 'Red', 'Green'])
	Values []ProductDetailVariantsOptionValue `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductDetailVariantsOption) RawJSON() string { return r.JSON.raw }
func (r *ProductDetailVariantsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One value of one variant option (e.g. 'Blue' under 'Color')
type ProductDetailVariantsOptionValue struct {
	// Whether the option value exists on the product, or is a configuration only
	// present on another variant of the same product. For example, a shirt that comes
	// in multiple colors, but only one color is available in Size XL.
	Exists bool `json:"exists" api:"required"`
	// The display value of the option value (e.g. 'Blue')
	Label string `json:"label" api:"required"`
	// The two availability values the public API emits on offers.
	//
	// Internal `AvailabilityStatus` values are collapsed to these via
	// `AvailabilityStatus.to_api()`.
	//
	// Any of "InStock", "OutOfStock".
	Available string `json:"available" api:"nullable"`
	// The product id that represents this value. Variants that point to different
	// products will have this field set, as well as thumbnail_url for displaying
	// selector icons.
	ProductID string `json:"product_id" api:"nullable"`
	// For options that reference different products, this is the URL of the thumbnail
	// image for the option value. E.g., a shoe that comes in multiple colors will have
	// an OptionValue for each color with a thumbnail_url set.
	ThumbnailURL string `json:"thumbnail_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exists       respjson.Field
		Label        respjson.Field
		Available    respjson.Field
		ProductID    respjson.Field
		ThumbnailURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductDetailVariantsOptionValue) RawJSON() string { return r.JSON.raw }
func (r *ProductDetailVariantsOptionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One effective selection on a product, post server-side relaxation.
type ProductDetailVariantsSelected struct {
	// The display value of the selected option (e.g. 'Blue', 'XL')
	Label string `json:"label" api:"required"`
	// The name of the selected option (e.g. 'Color', 'Size')
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductDetailVariantsSelected) RawJSON() string { return r.JSON.raw }
func (r *ProductDetailVariantsSelected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product image with metadata.
type ProductImage struct {
	URL     string `json:"url" api:"required"`
	AltText string `json:"alt_text" api:"nullable"`
	// Background-removed square image on Channel3 CDN when available. Use for product
	// grids; `url` is the regular hosted shot.
	CleanedURL  string `json:"cleaned_url" api:"nullable"`
	IsMainImage bool   `json:"is_main_image"`
	// Product image type classification for API responses.
	//
	// Any of "hero", "lifestyle", "on_model", "detail", "scale_reference",
	// "angle_view", "flat_lay", "in_use", "packaging", "size_chart",
	// "product_information", "merchant_information".
	ShotType ProductImageShotType `json:"shot_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		AltText     respjson.Field
		CleanedURL  respjson.Field
		IsMainImage respjson.Field
		ShotType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductImage) RawJSON() string { return r.JSON.raw }
func (r *ProductImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product image type classification for API responses.
type ProductImageShotType string

const (
	ProductImageShotTypeHero                ProductImageShotType = "hero"
	ProductImageShotTypeLifestyle           ProductImageShotType = "lifestyle"
	ProductImageShotTypeOnModel             ProductImageShotType = "on_model"
	ProductImageShotTypeDetail              ProductImageShotType = "detail"
	ProductImageShotTypeScaleReference      ProductImageShotType = "scale_reference"
	ProductImageShotTypeAngleView           ProductImageShotType = "angle_view"
	ProductImageShotTypeFlatLay             ProductImageShotType = "flat_lay"
	ProductImageShotTypeInUse               ProductImageShotType = "in_use"
	ProductImageShotTypePackaging           ProductImageShotType = "packaging"
	ProductImageShotTypeSizeChart           ProductImageShotType = "size_chart"
	ProductImageShotTypeProductInformation  ProductImageShotType = "product_information"
	ProductImageShotTypeMerchantInformation ProductImageShotType = "merchant_information"
)

type ProductOffer struct {
	// The two availability values the public API emits on offers.
	//
	// Internal `AvailabilityStatus` values are collapsed to these via
	// `AvailabilityStatus.to_api()`.
	//
	// Any of "InStock", "OutOfStock".
	Availability ProductOfferAvailability `json:"availability" api:"required"`
	Domain       string                   `json:"domain" api:"required"`
	Price        Price                    `json:"price" api:"required"`
	URL          string                   `json:"url" api:"required"`
	// Offer condition. 'refurbished' is deprecated: rejected as a filter value,
	// coerced to None on responses.
	//
	// Any of "new", "used".
	Condition ProductOfferCondition `json:"condition" api:"nullable"`
	// Physical dimensions of a product offer. Members are null when unknown.
	//
	// Values are standardized to the supported unit set; a merchant-stated value whose
	// unit is not one of those units is omitted rather than shown.
	Dimensions ProductOfferDimensions `json:"dimensions" api:"nullable"`
	// The maximum commission rate for the merchant, as a decimal fraction: 0 is no
	// commission, 0.5 is 50% commission. 'Max' because the actual commission rate may
	// be lower due to vendor-specific affiliate rules.
	MaxCommissionRate float64 `json:"max_commission_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availability      respjson.Field
		Domain            respjson.Field
		Price             respjson.Field
		URL               respjson.Field
		Condition         respjson.Field
		Dimensions        respjson.Field
		MaxCommissionRate respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductOffer) RawJSON() string { return r.JSON.raw }
func (r *ProductOffer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The two availability values the public API emits on offers.
//
// Internal `AvailabilityStatus` values are collapsed to these via
// `AvailabilityStatus.to_api()`.
type ProductOfferAvailability string

const (
	ProductOfferAvailabilityInStock    ProductOfferAvailability = "InStock"
	ProductOfferAvailabilityOutOfStock ProductOfferAvailability = "OutOfStock"
)

// Offer condition. 'refurbished' is deprecated: rejected as a filter value,
// coerced to None on responses.
type ProductOfferCondition string

const (
	ProductOfferConditionNew  ProductOfferCondition = "new"
	ProductOfferConditionUsed ProductOfferCondition = "used"
)

// Physical dimensions of a product offer. Members are null when unknown.
//
// Values are standardized to the supported unit set; a merchant-stated value whose
// unit is not one of those units is omitted rather than shown.
type ProductOfferDimensions struct {
	// A length measurement, in one of the supported length units.
	Height ProductOfferDimensionsHeight `json:"height" api:"nullable"`
	// A length measurement, in one of the supported length units.
	Length ProductOfferDimensionsLength `json:"length" api:"nullable"`
	// A weight measurement, in one of the supported weight units.
	Weight ProductOfferDimensionsWeight `json:"weight" api:"nullable"`
	// A length measurement, in one of the supported length units.
	Width ProductOfferDimensionsWidth `json:"width" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Height      respjson.Field
		Length      respjson.Field
		Weight      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductOfferDimensions) RawJSON() string { return r.JSON.raw }
func (r *ProductOfferDimensions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A length measurement, in one of the supported length units.
type ProductOfferDimensionsHeight struct {
	Number float64 `json:"number" api:"required"`
	// The unit from the request's dimension filters when one was given (the value is
	// converted to it); otherwise the unit the merchant stated.
	//
	// Any of "mm", "cm", "m", "in", "ft".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductOfferDimensionsHeight) RawJSON() string { return r.JSON.raw }
func (r *ProductOfferDimensionsHeight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A length measurement, in one of the supported length units.
type ProductOfferDimensionsLength struct {
	Number float64 `json:"number" api:"required"`
	// The unit from the request's dimension filters when one was given (the value is
	// converted to it); otherwise the unit the merchant stated.
	//
	// Any of "mm", "cm", "m", "in", "ft".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductOfferDimensionsLength) RawJSON() string { return r.JSON.raw }
func (r *ProductOfferDimensionsLength) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A weight measurement, in one of the supported weight units.
type ProductOfferDimensionsWeight struct {
	Number float64 `json:"number" api:"required"`
	// The unit from the request's dimension filters when one was given (the value is
	// converted to it); otherwise the unit the merchant stated.
	//
	// Any of "mg", "g", "kg", "oz", "lb".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductOfferDimensionsWeight) RawJSON() string { return r.JSON.raw }
func (r *ProductOfferDimensionsWeight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A length measurement, in one of the supported length units.
type ProductOfferDimensionsWidth struct {
	Number float64 `json:"number" api:"required"`
	// The unit from the request's dimension filters when one was given (the value is
	// converted to it); otherwise the unit the merchant stated.
	//
	// Any of "mm", "cm", "m", "in", "ft".
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductOfferDimensionsWidth) RawJSON() string { return r.JSON.raw }
func (r *ProductOfferDimensionsWidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Find products similar to a given product.
//
// The property ProductID is required.
type SimilarProductsRequestParam struct {
	// Canonical product ID to find similar products for.
	ProductID string `json:"product_id" api:"required"`
	// Optional limit on the number of results. Default is 20, max is 30.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Opaque token from a previous similar response to fetch the next page of results.
	PageToken param.Opt[string] `json:"page_token,omitzero"`
	// Optional locale configuration.
	Config LocaleConfigParam `json:"config,omitzero"`
	// Optional filters. Search will only consider products that match all of the
	// filters.
	Filters SearchFiltersParam `json:"filters,omitzero"`
	paramObj
}

func (r SimilarProductsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SimilarProductsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimilarProductsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductGetParams struct {
	// Optional user identifier to attribute clicks and sales to a user in your system.
	// Channel3 appends it to buy URLs in the response.
	XUserID param.Opt[string] `header:"x-user-id,omitzero" json:"-"`
	// ISO 3166-1 alpha-2 country code. Matches any country when unset; defaults to
	// 'US' only when language and currency are also unset.
	//
	// Any of "US", "GB", "EU", "AU", "CA", "IE", "DE", "AT", "FR", "BE", "IT", "ES",
	// "NL", "SE", "FI", "PT", "CZ", "GR", "RO".
	Country ProductGetParamsCountry `query:"country,omitzero" json:"-"`
	// ISO 4217 currency code. When unset, inferred from `country` (e.g. GB -> GBP);
	// falls back to 'USD' only when all three locale fields are unset.
	//
	// Any of "USD", "CAD", "AUD", "GBP", "EUR", "SEK", "CZK", "RON".
	Currency ProductGetParamsCurrency `query:"currency,omitzero" json:"-"`
	// ISO 639-1 language code. Matches any language when unset; defaults to 'en' only
	// when country and currency are also unset.
	//
	// Any of "en", "de", "fr", "it", "es", "nl", "sv", "fi", "pt", "cs", "el", "ro".
	Language ProductGetParamsLanguage `query:"language,omitzero" json:"-"`
	// Preferred unit for length dimensions (length/width/height). When unset,
	// dimensions are returned in the unit the merchant stated.
	//
	// Any of "mm", "cm", "m", "in", "ft".
	LengthUnit ProductGetParamsLengthUnit `query:"length_unit,omitzero" json:"-"`
	// Optional list of website IDs to constrain the buy URL to, relevant if multiple
	// merchants exist. Accepts website IDs or domains (e.g. "nike.com").
	WebsiteIDs []string `query:"website_ids,omitzero" json:"-"`
	// Preferred unit for weight dimensions. When unset, weight is returned in the unit
	// the merchant stated.
	//
	// Any of "mg", "g", "kg", "oz", "lb".
	WeightUnit ProductGetParamsWeightUnit `query:"weight_unit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProductGetParams]'s query parameters as `url.Values`.
func (r ProductGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ISO 3166-1 alpha-2 country code. Matches any country when unset; defaults to
// 'US' only when language and currency are also unset.
type ProductGetParamsCountry string

const (
	ProductGetParamsCountryUs ProductGetParamsCountry = "US"
	ProductGetParamsCountryGB ProductGetParamsCountry = "GB"
	ProductGetParamsCountryEu ProductGetParamsCountry = "EU"
	ProductGetParamsCountryAu ProductGetParamsCountry = "AU"
	ProductGetParamsCountryCa ProductGetParamsCountry = "CA"
	ProductGetParamsCountryIe ProductGetParamsCountry = "IE"
	ProductGetParamsCountryDe ProductGetParamsCountry = "DE"
	ProductGetParamsCountryAt ProductGetParamsCountry = "AT"
	ProductGetParamsCountryFr ProductGetParamsCountry = "FR"
	ProductGetParamsCountryBe ProductGetParamsCountry = "BE"
	ProductGetParamsCountryIt ProductGetParamsCountry = "IT"
	ProductGetParamsCountryEs ProductGetParamsCountry = "ES"
	ProductGetParamsCountryNl ProductGetParamsCountry = "NL"
	ProductGetParamsCountrySe ProductGetParamsCountry = "SE"
	ProductGetParamsCountryFi ProductGetParamsCountry = "FI"
	ProductGetParamsCountryPt ProductGetParamsCountry = "PT"
	ProductGetParamsCountryCz ProductGetParamsCountry = "CZ"
	ProductGetParamsCountryGr ProductGetParamsCountry = "GR"
	ProductGetParamsCountryRo ProductGetParamsCountry = "RO"
)

// ISO 4217 currency code. When unset, inferred from `country` (e.g. GB -> GBP);
// falls back to 'USD' only when all three locale fields are unset.
type ProductGetParamsCurrency string

const (
	ProductGetParamsCurrencyUsd ProductGetParamsCurrency = "USD"
	ProductGetParamsCurrencyCad ProductGetParamsCurrency = "CAD"
	ProductGetParamsCurrencyAud ProductGetParamsCurrency = "AUD"
	ProductGetParamsCurrencyGbp ProductGetParamsCurrency = "GBP"
	ProductGetParamsCurrencyEur ProductGetParamsCurrency = "EUR"
	ProductGetParamsCurrencySek ProductGetParamsCurrency = "SEK"
	ProductGetParamsCurrencyCzk ProductGetParamsCurrency = "CZK"
	ProductGetParamsCurrencyRon ProductGetParamsCurrency = "RON"
)

// ISO 639-1 language code. Matches any language when unset; defaults to 'en' only
// when country and currency are also unset.
type ProductGetParamsLanguage string

const (
	ProductGetParamsLanguageEn ProductGetParamsLanguage = "en"
	ProductGetParamsLanguageDe ProductGetParamsLanguage = "de"
	ProductGetParamsLanguageFr ProductGetParamsLanguage = "fr"
	ProductGetParamsLanguageIt ProductGetParamsLanguage = "it"
	ProductGetParamsLanguageEs ProductGetParamsLanguage = "es"
	ProductGetParamsLanguageNl ProductGetParamsLanguage = "nl"
	ProductGetParamsLanguageSv ProductGetParamsLanguage = "sv"
	ProductGetParamsLanguageFi ProductGetParamsLanguage = "fi"
	ProductGetParamsLanguagePt ProductGetParamsLanguage = "pt"
	ProductGetParamsLanguageCs ProductGetParamsLanguage = "cs"
	ProductGetParamsLanguageEl ProductGetParamsLanguage = "el"
	ProductGetParamsLanguageRo ProductGetParamsLanguage = "ro"
)

// Preferred unit for length dimensions (length/width/height). When unset,
// dimensions are returned in the unit the merchant stated.
type ProductGetParamsLengthUnit string

const (
	ProductGetParamsLengthUnitMm ProductGetParamsLengthUnit = "mm"
	ProductGetParamsLengthUnitCm ProductGetParamsLengthUnit = "cm"
	ProductGetParamsLengthUnitM  ProductGetParamsLengthUnit = "m"
	ProductGetParamsLengthUnitIn ProductGetParamsLengthUnit = "in"
	ProductGetParamsLengthUnitFt ProductGetParamsLengthUnit = "ft"
)

// Preferred unit for weight dimensions. When unset, weight is returned in the unit
// the merchant stated.
type ProductGetParamsWeightUnit string

const (
	ProductGetParamsWeightUnitMg ProductGetParamsWeightUnit = "mg"
	ProductGetParamsWeightUnitG  ProductGetParamsWeightUnit = "g"
	ProductGetParamsWeightUnitKg ProductGetParamsWeightUnit = "kg"
	ProductGetParamsWeightUnitOz ProductGetParamsWeightUnit = "oz"
	ProductGetParamsWeightUnitLb ProductGetParamsWeightUnit = "lb"
)

type ProductBrowseParams struct {
	// Filter-driven product listing with pagination (no free-text query).
	BrowseRequest BrowseRequestParam
	// Optional user identifier to attribute clicks and sales to a user in your system.
	// Channel3 appends it to buy URLs in the response.
	XUserID param.Opt[string] `header:"x-user-id,omitzero" json:"-"`
	paramObj
}

func (r ProductBrowseParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BrowseRequest)
}
func (r *ProductBrowseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductFindSimilarParams struct {
	// Find products similar to a given product.
	SimilarProductsRequest SimilarProductsRequestParam
	// Optional user identifier to attribute clicks and sales to a user in your system.
	// Channel3 appends it to buy URLs in the response.
	XUserID param.Opt[string] `header:"x-user-id,omitzero" json:"-"`
	paramObj
}

func (r ProductFindSimilarParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SimilarProductsRequest)
}
func (r *ProductFindSimilarParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductLookupParams struct {
	LookupRequest LookupRequestParam
	// Optional user identifier to attribute clicks and sales to a user in your system.
	// Channel3 appends it to buy URLs in the response.
	XUserID param.Opt[string] `header:"x-user-id,omitzero" json:"-"`
	paramObj
}

func (r ProductLookupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LookupRequest)
}
func (r *ProductLookupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductMonetizeParams struct {
	MonetizeRequest MonetizeRequestParam
	// Optional user identifier to attribute clicks and sales to a user in your system.
	// Channel3 appends it to buy URLs in the response.
	XUserID param.Opt[string] `header:"x-user-id,omitzero" json:"-"`
	paramObj
}

func (r ProductMonetizeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MonetizeRequest)
}
func (r *ProductMonetizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductSearchParams struct {
	// Search request with pagination support.
	SearchRequest SearchRequestParam
	// Optional user identifier to attribute clicks and sales to a user in your system.
	// Channel3 appends it to buy URLs in the response.
	XUserID param.Opt[string] `header:"x-user-id,omitzero" json:"-"`
	paramObj
}

func (r ProductSearchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SearchRequest)
}
func (r *ProductSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductSearchByImageParams struct {
	// Image-only search request.
	ImageSearchRequest ImageSearchRequestParam
	// Optional user identifier to attribute clicks and sales to a user in your system.
	// Channel3 appends it to buy URLs in the response.
	XUserID param.Opt[string] `header:"x-user-id,omitzero" json:"-"`
	paramObj
}

func (r ProductSearchByImageParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ImageSearchRequest)
}
func (r *ProductSearchByImageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
