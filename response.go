// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/channel3-ai/sdk-go/internal/apijson"
	shimjson "github.com/channel3-ai/sdk-go/internal/encoding/json"
	"github.com/channel3-ai/sdk-go/internal/requestconfig"
	"github.com/channel3-ai/sdk-go/option"
	"github.com/channel3-ai/sdk-go/packages/param"
	"github.com/channel3-ai/sdk-go/packages/ssestream"
)

// ResponseService contains methods and other services that help with interacting
// with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResponseService] method instead.
type ResponseService struct {
	options []option.RequestOption
}

// NewResponseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResponseService(opts ...option.RequestOption) (r ResponseService) {
	r = ResponseService{}
	r.options = opts
	return
}

// Run a shopping conversation turn.
func (r *ResponseService) NewStreaming(ctx context.Context, params ResponseNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[string]) {
	var (
		raw *http.Response
		err error
	)
	if !param.IsOmitted(params.XUserID) {
		opts = append(opts, option.WithHeader("x-user-id", fmt.Sprintf("%v", params.XUserID.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	path := "v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &raw, opts...)
	return ssestream.NewStream[string](ssestream.NewDecoder(raw), err)
}

type ChatRequestParam struct {
	ConversationID param.Opt[string]            `json:"conversation_id,omitzero"`
	Debug          param.Opt[bool]              `json:"debug,omitzero"`
	Attachments    []ChatRequestAttachmentParam `json:"attachments,omitzero"`
	Image          ChatRequestImageParam        `json:"image,omitzero"`
	Message        ChatRequestMessageParam      `json:"message,omitzero"`
	// Partner-supplied context pinned to the top of a conversation thread.
	Context  ConversationContextParam  `json:"context,omitzero"`
	Messages []ChatRequestMessageParam `json:"messages,omitzero"`
	paramObj
}

func (r ChatRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type ChatRequestAttachmentParam struct {
	URL string            `json:"url" api:"required"`
	Key param.Opt[string] `json:"key,omitzero"`
	paramObj
}

func (r ChatRequestAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatRequestAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatRequestAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatRequestImageParam struct {
	Base64 param.Opt[string] `json:"base64,omitzero"`
	URL    param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r ChatRequestImageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatRequestImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatRequestImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Role is required.
type ChatRequestMessageParam struct {
	Role  string                        `json:"role" api:"required"`
	Parts []ChatRequestMessagePartParam `json:"parts,omitzero"`
	paramObj
}

func (r ChatRequestMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatRequestMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatRequestMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ChatRequestMessagePartParam struct {
	// Any of "text", "tool", "image".
	Type       string            `json:"type,omitzero" api:"required"`
	Text       param.Opt[string] `json:"text,omitzero"`
	ToolCallID param.Opt[string] `json:"toolCallId,omitzero"`
	ToolName   param.Opt[string] `json:"toolName,omitzero"`
	URL        param.Opt[string] `json:"url,omitzero"`
	// Tool call kept in conversation history for model context only. Not streamed to
	// the UI and not exposed as a conversation item.
	ModelOnly        param.Opt[bool] `json:"modelOnly,omitzero"`
	Input            map[string]any  `json:"input,omitzero"`
	Output           map[string]any  `json:"output,omitzero"`
	SuggestedReplies []string        `json:"suggestedReplies,omitzero"`
	paramObj
}

func (r ChatRequestMessagePartParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatRequestMessagePartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatRequestMessagePartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ChatRequestMessagePartParam](
		"type", "text", "tool", "image",
	)
}

type ResponseNewParams struct {
	ChatRequest ChatRequestParam
	// Optional user identifier to attribute clicks and sales to a user in your system.
	// Channel3 appends it to buy URLs in the response.
	XUserID param.Opt[string] `header:"x-user-id,omitzero" json:"-"`
	paramObj
}

func (r ResponseNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ChatRequest)
}
func (r *ResponseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
