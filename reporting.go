// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"github.com/channel3-ai/sdk-go/internal/apijson"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/channel3-ai/sdk-go/packages/respjson"
)

// ReportingService contains methods and other services that help with interacting
// with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportingService] method instead.
type ReportingService struct {
	options      []option.RequestOption
	Clicks       ReportingClickService
	Transactions ReportingTransactionService
}

// NewReportingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReportingService(opts ...option.RequestOption) (r ReportingService) {
	r = ReportingService{}
	r.options = opts
	r.Clicks = NewReportingClickService(opts...)
	r.Transactions = NewReportingTransactionService(opts...)
	return
}

// Compact product reference on click/transaction items.
type AffiliateProduct struct {
	// Canonical product ID.
	ID string `json:"id" api:"required"`
	// Product image URL.
	ImageURL string `json:"image_url" api:"nullable"`
	// Product title.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ImageURL    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AffiliateProduct) RawJSON() string { return r.JSON.raw }
func (r *AffiliateProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
