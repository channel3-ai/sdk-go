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

// ConversationClientTokenService contains methods and other services that help
// with interacting with the channel3 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversationClientTokenService] method instead.
type ConversationClientTokenService struct {
	options []option.RequestOption
}

// NewConversationClientTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConversationClientTokenService(opts ...option.RequestOption) (r ConversationClientTokenService) {
	r = ConversationClientTokenService{}
	r.options = opts
	return
}

// Mint a short-lived, browser-safe token. With `conversation_id` the token
// continues and reads that thread; without it, the token's first turn creates the
// thread and binds the token to it.
func (r *ConversationClientTokenService) New(ctx context.Context, body ConversationClientTokenNewParams, opts ...option.RequestOption) (res *ClientTokenResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/conversations/client_tokens"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Revoke a client token immediately. The token travels in the request body, not
// the URL, so it stays out of access logs; only the minting vendor can revoke it.
func (r *ConversationClientTokenService) Revoke(ctx context.Context, body ConversationClientTokenRevokeParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/conversations/client_tokens/revoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type ClientTokenResponse struct {
	Token     string `json:"token" api:"required"`
	ExpiresAt int64  `json:"expires_at" api:"required"`
	TokenID   string `json:"token_id" api:"required"`
	// Any of "Bearer".
	TokenType ClientTokenResponseTokenType `json:"token_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExpiresAt   respjson.Field
		TokenID     respjson.Field
		TokenType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientTokenResponseTokenType string

const (
	ClientTokenResponseTokenTypeBearer ClientTokenResponseTokenType = "Bearer"
)

type CreateClientTokenRequestParam struct {
	ConversationID param.Opt[string] `json:"conversation_id,omitzero"`
	TtlSeconds     param.Opt[int64]  `json:"ttl_seconds,omitzero"`
	paramObj
}

func (r CreateClientTokenRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateClientTokenRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateClientTokenRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Token is required.
type RevokeClientTokenRequestParam struct {
	Token string `json:"token" api:"required"`
	paramObj
}

func (r RevokeClientTokenRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RevokeClientTokenRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RevokeClientTokenRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationClientTokenNewParams struct {
	CreateClientTokenRequest CreateClientTokenRequestParam
	paramObj
}

func (r ConversationClientTokenNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateClientTokenRequest)
}
func (r *ConversationClientTokenNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationClientTokenRevokeParams struct {
	RevokeClientTokenRequest RevokeClientTokenRequestParam
	paramObj
}

func (r ConversationClientTokenRevokeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RevokeClientTokenRequest)
}
func (r *ConversationClientTokenRevokeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
