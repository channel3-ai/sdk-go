// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package publicsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/public-sdk-go"
	"github.com/stainless-sdks/public-sdk-go/internal/testutil"
	"github.com/stainless-sdks/public-sdk-go/option"
)

func TestEnrichEnrichURL(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := publicsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Enrich.EnrichURL(context.TODO(), publicsdk.EnrichEnrichURLParams{
		EnrichRequest: publicsdk.EnrichRequestParam{
			URL: "url",
		},
	})
	if err != nil {
		var apierr *publicsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
