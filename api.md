# Products

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#AvailabilityStatus">AvailabilityStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ImageSearchRequestParam">ImageSearchRequestParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#LocaleConfigParam">LocaleConfigParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#LookupRequestParam">LookupRequestParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SimilarProductsRequestParam">SimilarProductsRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#LookupResponse">LookupResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Price">Price</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductBrand">ProductBrand</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductDetail">ProductDetail</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductImage">ProductImage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductOffer">ProductOffer</a>

Methods:

- <code title="get /v1/products/{product_id}">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, productID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductGetParams">ProductGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductDetail">ProductDetail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/similar">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductService.FindSimilar">FindSimilar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductFindSimilarParams">ProductFindSimilarParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination#SearchPage">SearchPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductDetail">ProductDetail</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/lookup">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductLookupParams">ProductLookupParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#LookupResponse">LookupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/search">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductSearchParams">ProductSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination#SearchPage">SearchPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductDetail">ProductDetail</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/image-search">client.Products.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductService.SearchByImage">SearchByImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductSearchByImageParams">ProductSearchByImageParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination#SearchPage">SearchPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#ProductDetail">ProductDetail</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Brands

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Brand">Brand</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchBrandsResponse">SearchBrandsResponse</a>

Methods:

- <code title="get /v1/brands/{brand_id}">client.Brands.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#BrandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, brandID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Brand">Brand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/brands">client.Brands.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#BrandService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#BrandListParams">BrandListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Brand">Brand</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/brands">client.Brands.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#BrandService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#BrandFindParams">BrandFindParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Brand">Brand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/brands/search">client.Brands.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#BrandService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#BrandSearchParams">BrandSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchBrandsResponse">SearchBrandsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Categories

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Category">Category</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#CategoryAttribute">CategoryAttribute</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#CategoryRef">CategoryRef</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#CategorySummary">CategorySummary</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PaginatedListCategoriesResponse">PaginatedListCategoriesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchCategoriesResponse">SearchCategoriesResponse</a>

Methods:

- <code title="get /v1/categories/{slug}">client.Categories.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#CategoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, slug <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Category">Category</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/categories">client.Categories.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#CategoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#CategoryListParams">CategoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination#CategoryPage">CategoryPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#CategorySummary">CategorySummary</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/categories/search">client.Categories.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#CategoryService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#CategorySearchParams">CategorySearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchCategoriesResponse">SearchCategoriesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Websites

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Website">Website</a>

Methods:

- <code title="get /v0/websites">client.Websites.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#WebsiteService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#WebsiteGetParams">WebsiteGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Website">Website</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PriceTracking

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#StartTrackingRequestParam">StartTrackingRequestParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#StopTrackingRequestParam">StopTrackingRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PaginatedSubscriptionsResponse">PaginatedSubscriptionsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceHistory">PriceHistory</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceHistoryPoint">PriceHistoryPoint</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceStatistics">PriceStatistics</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Subscription">Subscription</a>

Methods:

- <code title="get /v0/price-tracking/subscriptions">client.PriceTracking.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceTrackingService.ListSubscriptions">ListSubscriptions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceTrackingListSubscriptionsParams">PriceTrackingListSubscriptionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Subscription">Subscription</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/price-tracking/history/{canonical_product_id}">client.PriceTracking.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceTrackingService.GetHistory">GetHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, canonicalProductID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceTrackingGetHistoryParams">PriceTrackingGetHistoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceHistory">PriceHistory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/price-tracking/start">client.PriceTracking.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceTrackingService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceTrackingStartParams">PriceTrackingStartParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/price-tracking/stop">client.PriceTracking.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceTrackingService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#PriceTrackingStopParams">PriceTrackingStopParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Search

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchConfigParam">SearchConfigParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchFilterPriceParam">SearchFilterPriceParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchFiltersParam">SearchFiltersParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchRequestParam">SearchRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchResponse">SearchResponse</a>

Methods:

- <code title="post /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchService.Perform">Perform</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchPerformParams">SearchPerformParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Enrich

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#EnrichRequestParam">EnrichRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#EnrichEnrichURLResponse">EnrichEnrichURLResponse</a>

Methods:

- <code title="post /v0/enrich">client.Enrich.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#EnrichService.EnrichURL">EnrichURL</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#EnrichEnrichURLParams">EnrichEnrichURLParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go">publicsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/public-sdk-go#EnrichEnrichURLResponse">EnrichEnrichURLResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
