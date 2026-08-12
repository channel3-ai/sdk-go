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
	"github.com/channel3-ai/sdk-go/internal/apiquery"
	shimjson "github.com/channel3-ai/sdk-go/internal/encoding/json"
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
}

// NewConversationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConversationService(opts ...option.RequestOption) (r ConversationService) {
	r = ConversationService{}
	r.options = opts
	return
}

// Run one conversation turn. Omit `conversation_id` to create the thread with this
// turn; pass it to continue an existing thread.
func (r *ConversationService) New(ctx context.Context, params ConversationNewParams, opts ...option.RequestOption) (res *TurnResult, err error) {
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "v1/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Thread metadata plus one page of its message history. Paginate `items` with
// `limit` and `cursor`.
func (r *ConversationService) Get(ctx context.Context, conversationID string, query ConversationGetParams, opts ...option.RequestOption) (res *ConversationDetail, err error) {
	opts = slices.Concat(r.options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/conversations/%s", url.PathEscape(conversationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AssistantMessage struct {
	Parts []AssistantMessagePartUnion `json:"parts"`
	// Any of "assistant".
	Role AssistantMessageRole `json:"role"`
	// Tap-ready follow-up messages offered after this reply.
	Suggestions []string `json:"suggestions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parts       respjson.Field
		Role        respjson.Field
		Suggestions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantMessage) RawJSON() string { return r.JSON.raw }
func (r *AssistantMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantMessagePartUnion contains all possible properties and values from
// [TextPart], [ToolPart].
//
// Use the [AssistantMessagePartUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssistantMessagePartUnion struct {
	// This field is from variant [TextPart].
	Text string `json:"text"`
	// Any of "text", "tool".
	Type string `json:"type"`
	// This field is from variant [ToolPart].
	ToolCallID string `json:"tool_call_id"`
	// This field is from variant [ToolPart].
	ToolName string `json:"tool_name"`
	// This field is from variant [ToolPart].
	Input ToolPartInputUnion `json:"input"`
	// This field is from variant [ToolPart].
	Output ToolPartOutputUnion `json:"output"`
	JSON   struct {
		Text       respjson.Field
		Type       respjson.Field
		ToolCallID respjson.Field
		ToolName   respjson.Field
		Input      respjson.Field
		Output     respjson.Field
		raw        string
	} `json:"-"`
}

// anyAssistantMessagePart is implemented by each variant of
// [AssistantMessagePartUnion] to add type safety for the return type of
// [AssistantMessagePartUnion.AsAny]
type anyAssistantMessagePart interface {
	implAssistantMessagePartUnion()
}

func (TextPart) implAssistantMessagePartUnion() {}
func (ToolPart) implAssistantMessagePartUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AssistantMessagePartUnion.AsAny().(type) {
//	case channel3go.TextPart:
//	case channel3go.ToolPart:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AssistantMessagePartUnion) AsAny() anyAssistantMessagePart {
	switch u.Type {
	case "text":
		return u.AsText()
	case "tool":
		return u.AsTool()
	}
	return nil
}

func (u AssistantMessagePartUnion) AsText() (v TextPart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantMessagePartUnion) AsTool() (v ToolPart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssistantMessagePartUnion) RawJSON() string { return u.JSON.raw }

func (r *AssistantMessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantMessageRole string

const (
	AssistantMessageRoleAssistant AssistantMessageRole = "assistant"
)

// Client-facing catalog tool result shown on the stream and on `ToolPart`.
type CatalogDisplayPayload struct {
	NextPageToken string          `json:"next_page_token" api:"nullable"`
	Products      []ProductDetail `json:"products"`
	ExtraFields   map[string]any  `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageToken respjson.Field
		Products      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogDisplayPayload) RawJSON() string { return r.JSON.raw }
func (r *CatalogDisplayPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogToolError struct {
	Error string `json:"error" api:"required"`
	// Any of true.
	IsError  bool            `json:"isError"`
	Products []ProductDetail `json:"products"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		IsError     respjson.Field
		Products    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogToolError) RawJSON() string { return r.JSON.raw }
func (r *CatalogToolError) UnmarshalJSON(data []byte) error {
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

// Thread metadata plus one page of its message history.
type ConversationDetail struct {
	ID        string                        `json:"id" api:"required"`
	CreatedAt int64                         `json:"created_at" api:"required"`
	Items     []ConversationDetailItemUnion `json:"items" api:"required"`
	// Partner-supplied context pinned to the top of a conversation thread.
	Context ConversationContext `json:"context" api:"nullable"`
	HasMore bool                `json:"has_more"`
	// Pass as `cursor` to fetch the next page. Null when no more items.
	NextCursor string `json:"next_cursor" api:"nullable"`
	UserID     string `json:"user_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Items       respjson.Field
		Context     respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationDetail) RawJSON() string { return r.JSON.raw }
func (r *ConversationDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationDetailItemUnion contains all possible properties and values from
// [UserMessage], [AssistantMessage].
//
// Use the [ConversationDetailItemUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationDetailItemUnion struct {
	// This field is a union of [[]UserMessagePartUnion], [[]AssistantMessagePartUnion]
	Parts ConversationDetailItemUnionParts `json:"parts"`
	// Any of "user", "assistant".
	Role string `json:"role"`
	// This field is from variant [AssistantMessage].
	Suggestions []string `json:"suggestions"`
	JSON        struct {
		Parts       respjson.Field
		Role        respjson.Field
		Suggestions respjson.Field
		raw         string
	} `json:"-"`
}

// anyConversationDetailItem is implemented by each variant of
// [ConversationDetailItemUnion] to add type safety for the return type of
// [ConversationDetailItemUnion.AsAny]
type anyConversationDetailItem interface {
	implConversationDetailItemUnion()
}

func (UserMessage) implConversationDetailItemUnion()      {}
func (AssistantMessage) implConversationDetailItemUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationDetailItemUnion.AsAny().(type) {
//	case channel3go.UserMessage:
//	case channel3go.AssistantMessage:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationDetailItemUnion) AsAny() anyConversationDetailItem {
	switch u.Role {
	case "user":
		return u.AsUser()
	case "assistant":
		return u.AsAssistant()
	}
	return nil
}

func (u ConversationDetailItemUnion) AsUser() (v UserMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationDetailItemUnion) AsAssistant() (v AssistantMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationDetailItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ConversationDetailItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationDetailItemUnionParts is an implicit subunion of
// [ConversationDetailItemUnion]. ConversationDetailItemUnionParts provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConversationDetailItemUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfParts]
type ConversationDetailItemUnionParts struct {
	// This field will be present if the value is a [[]UserMessagePartUnion] instead of
	// an object.
	OfParts []UserMessagePartUnion `json:",inline"`
	JSON    struct {
		OfParts respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ConversationDetailItemUnionParts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run one turn. Without `conversation_id` a new thread is created first.
//
// The property Message is required.
type CreateTurnRequestParam struct {
	Message UserMessageParam `json:"message,omitzero" api:"required"`
	// Existing thread to continue; when omitted, a new thread is created.
	ConversationID param.Opt[string] `json:"conversation_id,omitzero"`
	// Stream turn events over SSE (default) or return the assembled turn as JSON.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// Partner-supplied context pinned to the top of a conversation thread.
	Context ConversationContextParam `json:"context,omitzero"`
	// Search filters for the search API.
	Filters SearchFiltersParam `json:"filters,omitzero"`
	paramObj
}

func (r CreateTurnRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateTurnRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateTurnRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image by URL. `data:` URIs are uploaded and rewritten server-side.
type ImagePart struct {
	URL string `json:"url" api:"required"`
	// Any of "image".
	Type ImagePartType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImagePart) RawJSON() string { return r.JSON.raw }
func (r *ImagePart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ImagePart to a ImagePartParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ImagePartParam.Overrides()
func (r ImagePart) ToParam() ImagePartParam {
	return param.Override[ImagePartParam](json.RawMessage(r.RawJSON()))
}

type ImagePartType string

const (
	ImagePartTypeImage ImagePartType = "image"
)

// An image by URL. `data:` URIs are uploaded and rewritten server-side.
//
// The property URL is required.
type ImagePartParam struct {
	URL string `json:"url" api:"required"`
	// Any of "image".
	Type ImagePartType `json:"type,omitzero"`
	paramObj
}

func (r ImagePartParam) MarshalJSON() (data []byte, err error) {
	type shadow ImagePartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImagePartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProductIDsInput struct {
	ProductIDs []string `json:"product_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductIDs  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductIDsInput) RawJSON() string { return r.JSON.raw }
func (r *ProductIDsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchProductsInput struct {
	Query string `json:"query" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Query       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchProductsInput) RawJSON() string { return r.JSON.raw }
func (r *SearchProductsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextPart struct {
	Text string `json:"text" api:"required"`
	// Any of "text".
	Type TextPartType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextPart) RawJSON() string { return r.JSON.raw }
func (r *TextPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TextPart to a TextPartParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TextPartParam.Overrides()
func (r TextPart) ToParam() TextPartParam {
	return param.Override[TextPartParam](json.RawMessage(r.RawJSON()))
}

type TextPartType string

const (
	TextPartTypeText TextPartType = "text"
)

// The property Text is required.
type TextPartParam struct {
	Text string `json:"text" api:"required"`
	// Any of "text".
	Type TextPartType `json:"type,omitzero"`
	paramObj
}

func (r TextPartParam) MarshalJSON() (data []byte, err error) {
	type shadow TextPartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextPartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One catalog tool call and its display payload from an assistant turn.
type ToolPart struct {
	ToolCallID string             `json:"tool_call_id" api:"required"`
	ToolName   string             `json:"tool_name" api:"required"`
	Input      ToolPartInputUnion `json:"input"`
	// Client-facing catalog tool result shown on the stream and on `ToolPart`.
	Output ToolPartOutputUnion `json:"output"`
	// Any of "tool".
	Type ToolPartType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Input       respjson.Field
		Output      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolPart) RawJSON() string { return r.JSON.raw }
func (r *ToolPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolPartInputUnion contains all possible properties and values from
// [SearchProductsInput], [ProductIDsInput].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ToolPartInputUnion struct {
	// This field is from variant [SearchProductsInput].
	Query string `json:"query"`
	// This field is from variant [ProductIDsInput].
	ProductIDs []string `json:"product_ids"`
	JSON       struct {
		Query      respjson.Field
		ProductIDs respjson.Field
		raw        string
	} `json:"-"`
}

func (u ToolPartInputUnion) AsSearchProductsInput() (v SearchProductsInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolPartInputUnion) AsProductIDsInput() (v ProductIDsInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolPartInputUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolPartInputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolPartOutputUnion contains all possible properties and values from
// [CatalogDisplayPayload], [CatalogToolError].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ToolPartOutputUnion struct {
	// This field is from variant [CatalogDisplayPayload].
	NextPageToken string          `json:"next_page_token"`
	Products      []ProductDetail `json:"products"`
	// This field is from variant [CatalogToolError].
	Error string `json:"error"`
	// This field is from variant [CatalogToolError].
	IsError bool `json:"isError"`
	JSON    struct {
		NextPageToken respjson.Field
		Products      respjson.Field
		Error         respjson.Field
		IsError       respjson.Field
		raw           string
	} `json:"-"`
}

func (u ToolPartOutputUnion) AsCatalogDisplayPayload() (v CatalogDisplayPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolPartOutputUnion) AsCatalogToolError() (v CatalogToolError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolPartOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolPartOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolPartType string

const (
	ToolPartTypeTool ToolPartType = "tool"
)

// Buffered equivalent of a streamed turn (`stream: false`).
type TurnResult struct {
	ConversationID string           `json:"conversation_id" api:"required"`
	Message        AssistantMessage `json:"message" api:"required"`
	TurnID         string           `json:"turn_id" api:"required"`
	Usage          TurnUsage        `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversationID respjson.Field
		Message        respjson.Field
		TurnID         respjson.Field
		Usage          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnResult) RawJSON() string { return r.JSON.raw }
func (r *TurnResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TurnUsage struct {
	// API credits charged for this turn (turn fee plus catalog searches).
	CreditsCharged int64 `json:"credits_charged" api:"required"`
	// Catalog searches executed during this turn.
	SearchesRun int64 `json:"searches_run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditsCharged respjson.Field
		SearchesRun    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnUsage) RawJSON() string { return r.JSON.raw }
func (r *TurnUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMessage struct {
	Parts []UserMessagePartUnion `json:"parts"`
	// Any of "user".
	Role UserMessageRole `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parts       respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserMessage) RawJSON() string { return r.JSON.raw }
func (r *UserMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserMessage to a UserMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserMessageParam.Overrides()
func (r UserMessage) ToParam() UserMessageParam {
	return param.Override[UserMessageParam](json.RawMessage(r.RawJSON()))
}

// UserMessagePartUnion contains all possible properties and values from
// [TextPart], [ImagePart].
//
// Use the [UserMessagePartUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UserMessagePartUnion struct {
	// This field is from variant [TextPart].
	Text string `json:"text"`
	// Any of "text", "image".
	Type string `json:"type"`
	// This field is from variant [ImagePart].
	URL  string `json:"url"`
	JSON struct {
		Text respjson.Field
		Type respjson.Field
		URL  respjson.Field
		raw  string
	} `json:"-"`
}

// anyUserMessagePart is implemented by each variant of [UserMessagePartUnion] to
// add type safety for the return type of [UserMessagePartUnion.AsAny]
type anyUserMessagePart interface {
	implUserMessagePartUnion()
}

func (TextPart) implUserMessagePartUnion()  {}
func (ImagePart) implUserMessagePartUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UserMessagePartUnion.AsAny().(type) {
//	case channel3go.TextPart:
//	case channel3go.ImagePart:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UserMessagePartUnion) AsAny() anyUserMessagePart {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image":
		return u.AsImage()
	}
	return nil
}

func (u UserMessagePartUnion) AsText() (v TextPart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserMessagePartUnion) AsImage() (v ImagePart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserMessagePartUnion) RawJSON() string { return u.JSON.raw }

func (r *UserMessagePartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMessageRole string

const (
	UserMessageRoleUser UserMessageRole = "user"
)

type UserMessageParam struct {
	Parts []UserMessagePartUnionParam `json:"parts,omitzero"`
	// Any of "user".
	Role UserMessageRole `json:"role,omitzero"`
	paramObj
}

func (r UserMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow UserMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UserMessagePartUnionParam struct {
	OfText  *TextPartParam  `json:",omitzero,inline"`
	OfImage *ImagePartParam `json:",omitzero,inline"`
	paramUnion
}

func (u UserMessagePartUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImage)
}
func (u *UserMessagePartUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[UserMessagePartUnionParam](
		"type",
		apijson.Discriminator[TextPartParam]("text"),
		apijson.Discriminator[ImagePartParam]("image"),
	)
}

type ConversationNewParams struct {
	// Run one turn. Without `conversation_id` a new thread is created first.
	CreateTurnRequest CreateTurnRequestParam
	// Optional user identifier to attribute clicks and sales to a user in your system.
	// Channel3 appends it to buy URLs in the response.
	XUserID param.Opt[string] `header:"x-user-id,omitzero" json:"-"`
	paramObj
}

func (r ConversationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateTurnRequest)
}
func (r *ConversationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationGetParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConversationGetParams]'s query parameters as `url.Values`.
func (r ConversationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
