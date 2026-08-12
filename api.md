# Products

Params Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#BrowseRequestParam">BrowseRequestParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ImageSearchRequestParam">ImageSearchRequestParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#LocaleConfigParam">LocaleConfigParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#LookupRequestParam">LookupRequestParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#MonetizeRequestParam">MonetizeRequestParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SimilarProductsRequestParam">SimilarProductsRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#LookupResponse">LookupResponse</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#MonetizeOffer">MonetizeOffer</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#MonetizeResponse">MonetizeResponse</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Price">Price</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductBrand">ProductBrand</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductDetail">ProductDetail</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductImage">ProductImage</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductOffer">ProductOffer</a>

Methods:

- <code title="get /v1/products/{product_id}">client.Products.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, productID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductGetParams">ProductGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductDetail">ProductDetail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/browse">client.Products.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductService.Browse">Browse</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductBrowseParams">ProductBrowseParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination#SearchPage">SearchPage</a>[<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductDetail">ProductDetail</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/similar">client.Products.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductService.FindSimilar">FindSimilar</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductFindSimilarParams">ProductFindSimilarParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination#SearchPage">SearchPage</a>[<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductDetail">ProductDetail</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/lookup">client.Products.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductLookupParams">ProductLookupParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#LookupResponse">LookupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/monetize">client.Products.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductService.Monetize">Monetize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductMonetizeParams">ProductMonetizeParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#MonetizeResponse">MonetizeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/search">client.Products.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductSearchParams">ProductSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination#SearchPage">SearchPage</a>[<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductDetail">ProductDetail</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/image-search">client.Products.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductService.SearchByImage">SearchByImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductSearchByImageParams">ProductSearchByImageParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination#SearchPage">SearchPage</a>[<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductDetail">ProductDetail</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Reporting

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ReportingProduct">ReportingProduct</a>

## Clicks

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Click">Click</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ClicksResponse">ClicksResponse</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ClicksSummary">ClicksSummary</a>

Methods:

- <code title="get /v1/reporting/clicks">client.Reporting.Clicks.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ReportingClickService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ReportingClickListParams">ReportingClickListParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination#AnalyticsPage">AnalyticsPage</a>[<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Click">Click</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PublicTransactionStatus">PublicTransactionStatus</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#TransactionsResponse">TransactionsResponse</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#TransactionsSummary">TransactionsSummary</a>

Methods:

- <code title="get /v1/reporting/transactions">client.Reporting.Transactions.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ReportingTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ReportingTransactionListParams">ReportingTransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination#AnalyticsPage">AnalyticsPage</a>[<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Transaction">Transaction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Brands

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Brand">Brand</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchBrandsResponse">SearchBrandsResponse</a>

Methods:

- <code title="get /v1/brands/{brand_id}">client.Brands.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#BrandService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, brandID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#BrandGetParams">BrandGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Brand">Brand</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/brands">client.Brands.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#BrandService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#BrandListParams">BrandListParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Brand">Brand</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/brands/search">client.Brands.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#BrandService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#BrandSearchParams">BrandSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchBrandsResponse">SearchBrandsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Categories

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Category">Category</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CategoryAttribute">CategoryAttribute</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CategoryRef">CategoryRef</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CategorySummary">CategorySummary</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PaginatedListCategoriesResponse">PaginatedListCategoriesResponse</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchCategoriesResponse">SearchCategoriesResponse</a>

Methods:

- <code title="get /v1/categories/{slug}">client.Categories.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CategoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, slug <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Category">Category</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/categories">client.Categories.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CategoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CategoryListParams">CategoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination#CategoryPage">CategoryPage</a>[<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CategorySummary">CategorySummary</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/categories/search">client.Categories.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CategoryService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CategorySearchParams">CategorySearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchCategoriesResponse">SearchCategoriesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Websites

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Website">Website</a>

Methods:

- <code title="get /v0/websites">client.Websites.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#WebsiteService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#WebsiteGetParams">WebsiteGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Website">Website</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PriceTracking

Params Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#StartTrackingRequestParam">StartTrackingRequestParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#StopTrackingRequestParam">StopTrackingRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PaginatedSubscriptionsResponse">PaginatedSubscriptionsResponse</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceHistory">PriceHistory</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceHistoryPoint">PriceHistoryPoint</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceStatistics">PriceStatistics</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Subscription">Subscription</a>

Methods:

- <code title="get /v0/price-tracking/subscriptions">client.PriceTracking.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceTrackingService.ListSubscriptions">ListSubscriptions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceTrackingListSubscriptionsParams">PriceTrackingListSubscriptionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go/packages/pagination#CursorPage">CursorPage</a>[<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Subscription">Subscription</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v0/price-tracking/history/{canonical_product_id}">client.PriceTracking.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceTrackingService.GetHistory">GetHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, canonicalProductID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceTrackingGetHistoryParams">PriceTrackingGetHistoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceHistory">PriceHistory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/price-tracking/start">client.PriceTracking.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceTrackingService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceTrackingStartParams">PriceTrackingStartParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v0/price-tracking/stop">client.PriceTracking.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceTrackingService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#PriceTrackingStopParams">PriceTrackingStopParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#Subscription">Subscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Conversations

Params Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ConversationContextParam">ConversationContextParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CreateTurnRequestParam">CreateTurnRequestParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ImagePartParam">ImagePartParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#TextPartParam">TextPartParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#UserMessageParam">UserMessageParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#AssistantMessage">AssistantMessage</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CatalogDisplayPayload">CatalogDisplayPayload</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#CatalogToolError">CatalogToolError</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ConversationContext">ConversationContext</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ConversationDetail">ConversationDetail</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ImagePart">ImagePart</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ProductIDsInput">ProductIDsInput</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchProductsInput">SearchProductsInput</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#TextPart">TextPart</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ToolPart">ToolPart</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#TurnResult">TurnResult</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#TurnUsage">TurnUsage</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#UserMessage">UserMessage</a>

Methods:

- <code title="post /v1/conversations">client.Conversations.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ConversationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ConversationNewParams">ConversationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#TurnResult">TurnResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/conversations/{conversation_id}">client.Conversations.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ConversationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ConversationGetParams">ConversationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#ConversationDetail">ConversationDetail</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Search

Params Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchConfigParam">SearchConfigParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchFilterPriceParam">SearchFilterPriceParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchFiltersParam">SearchFiltersParam</a>
- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchRequestParam">SearchRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchResponse">SearchResponse</a>

Methods:

- <code title="post /v1/search">client.Search.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchService.Perform">Perform</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchPerformParams">SearchPerformParams</a>) (\*<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go">channel3go</a>.<a href="https://pkg.go.dev/github.com/channel3-ai/sdk-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Enrich
