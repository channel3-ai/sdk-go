// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/channel3-ai/sdk-go"
	"github.com/channel3-ai/sdk-go/internal/testutil"
	"github.com/channel3-ai/sdk-go/option"
)

func TestConversationNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := channel3go.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Conversations.New(context.TODO(), channel3go.ConversationNewParams{
		CreateTurnRequest: channel3go.CreateTurnRequestParam{
			Message: channel3go.UserMessageParam{
				Parts: []channel3go.UserMessagePartUnionParam{{
					OfText: &channel3go.TextPartParam{
						Text: "text",
						Type: channel3go.TextPartTypeText,
					},
				}},
				Role: channel3go.UserMessageRoleUser,
			},
			Context: channel3go.ConversationContextParam{
				ApplicationContext: channel3go.String("application_context"),
				UserContext:        channel3go.String("user_context"),
			},
			ConversationID: channel3go.String("conversation_id"),
			Filters: channel3go.SearchFiltersParam{
				Age: []string{"newborn"},
				Attributes: map[string][]string{
					"foo": {"string"},
				},
				Availability: []channel3go.AvailabilityStatus{channel3go.AvailabilityStatusInStock},
				BrandIDs:     []string{"string"},
				CategoryIDs:  []string{"string"},
				Colors: channel3go.SearchFiltersColorsParam{
					Palette: []channel3go.SearchFiltersColorsPaletteParam{{
						Hex:        "hex",
						Percentage: channel3go.Float(0),
					}},
					Match: "strict",
				},
				Conditions: []string{"new"},
				Dimensions: channel3go.SearchFiltersDimensionsParam{
					Height: channel3go.SearchFiltersDimensionsHeightParam{
						Unit: "mm",
						Max:  channel3go.Float(0),
						Min:  channel3go.Float(0),
					},
					Length: channel3go.SearchFiltersDimensionsLengthParam{
						Unit: "mm",
						Max:  channel3go.Float(0),
						Min:  channel3go.Float(0),
					},
					Weight: channel3go.SearchFiltersDimensionsWeightParam{
						Unit: "mg",
						Max:  channel3go.Float(0),
						Min:  channel3go.Float(0),
					},
					Width: channel3go.SearchFiltersDimensionsWidthParam{
						Unit: "mm",
						Max:  channel3go.Float(0),
						Min:  channel3go.Float(0),
					},
				},
				ExcludeBrandIDs:    []string{"string"},
				ExcludeCategoryIDs: []string{"string"},
				ExcludeWebsiteIDs:  []string{"string"},
				Gender:             channel3go.SearchFiltersGenderMale,
				Price: channel3go.SearchFilterPriceParam{
					MaxPrice: channel3go.Float(0),
					MinPrice: channel3go.Float(0),
				},
				Sale:       channel3go.SearchFiltersSaleOnSale,
				WebsiteIDs: []string{"string"},
			},
			Stream: channel3go.Bool(true),
		},
		XUserID: channel3go.String("x-user-id"),
	})
	if err != nil {
		var apierr *channel3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConversationGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := channel3go.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Conversations.Get(
		context.TODO(),
		"conversation_id",
		channel3go.ConversationGetParams{
			Cursor: channel3go.String("cursor"),
			Limit:  channel3go.Int(1),
		},
	)
	if err != nil {
		var apierr *channel3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
