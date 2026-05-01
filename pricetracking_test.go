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

func TestPriceTrackingGetHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.PriceTracking.GetHistory(
		context.TODO(),
		"canonical_product_id",
		publicsdk.PriceTrackingGetHistoryParams{
			Days: publicsdk.Int(1),
		},
	)
	if err != nil {
		var apierr *publicsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPriceTrackingListSubscriptionsWithOptionalParams(t *testing.T) {
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
	_, err := client.PriceTracking.ListSubscriptions(context.TODO(), publicsdk.PriceTrackingListSubscriptionsParams{
		Cursor: publicsdk.String("cursor"),
		Limit:  publicsdk.Int(1),
	})
	if err != nil {
		var apierr *publicsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPriceTrackingGetHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.PriceTracking.GetHistory(
		context.TODO(),
		"canonical_product_id",
		publicsdk.PriceTrackingGetHistoryParams{
			Days: publicsdk.Int(1),
		},
	)
	if err != nil {
		var apierr *publicsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPriceTrackingStart(t *testing.T) {
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
	_, err := client.PriceTracking.Start(context.TODO(), publicsdk.PriceTrackingStartParams{
		StartTrackingRequest: publicsdk.StartTrackingRequestParam{
			CanonicalProductID: "canonical_product_id",
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

func TestPriceTrackingStop(t *testing.T) {
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
	_, err := client.PriceTracking.Stop(context.TODO(), publicsdk.PriceTrackingStopParams{
		StopTrackingRequest: publicsdk.StopTrackingRequestParam{
			CanonicalProductID: "canonical_product_id",
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
