// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/channel3-ai/sdk-go/internal/apijson"
	"github.com/channel3-ai/sdk-go/internal/apiquery"
	"github.com/channel3-ai/sdk-go/internal/requestconfig"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/channel3-ai/sdk-go/packages/respjson"
)

// WebsiteService contains methods and other services that help with interacting
// with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebsiteService] method instead.
type WebsiteService struct {
	options []option.RequestOption
}

// NewWebsiteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebsiteService(opts ...option.RequestOption) (r WebsiteService) {
	r = WebsiteService{}
	r.options = opts
	return
}

// Resolve a website URL to its ID and best_commission_rate. Tip: website_ids
// filters accept domains directly, so this lookup is most useful for retrieving
// commission rates.
func (r *WebsiteService) Get(ctx context.Context, query WebsiteGetParams, opts ...option.RequestOption) (res *Website, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v0/websites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Resolve a website URL to its ID and best_commission_rate. Tip: website_ids
// filters accept domains directly, so this lookup is most useful for retrieving
// commission rates.
//
// Deprecated: use `retrieve` instead; will be removed in the next major version
func (r *WebsiteService) Find(ctx context.Context, body WebsiteFindParams, opts ...option.RequestOption) (res *Website, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v0/websites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, body, &res, opts...)
	return res, err
}

type Website struct {
	ID  string `json:"id" api:"required"`
	URL string `json:"url" api:"required"`
	// The maximum commission rate for the website in the requested country (default
	// 'US'), as a percentage
	BestCommissionRate float64 `json:"best_commission_rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		URL                respjson.Field
		BestCommissionRate respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Website) RawJSON() string { return r.JSON.raw }
func (r *Website) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebsiteGetParams struct {
	Query string `query:"query" api:"required" json:"-"`
	// ISO 3166-1 alpha-2 country code that `best_commission_rate` is scoped to.
	// Defaults to 'US' when unset.
	//
	// Any of "US", "GB", "EU", "AU", "CA", "IE", "DE", "AT", "FR", "BE", "IT", "ES",
	// "NL", "SE", "FI", "PT", "CZ", "GR", "RO".
	Country WebsiteGetParamsCountry `query:"country,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebsiteGetParams]'s query parameters as `url.Values`.
func (r WebsiteGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ISO 3166-1 alpha-2 country code that `best_commission_rate` is scoped to.
// Defaults to 'US' when unset.
type WebsiteGetParamsCountry string

const (
	WebsiteGetParamsCountryUs WebsiteGetParamsCountry = "US"
	WebsiteGetParamsCountryGB WebsiteGetParamsCountry = "GB"
	WebsiteGetParamsCountryEu WebsiteGetParamsCountry = "EU"
	WebsiteGetParamsCountryAu WebsiteGetParamsCountry = "AU"
	WebsiteGetParamsCountryCa WebsiteGetParamsCountry = "CA"
	WebsiteGetParamsCountryIe WebsiteGetParamsCountry = "IE"
	WebsiteGetParamsCountryDe WebsiteGetParamsCountry = "DE"
	WebsiteGetParamsCountryAt WebsiteGetParamsCountry = "AT"
	WebsiteGetParamsCountryFr WebsiteGetParamsCountry = "FR"
	WebsiteGetParamsCountryBe WebsiteGetParamsCountry = "BE"
	WebsiteGetParamsCountryIt WebsiteGetParamsCountry = "IT"
	WebsiteGetParamsCountryEs WebsiteGetParamsCountry = "ES"
	WebsiteGetParamsCountryNl WebsiteGetParamsCountry = "NL"
	WebsiteGetParamsCountrySe WebsiteGetParamsCountry = "SE"
	WebsiteGetParamsCountryFi WebsiteGetParamsCountry = "FI"
	WebsiteGetParamsCountryPt WebsiteGetParamsCountry = "PT"
	WebsiteGetParamsCountryCz WebsiteGetParamsCountry = "CZ"
	WebsiteGetParamsCountryGr WebsiteGetParamsCountry = "GR"
	WebsiteGetParamsCountryRo WebsiteGetParamsCountry = "RO"
)

type WebsiteFindParams struct {
	Query string `query:"query" api:"required" json:"-"`
	// ISO 3166-1 alpha-2 country code that `best_commission_rate` is scoped to.
	// Defaults to 'US' when unset.
	//
	// Any of "US", "GB", "EU", "AU", "CA", "IE", "DE", "AT", "FR", "BE", "IT", "ES",
	// "NL", "SE", "FI", "PT", "CZ", "GR", "RO".
	Country WebsiteFindParamsCountry `query:"country,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebsiteFindParams]'s query parameters as `url.Values`.
func (r WebsiteFindParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ISO 3166-1 alpha-2 country code that `best_commission_rate` is scoped to.
// Defaults to 'US' when unset.
type WebsiteFindParamsCountry string

const (
	WebsiteFindParamsCountryUs WebsiteFindParamsCountry = "US"
	WebsiteFindParamsCountryGB WebsiteFindParamsCountry = "GB"
	WebsiteFindParamsCountryEu WebsiteFindParamsCountry = "EU"
	WebsiteFindParamsCountryAu WebsiteFindParamsCountry = "AU"
	WebsiteFindParamsCountryCa WebsiteFindParamsCountry = "CA"
	WebsiteFindParamsCountryIe WebsiteFindParamsCountry = "IE"
	WebsiteFindParamsCountryDe WebsiteFindParamsCountry = "DE"
	WebsiteFindParamsCountryAt WebsiteFindParamsCountry = "AT"
	WebsiteFindParamsCountryFr WebsiteFindParamsCountry = "FR"
	WebsiteFindParamsCountryBe WebsiteFindParamsCountry = "BE"
	WebsiteFindParamsCountryIt WebsiteFindParamsCountry = "IT"
	WebsiteFindParamsCountryEs WebsiteFindParamsCountry = "ES"
	WebsiteFindParamsCountryNl WebsiteFindParamsCountry = "NL"
	WebsiteFindParamsCountrySe WebsiteFindParamsCountry = "SE"
	WebsiteFindParamsCountryFi WebsiteFindParamsCountry = "FI"
	WebsiteFindParamsCountryPt WebsiteFindParamsCountry = "PT"
	WebsiteFindParamsCountryCz WebsiteFindParamsCountry = "CZ"
	WebsiteFindParamsCountryGr WebsiteFindParamsCountry = "GR"
	WebsiteFindParamsCountryRo WebsiteFindParamsCountry = "RO"
)
