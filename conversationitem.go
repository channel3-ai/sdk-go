// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/channel3-ai/sdk-go/internal/requestconfig"
	"github.com/channel3-ai/sdk-go/option"
)

// ConversationItemService contains methods and other services that help with
// interacting with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationItemService] method instead.
type ConversationItemService struct {
	options []option.RequestOption
}

// NewConversationItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConversationItemService(opts ...option.RequestOption) (r ConversationItemService) {
	r = ConversationItemService{}
	r.options = opts
	return
}

// Return persisted messages for a conversation.
func (r *ConversationItemService) List(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *ConversationItemsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/conversations/%s/items", url.PathEscape(conversationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
