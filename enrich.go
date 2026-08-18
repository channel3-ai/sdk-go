// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"github.com/channel3-ai/sdk-go/option"
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
