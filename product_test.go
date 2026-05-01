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

func TestProductGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.Get(
		context.TODO(),
		"product_id",
		publicsdk.ProductGetParams{
			Country:    publicsdk.ProductGetParamsCountryUs,
			Currency:   publicsdk.ProductGetParamsCurrencyUsd,
			Language:   publicsdk.ProductGetParamsLanguageEn,
			WebsiteIDs: []string{"string"},
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

func TestProductFindSimilarWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.FindSimilar(context.TODO(), publicsdk.ProductFindSimilarParams{
		SimilarProductsRequest: publicsdk.SimilarProductsRequestParam{
			ProductID: "product_id",
			Config: publicsdk.LocaleConfigParam{
				Country:  publicsdk.LocaleConfigCountryUs,
				Currency: publicsdk.LocaleConfigCurrencyUsd,
				Language: publicsdk.LocaleConfigLanguageEn,
			},
			Filters: publicsdk.SearchFiltersParam{
				Age:                []string{"newborn"},
				Availability:       []publicsdk.AvailabilityStatus{publicsdk.AvailabilityStatusInStock},
				BrandIDs:           []string{"string"},
				CategoryIDs:        []string{"string"},
				Condition:          publicsdk.SearchFiltersConditionNew,
				ExcludeBrandIDs:    []string{"string"},
				ExcludeCategoryIDs: []string{"string"},
				ExcludeWebsiteIDs:  []string{"string"},
				Gender:             publicsdk.SearchFiltersGenderMale,
				Price: publicsdk.SearchFilterPriceParam{
					MaxPrice: publicsdk.Float(0),
					MinPrice: publicsdk.Float(0),
				},
				WebsiteIDs: []string{"string"},
			},
			Limit:     publicsdk.Int(1),
			PageToken: publicsdk.String("page_token"),
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

func TestProductLookupWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.Lookup(context.TODO(), publicsdk.ProductLookupParams{
		LookupRequest: publicsdk.LookupRequestParam{
			URL:               "url",
			MaxStalenessHours: publicsdk.Int(1),
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

func TestProductSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.Search(context.TODO(), publicsdk.ProductSearchParams{
		SearchRequest: publicsdk.SearchRequestParam{
			Base64Image: publicsdk.String("base64_image"),
			Config: publicsdk.SearchConfigParam{
				Country:           publicsdk.SearchConfigCountryUs,
				Currency:          publicsdk.SearchConfigCurrencyUsd,
				KeywordSearchOnly: publicsdk.Bool(true),
				Language:          publicsdk.SearchConfigLanguageEn,
			},
			Filters: publicsdk.SearchFiltersParam{
				Age:                []string{"newborn"},
				Availability:       []publicsdk.AvailabilityStatus{publicsdk.AvailabilityStatusInStock},
				BrandIDs:           []string{"string"},
				CategoryIDs:        []string{"string"},
				Condition:          publicsdk.SearchFiltersConditionNew,
				ExcludeBrandIDs:    []string{"string"},
				ExcludeCategoryIDs: []string{"string"},
				ExcludeWebsiteIDs:  []string{"string"},
				Gender:             publicsdk.SearchFiltersGenderMale,
				Price: publicsdk.SearchFilterPriceParam{
					MaxPrice: publicsdk.Float(0),
					MinPrice: publicsdk.Float(0),
				},
				WebsiteIDs: []string{"string"},
			},
			ImageURL:  publicsdk.String("image_url"),
			Limit:     publicsdk.Int(1),
			PageToken: publicsdk.String("page_token"),
			Query:     publicsdk.String("query"),
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

func TestProductSearchByImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.SearchByImage(context.TODO(), publicsdk.ProductSearchByImageParams{
		ImageSearchRequest: publicsdk.ImageSearchRequestParam{
			Base64Image: publicsdk.String("base64_image"),
			Config: publicsdk.LocaleConfigParam{
				Country:  publicsdk.LocaleConfigCountryUs,
				Currency: publicsdk.LocaleConfigCurrencyUsd,
				Language: publicsdk.LocaleConfigLanguageEn,
			},
			Filters: publicsdk.SearchFiltersParam{
				Age:                []string{"newborn"},
				Availability:       []publicsdk.AvailabilityStatus{publicsdk.AvailabilityStatusInStock},
				BrandIDs:           []string{"string"},
				CategoryIDs:        []string{"string"},
				Condition:          publicsdk.SearchFiltersConditionNew,
				ExcludeBrandIDs:    []string{"string"},
				ExcludeCategoryIDs: []string{"string"},
				ExcludeWebsiteIDs:  []string{"string"},
				Gender:             publicsdk.SearchFiltersGenderMale,
				Price: publicsdk.SearchFilterPriceParam{
					MaxPrice: publicsdk.Float(0),
					MinPrice: publicsdk.Float(0),
				},
				WebsiteIDs: []string{"string"},
			},
			ImageURL:  publicsdk.String("image_url"),
			Limit:     publicsdk.Int(1),
			PageToken: publicsdk.String("page_token"),
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
