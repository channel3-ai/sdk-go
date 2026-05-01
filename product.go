// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package publicsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/public-sdk-go/internal/apijson"
	"github.com/stainless-sdks/public-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/public-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/public-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/public-sdk-go/option"
	"github.com/stainless-sdks/public-sdk-go/packages/pagination"
	"github.com/stainless-sdks/public-sdk-go/packages/param"
	"github.com/stainless-sdks/public-sdk-go/packages/respjson"
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
func (r *ProductService) Get(ctx context.Context, productID string, query ProductGetParams, opts ...option.RequestOption) (res *ProductDetail, err error) {
	opts = slices.Concat(r.options, opts)
	if productID == "" {
		err = errors.New("missing required product_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/products/%s", url.PathEscape(productID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Find products similar to a given product.
//
// Consider setting `filters` to narrow results to the same gender, brand,
// category, price range, etc. when you only want similar items within a specific
// slice of the catalog.
func (r *ProductService) FindSimilar(ctx context.Context, body ProductFindSimilarParams, opts ...option.RequestOption) (res *pagination.SearchPage[ProductDetail], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/similar"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
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
func (r *ProductService) FindSimilarAutoPaging(ctx context.Context, body ProductFindSimilarParams, opts ...option.RequestOption) *pagination.SearchPageAutoPager[ProductDetail] {
	return pagination.NewSearchPageAutoPager(r.FindSimilar(ctx, body, opts...))
}

// Retrieve product information for any supported product URL.
//
// Returns the same Product model as GET /v1/products/{product_id}. The product_id
// in the response can be used with the Product Detail endpoint.
func (r *ProductService) Lookup(ctx context.Context, body ProductLookupParams, opts ...option.RequestOption) (res *LookupResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Search for products with pagination support.
//
// At least one of `query`, `image_url`, or `base64_image` must be provided;
// requests with none of these will return 422.
func (r *ProductService) Search(ctx context.Context, body ProductSearchParams, opts ...option.RequestOption) (res *pagination.SearchPage[ProductDetail], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/search"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
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
// At least one of `query`, `image_url`, or `base64_image` must be provided;
// requests with none of these will return 422.
func (r *ProductService) SearchAutoPaging(ctx context.Context, body ProductSearchParams, opts ...option.RequestOption) *pagination.SearchPageAutoPager[ProductDetail] {
	return pagination.NewSearchPageAutoPager(r.Search(ctx, body, opts...))
}

// Search the catalog by image (URL or base64), with pagination support.
//
// Provide exactly one of `image_url` or `base64_image`. For text or text+image
// search, use `POST /v1/search`.
func (r *ProductService) SearchByImage(ctx context.Context, body ProductSearchByImageParams, opts ...option.RequestOption) (res *pagination.SearchPage[ProductDetail], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/image-search"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
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
func (r *ProductService) SearchByImageAutoPaging(ctx context.Context, body ProductSearchByImageParams, opts ...option.RequestOption) *pagination.SearchPageAutoPager[ProductDetail] {
	return pagination.NewSearchPageAutoPager(r.SearchByImage(ctx, body, opts...))
}

type AvailabilityStatus string

const (
	AvailabilityStatusInStock             AvailabilityStatus = "InStock"
	AvailabilityStatusLimitedAvailability AvailabilityStatus = "LimitedAvailability"
	AvailabilityStatusPreOrder            AvailabilityStatus = "PreOrder"
	AvailabilityStatusBackOrder           AvailabilityStatus = "BackOrder"
	AvailabilityStatusSoldOut             AvailabilityStatus = "SoldOut"
	AvailabilityStatusOutOfStock          AvailabilityStatus = "OutOfStock"
	AvailabilityStatusDiscontinued        AvailabilityStatus = "Discontinued"
	AvailabilityStatusUnknown             AvailabilityStatus = "Unknown"
)

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
	// ISO 3166-1 alpha-2 country code. May stay unset for pan-region storefronts (e.g.
	// `currency=EUR` with no specific country).
	//
	// Any of "US", "GB", "EU", "AU", "CA", "IE", "DE", "AT", "FR", "BE", "IT", "ES",
	// "NL", "SE", "FI", "PT", "CZ".
	Country LocaleConfigCountry `json:"country,omitzero"`
	// ISO 4217 currency code. When unset, inferred from `country` (e.g. `GB` → `GBP`),
	// defaulting to `USD`.
	//
	// Any of "USD", "CAD", "AUD", "GBP", "EUR", "SEK", "CZK".
	Currency LocaleConfigCurrency `json:"currency,omitzero"`
	// ISO 639-1 language code. When unset, inferred from `country` (preferred) then
	// `currency`, defaulting to `en`.
	//
	// Any of "en", "de", "fr", "it", "es", "nl", "sv", "fi", "pt", "cs".
	Language LocaleConfigLanguage `json:"language,omitzero"`
	paramObj
}

func (r LocaleConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow LocaleConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocaleConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ISO 3166-1 alpha-2 country code. May stay unset for pan-region storefronts (e.g.
// `currency=EUR` with no specific country).
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
)

// ISO 4217 currency code. When unset, inferred from `country` (e.g. `GB` → `GBP`),
// defaulting to `USD`.
type LocaleConfigCurrency string

const (
	LocaleConfigCurrencyUsd LocaleConfigCurrency = "USD"
	LocaleConfigCurrencyCad LocaleConfigCurrency = "CAD"
	LocaleConfigCurrencyAud LocaleConfigCurrency = "AUD"
	LocaleConfigCurrencyGbp LocaleConfigCurrency = "GBP"
	LocaleConfigCurrencyEur LocaleConfigCurrency = "EUR"
	LocaleConfigCurrencySek LocaleConfigCurrency = "SEK"
	LocaleConfigCurrencyCzk LocaleConfigCurrency = "CZK"
)

// ISO 639-1 language code. When unset, inferred from `country` (preferred) then
// `currency`, defaulting to `en`.
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
	Brands      []ProductBrand `json:"brands"`
	Categories  []string       `json:"categories"`
	Description string         `json:"description" api:"nullable"`
	// Any of "male", "female", "unisex".
	Gender      ProductDetailGender `json:"gender" api:"nullable"`
	Images      []ProductImage      `json:"images"`
	KeyFeatures []string            `json:"key_features" api:"nullable"`
	Materials   []string            `json:"materials" api:"nullable"`
	// All merchant offers for this product in the requested locale.
	Offers []ProductOffer `json:"offers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Title       respjson.Field
		Age         respjson.Field
		Brands      respjson.Field
		Categories  respjson.Field
		Description respjson.Field
		Gender      respjson.Field
		Images      respjson.Field
		KeyFeatures respjson.Field
		Materials   respjson.Field
		Offers      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

type ProductDetailGender string

const (
	ProductDetailGenderMale   ProductDetailGender = "male"
	ProductDetailGenderFemale ProductDetailGender = "female"
	ProductDetailGenderUnisex ProductDetailGender = "unisex"
)

// Product image with metadata.
type ProductImage struct {
	URL         string `json:"url" api:"required"`
	AltText     string `json:"alt_text" api:"nullable"`
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
	// Any of "InStock", "OutOfStock".
	Availability ProductOfferAvailability `json:"availability" api:"required"`
	Domain       string                   `json:"domain" api:"required"`
	Price        Price                    `json:"price" api:"required"`
	URL          string                   `json:"url" api:"required"`
	// The maximum commission rate for the merchant, as a percentage. 0 is no
	// commission. 0.5 is 50% commission. 'Max' because the actual commission rate may
	// be lower due to vendor-specific affiliate rules.
	MaxCommissionRate float64 `json:"max_commission_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Availability      respjson.Field
		Domain            respjson.Field
		Price             respjson.Field
		URL               respjson.Field
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

type ProductOfferAvailability string

const (
	ProductOfferAvailabilityInStock    ProductOfferAvailability = "InStock"
	ProductOfferAvailabilityOutOfStock ProductOfferAvailability = "OutOfStock"
)

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
	// ISO 3166-1 alpha-2 country code. Matches any country when unset; defaults to
	// 'US' only when language and currency are also unset.
	//
	// Any of "US", "GB", "EU", "AU", "CA", "IE", "DE", "AT", "FR", "BE", "IT", "ES",
	// "NL", "SE", "FI", "PT", "CZ".
	Country ProductGetParamsCountry `query:"country,omitzero" json:"-"`
	// ISO 4217 currency code. When unset, inferred from `country` (e.g. GB -> GBP);
	// falls back to 'USD' only when all three locale fields are unset.
	//
	// Any of "USD", "CAD", "AUD", "GBP", "EUR", "SEK", "CZK".
	Currency ProductGetParamsCurrency `query:"currency,omitzero" json:"-"`
	// ISO 639-1 language code. Matches any language when unset; defaults to 'en' only
	// when country and currency are also unset.
	//
	// Any of "en", "de", "fr", "it", "es", "nl", "sv", "fi", "pt", "cs".
	Language ProductGetParamsLanguage `query:"language,omitzero" json:"-"`
	// Optional list of website IDs to constrain the buy URL to, relevant if multiple
	// merchants exist
	WebsiteIDs []string `query:"website_ids,omitzero" json:"-"`
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
)

type ProductFindSimilarParams struct {
	// Find products similar to a given product.
	SimilarProductsRequest SimilarProductsRequestParam
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
	paramObj
}

func (r ProductLookupParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LookupRequest)
}
func (r *ProductLookupParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductSearchParams struct {
	// Search request with pagination support.
	SearchRequest SearchRequestParam
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
	paramObj
}

func (r ProductSearchByImageParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ImageSearchRequest)
}
func (r *ProductSearchByImageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
