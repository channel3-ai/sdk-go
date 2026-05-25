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

// EnrichService contains methods and other services that help with interacting
// with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnrichService] method instead.
//
// Deprecated: enrich is deprecated; migrate to `products.lookup`. This resource
// will be removed in the next major version.
type EnrichService struct {
	options []option.RequestOption
}

// NewEnrichService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEnrichService(opts ...option.RequestOption) (r EnrichService) {
	r = EnrichService{}
	r.options = opts
	return
}

// **Deprecated** — use POST /v1/lookup instead.
//
// Search by product URL, get back full product information from Channel3's product
// database.
//
// If the product is not found in the database, the endpoint will attempt real-time
// retrieval from the product page. This fallback returns basic product information
// (price, images, title) without full enrichment.
//
// Deprecated: use `products.lookup` instead; will be removed in the next major
// version
func (r *EnrichService) EnrichURL(ctx context.Context, body EnrichEnrichURLParams, opts ...option.RequestOption) (res *EnrichEnrichURLResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v0/enrich"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The property URL is required.
type EnrichRequestParam struct {
	// The URL of the product to enrich
	URL string `json:"url" api:"required"`
	paramObj
}

func (r EnrichRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow EnrichRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnrichRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// v0 product detail with deprecated fields.
type EnrichEnrichURLResponse struct {
	ID string `json:"id" api:"required"`
	// Deprecated, use offers field
	//
	// Any of "InStock", "OutOfStock".
	//
	// Deprecated: deprecated
	Availability EnrichEnrichURLResponseAvailability `json:"availability" api:"required"`
	// Deprecated, use offers field
	//
	// Deprecated: deprecated
	Price Price  `json:"price" api:"required"`
	Title string `json:"title" api:"required"`
	// Deprecated, use offers field
	//
	// Deprecated: deprecated
	URL string `json:"url" api:"required"`
	// Target age group. Age-agnostic products are typically returned as 'adult'.
	//
	// Any of "newborn", "infant", "toddler", "kids", "adult".
	Age EnrichEnrichURLResponseAge `json:"age" api:"nullable"`
	// Deprecated: deprecated
	BrandID string `json:"brand_id" api:"nullable"`
	// Deprecated: deprecated
	BrandName string `json:"brand_name" api:"nullable"`
	// Ordered list of brands.
	Brands []ProductBrand `json:"brands"`
	// Deprecated: deprecated
	Categories []string `json:"categories"`
	// Lean category representation used in search hits and list rows.
	Category    CategorySummary `json:"category" api:"nullable"`
	Description string          `json:"description" api:"nullable"`
	// Any of "male", "female", "unisex".
	Gender EnrichEnrichURLResponseGender `json:"gender" api:"nullable"`
	// List of image URLs (deprecated, use images field)
	//
	// Deprecated: deprecated
	ImageURLs   []string                       `json:"image_urls"`
	Images      []EnrichEnrichURLResponseImage `json:"images"`
	KeyFeatures []string                       `json:"key_features" api:"nullable"`
	Materials   []string                       `json:"materials" api:"nullable"`
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
	Variants EnrichEnrichURLResponseVariants `json:"variants" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Availability         respjson.Field
		Price                respjson.Field
		Title                respjson.Field
		URL                  respjson.Field
		Age                  respjson.Field
		BrandID              respjson.Field
		BrandName            respjson.Field
		Brands               respjson.Field
		Categories           respjson.Field
		Category             respjson.Field
		Description          respjson.Field
		Gender               respjson.Field
		ImageURLs            respjson.Field
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
func (r EnrichEnrichURLResponse) RawJSON() string { return r.JSON.raw }
func (r *EnrichEnrichURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated, use offers field
type EnrichEnrichURLResponseAvailability string

const (
	EnrichEnrichURLResponseAvailabilityInStock    EnrichEnrichURLResponseAvailability = "InStock"
	EnrichEnrichURLResponseAvailabilityOutOfStock EnrichEnrichURLResponseAvailability = "OutOfStock"
)

// Target age group. Age-agnostic products are typically returned as 'adult'.
type EnrichEnrichURLResponseAge string

const (
	EnrichEnrichURLResponseAgeNewborn EnrichEnrichURLResponseAge = "newborn"
	EnrichEnrichURLResponseAgeInfant  EnrichEnrichURLResponseAge = "infant"
	EnrichEnrichURLResponseAgeToddler EnrichEnrichURLResponseAge = "toddler"
	EnrichEnrichURLResponseAgeKids    EnrichEnrichURLResponseAge = "kids"
	EnrichEnrichURLResponseAgeAdult   EnrichEnrichURLResponseAge = "adult"
)

type EnrichEnrichURLResponseGender string

const (
	EnrichEnrichURLResponseGenderMale   EnrichEnrichURLResponseGender = "male"
	EnrichEnrichURLResponseGenderFemale EnrichEnrichURLResponseGender = "female"
	EnrichEnrichURLResponseGenderUnisex EnrichEnrichURLResponseGender = "unisex"
)

// v0 product image with deprecated photo_quality field.
type EnrichEnrichURLResponseImage struct {
	URL         string `json:"url" api:"required"`
	AltText     string `json:"alt_text" api:"nullable"`
	IsMainImage bool   `json:"is_main_image"`
	// Photo quality classification for API responses.
	//
	// Any of "professional", "ugc", "poor".
	//
	// Deprecated: deprecated
	PhotoQuality string `json:"photo_quality" api:"nullable"`
	// Product image type classification for API responses.
	//
	// Any of "hero", "lifestyle", "on_model", "detail", "scale_reference",
	// "angle_view", "flat_lay", "in_use", "packaging", "size_chart",
	// "product_information", "merchant_information".
	ShotType string `json:"shot_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL          respjson.Field
		AltText      respjson.Field
		IsMainImage  respjson.Field
		PhotoQuality respjson.Field
		ShotType     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnrichEnrichURLResponseImage) RawJSON() string { return r.JSON.raw }
func (r *EnrichEnrichURLResponseImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Wrapper for variant-interaction state on a Product.
//
// Holds `options` and `selected`. `options` represent all of the configuration
// options for the product. `selected` represents the currently selected option
// values.
type EnrichEnrichURLResponseVariants struct {
	Options  []EnrichEnrichURLResponseVariantsOption   `json:"options" api:"required"`
	Selected []EnrichEnrichURLResponseVariantsSelected `json:"selected" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Options     respjson.Field
		Selected    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnrichEnrichURLResponseVariants) RawJSON() string { return r.JSON.raw }
func (r *EnrichEnrichURLResponseVariants) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One dimension of a product family (e.g. 'Color', 'Size').
type EnrichEnrichURLResponseVariantsOption struct {
	// The name of the option (e.g. 'Color', 'Size')
	Name string `json:"name" api:"required"`
	// The values of the option (e.g. ['Blue', 'Red', 'Green'])
	Values []EnrichEnrichURLResponseVariantsOptionValue `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnrichEnrichURLResponseVariantsOption) RawJSON() string { return r.JSON.raw }
func (r *EnrichEnrichURLResponseVariantsOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One value of one variant option (e.g. 'Blue' under 'Color')
type EnrichEnrichURLResponseVariantsOptionValue struct {
	// Whether the option value exists on the product, or is a configuration only
	// present on another variant of the same product. For example, a shirt that comes
	// in multiple colors, but only one color is available in Size XL.
	Exists bool `json:"exists" api:"required"`
	// The display value of the option value (e.g. 'Blue')
	Label string `json:"label" api:"required"`
	// The availability status of the option value. None when returned on search
	// results, hydrated only on get product detail requests.
	//
	// Any of "InStock", "LimitedAvailability", "PreOrder", "BackOrder", "SoldOut",
	// "OutOfStock", "Discontinued", "Unknown".
	Available AvailabilityStatus `json:"available" api:"nullable"`
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
func (r EnrichEnrichURLResponseVariantsOptionValue) RawJSON() string { return r.JSON.raw }
func (r *EnrichEnrichURLResponseVariantsOptionValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One effective selection on a product, post server-side relaxation.
type EnrichEnrichURLResponseVariantsSelected struct {
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
func (r EnrichEnrichURLResponseVariantsSelected) RawJSON() string { return r.JSON.raw }
func (r *EnrichEnrichURLResponseVariantsSelected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnrichEnrichURLParams struct {
	EnrichRequest EnrichRequestParam
	paramObj
}

func (r EnrichEnrichURLParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EnrichRequest)
}
func (r *EnrichEnrichURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
