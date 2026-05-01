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

func TestCategoryGet(t *testing.T) {
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
	_, err := client.Categories.Get(context.TODO(), "slug")
	if err != nil {
		var apierr *channel3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCategoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Categories.List(context.TODO(), channel3go.CategoryListParams{
		Page:      channel3go.Int(1),
		PageSize:  channel3go.Int(1),
		RootsOnly: channel3go.Bool(true),
	})
	if err != nil {
		var apierr *channel3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCategorySearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Categories.Search(context.TODO(), channel3go.CategorySearchParams{
		Query: "x",
		Limit: channel3go.Int(1),
	})
	if err != nil {
		var apierr *channel3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
