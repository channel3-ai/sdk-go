// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/channel3-ai/sdk-go/internal/apijson"
	"github.com/channel3-ai/sdk-go/internal/requestconfig"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/channel3-ai/sdk-go/packages/param"
	"github.com/channel3-ai/sdk-go/packages/respjson"
)

// ConversationService contains methods and other services that help with
// interacting with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationService] method instead.
type ConversationService struct {
	options []option.RequestOption
	Items   ConversationItemService
}

// NewConversationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConversationService(opts ...option.RequestOption) (r ConversationService) {
	r = ConversationService{}
	r.options = opts
	r.Items = NewConversationItemService(opts...)
	return
}

// Return metadata for a conversation thread.
func (r *ConversationService) Get(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/conversations/%s", url.PathEscape(conversationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Conversation struct {
	ID        string `json:"id" api:"required"`
	CreatedAt int64  `json:"created_at" api:"required"`
	// Partner-supplied context pinned to the top of a conversation thread.
	Context ConversationContext `json:"context" api:"nullable"`
	// Free-form key/value pairs the caller attached to the thread.
	Metadata map[string]any `json:"metadata"`
	UserID   string         `json:"user_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Context     respjson.Field
		Metadata    respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Conversation) RawJSON() string { return r.JSON.raw }
func (r *Conversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Partner-supplied context pinned to the top of a conversation thread.
type ConversationContext struct {
	// What platform or surface is hosting this conversation.
	ApplicationContext string `json:"application_context" api:"nullable"`
	// Who the conversation is with (profile, preferences, session facts).
	UserContext string `json:"user_context" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationContext respjson.Field
		UserContext        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationContext) RawJSON() string { return r.JSON.raw }
func (r *ConversationContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationContext to a ConversationContextParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationContextParam.Overrides()
func (r ConversationContext) ToParam() ConversationContextParam {
	return param.Override[ConversationContextParam](json.RawMessage(r.RawJSON()))
}

// Partner-supplied context pinned to the top of a conversation thread.
type ConversationContextParam struct {
	// What platform or surface is hosting this conversation.
	ApplicationContext param.Opt[string] `json:"application_context,omitzero"`
	// Who the conversation is with (profile, preferences, session facts).
	UserContext param.Opt[string] `json:"user_context,omitzero"`
	paramObj
}

func (r ConversationContextParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationContextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationContextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemsResponse struct {
	Items []ConversationItemsResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemsResponseItem struct {
	Role  string                              `json:"role" api:"required"`
	Parts []ConversationItemsResponseItemPart `json:"parts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Parts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationItemsResponseItemPart struct {
	// Any of "text", "tool", "image".
	Type  string         `json:"type" api:"required"`
	Input map[string]any `json:"input" api:"nullable"`
	// Tool call kept in conversation history for model context only. Not streamed to
	// the UI and not exposed as a conversation item.
	ModelOnly        bool           `json:"modelOnly"`
	Output           map[string]any `json:"output" api:"nullable"`
	SuggestedReplies []string       `json:"suggestedReplies" api:"nullable"`
	Text             string         `json:"text" api:"nullable"`
	ToolCallID       string         `json:"toolCallId" api:"nullable"`
	ToolName         string         `json:"toolName" api:"nullable"`
	URL              string         `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type             respjson.Field
		Input            respjson.Field
		ModelOnly        respjson.Field
		Output           respjson.Field
		SuggestedReplies respjson.Field
		Text             respjson.Field
		ToolCallID       respjson.Field
		ToolName         respjson.Field
		URL              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationItemsResponseItemPart) RawJSON() string { return r.JSON.raw }
func (r *ConversationItemsResponseItemPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
