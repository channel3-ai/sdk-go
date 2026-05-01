// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go_test

import (
	"context"
	"os"
	"testing"

	"github.com/channel3-ai/sdk-go"
	"github.com/channel3-ai/sdk-go/internal/testutil"
	"github.com/channel3-ai/sdk-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Products.SearchAutoPaging(context.TODO(), channel3go.ProductSearchParams{
		SearchRequest: channel3go.SearchRequestParam{},
	})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		product := iter.Current()
		t.Logf("%+v\n", product.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
