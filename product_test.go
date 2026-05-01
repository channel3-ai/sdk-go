// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package channel3go_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/public-sdk-go"
	"github.com/stainless-sdks/public-sdk-go/internal/testutil"
	"github.com/stainless-sdks/public-sdk-go/option"
)

func TestProductGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.Get(
		context.TODO(),
		"product_id",
		channel3go.ProductGetParams{
			Country:    channel3go.ProductGetParamsCountryUs,
			Currency:   channel3go.ProductGetParamsCurrencyUsd,
			Language:   channel3go.ProductGetParamsLanguageEn,
			WebsiteIDs: []string{"string"},
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

func TestProductFindSimilarWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.FindSimilar(context.TODO(), channel3go.ProductFindSimilarParams{
		SimilarProductsRequest: channel3go.SimilarProductsRequestParam{
			ProductID: "product_id",
			Config: channel3go.LocaleConfigParam{
				Country:  channel3go.LocaleConfigCountryUs,
				Currency: channel3go.LocaleConfigCurrencyUsd,
				Language: channel3go.LocaleConfigLanguageEn,
			},
			Filters: channel3go.SearchFiltersParam{
				Age:                []string{"newborn"},
				Availability:       []channel3go.AvailabilityStatus{channel3go.AvailabilityStatusInStock},
				BrandIDs:           []string{"string"},
				CategoryIDs:        []string{"string"},
				Condition:          channel3go.SearchFiltersConditionNew,
				ExcludeBrandIDs:    []string{"string"},
				ExcludeCategoryIDs: []string{"string"},
				ExcludeWebsiteIDs:  []string{"string"},
				Gender:             channel3go.SearchFiltersGenderMale,
				Price: channel3go.SearchFilterPriceParam{
					MaxPrice: channel3go.Float(0),
					MinPrice: channel3go.Float(0),
				},
				WebsiteIDs: []string{"string"},
			},
			Limit:     channel3go.Int(1),
			PageToken: channel3go.String("page_token"),
		},
	})
	if err != nil {
		var apierr *channel3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductLookupWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.Lookup(context.TODO(), channel3go.ProductLookupParams{
		LookupRequest: channel3go.LookupRequestParam{
			URL:               "url",
			MaxStalenessHours: channel3go.Int(1),
		},
	})
	if err != nil {
		var apierr *channel3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.Search(context.TODO(), channel3go.ProductSearchParams{
		SearchRequest: channel3go.SearchRequestParam{
			Base64Image: channel3go.String("base64_image"),
			Config: channel3go.SearchConfigParam{
				Country:           channel3go.SearchConfigCountryUs,
				Currency:          channel3go.SearchConfigCurrencyUsd,
				KeywordSearchOnly: channel3go.Bool(true),
				Language:          channel3go.SearchConfigLanguageEn,
			},
			Filters: channel3go.SearchFiltersParam{
				Age:                []string{"newborn"},
				Availability:       []channel3go.AvailabilityStatus{channel3go.AvailabilityStatusInStock},
				BrandIDs:           []string{"string"},
				CategoryIDs:        []string{"string"},
				Condition:          channel3go.SearchFiltersConditionNew,
				ExcludeBrandIDs:    []string{"string"},
				ExcludeCategoryIDs: []string{"string"},
				ExcludeWebsiteIDs:  []string{"string"},
				Gender:             channel3go.SearchFiltersGenderMale,
				Price: channel3go.SearchFilterPriceParam{
					MaxPrice: channel3go.Float(0),
					MinPrice: channel3go.Float(0),
				},
				WebsiteIDs: []string{"string"},
			},
			ImageURL:  channel3go.String("image_url"),
			Limit:     channel3go.Int(1),
			PageToken: channel3go.String("page_token"),
			Query:     channel3go.String("query"),
		},
	})
	if err != nil {
		var apierr *channel3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductSearchByImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.SearchByImage(context.TODO(), channel3go.ProductSearchByImageParams{
		ImageSearchRequest: channel3go.ImageSearchRequestParam{
			Base64Image: channel3go.String("base64_image"),
			Config: channel3go.LocaleConfigParam{
				Country:  channel3go.LocaleConfigCountryUs,
				Currency: channel3go.LocaleConfigCurrencyUsd,
				Language: channel3go.LocaleConfigLanguageEn,
			},
			Filters: channel3go.SearchFiltersParam{
				Age:                []string{"newborn"},
				Availability:       []channel3go.AvailabilityStatus{channel3go.AvailabilityStatusInStock},
				BrandIDs:           []string{"string"},
				CategoryIDs:        []string{"string"},
				Condition:          channel3go.SearchFiltersConditionNew,
				ExcludeBrandIDs:    []string{"string"},
				ExcludeCategoryIDs: []string{"string"},
				ExcludeWebsiteIDs:  []string{"string"},
				Gender:             channel3go.SearchFiltersGenderMale,
				Price: channel3go.SearchFilterPriceParam{
					MaxPrice: channel3go.Float(0),
					MinPrice: channel3go.Float(0),
				},
				WebsiteIDs: []string{"string"},
			},
			ImageURL:  channel3go.String("image_url"),
			Limit:     channel3go.Int(1),
			PageToken: channel3go.String("page_token"),
		},
	})
	if err != nil {
		var apierr *channel3go.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
