# Changelog

## 0.5.0 (2026-08-18)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/channel3-ai/sdk-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** add cleaned_url to enrich and product image responses ([b529715](https://github.com/channel3-ai/sdk-go/commit/b5297155a152c52211c00463956edb2544bb8830))
* **api:** add clicks and transactions list endpoints ([6667d99](https://github.com/channel3-ai/sdk-go/commit/6667d995a22d6b95c5859555232c093854cafd6b))
* **api:** add client_tokens resource, DiscussionPart, update availability types ([78dab8e](https://github.com/channel3-ai/sdk-go/commit/78dab8e3c4ee3d6d4b8031cee807b5ae875c2f31))
* **api:** add conversation items pagination, remove responses endpoint and streaming ([ef7e818](https://github.com/channel3-ai/sdk-go/commit/ef7e818b401b794c4d368a7bfc3df53990758159))
* **api:** add filters field to chat request ([313920e](https://github.com/channel3-ai/sdk-go/commit/313920eea173d0e5fac3608f26558026e4115713))
* **api:** add length_unit/weight_unit options, rename preferred_* fields across models ([7a87120](https://github.com/channel3-ai/sdk-go/commit/7a8712045878d9e1e9dcfaec732e7a0b245de99f))
* **api:** add mode parameter, deprecate keyword_search_only in search config ([e49c3b3](https://github.com/channel3-ai/sdk-go/commit/e49c3b3ab15fb3eda2d6e7f8d8e20181f65f0fcb))
* **api:** add monetize method to products ([7c9c44e](https://github.com/channel3-ai/sdk-go/commit/7c9c44efbdc27a20e8d3d5a7b1bd723665494383))
* **api:** add responses/conversations endpoints and conversation types ([de92a5c](https://github.com/channel3-ai/sdk-go/commit/de92a5ccb9c5d7130b884eefca272b71839e4a23))
* **api:** add sort parameter to product browse method ([83c3e2e](https://github.com/channel3-ai/sdk-go/commit/83c3e2ed6ae5aaa5ff30bc9ba75c34285bfe7b4b))
* Experimental ClickHouse scraping telemetry (+ Grafana dashboards) ([86cb012](https://github.com/channel3-ai/sdk-go/commit/86cb012edb976cb3cb160579ecea21e9247ee9b6))
* Index offer dimensions into OpenSearch and expose dimension search filters ([37a7830](https://github.com/channel3-ai/sdk-go/commit/37a7830eeb58fc619000107cd0becba14af1fdfc))
* **stlc:** configurable CI runner and private-production-repo support in workflow templates ([8bcb173](https://github.com/channel3-ai/sdk-go/commit/8bcb173ef0954eb29d6e52bcd85e154e2417a8cf))


### Bug Fixes

* **types:** correct is_error JSON field name in CatalogToolError ([00458c3](https://github.com/channel3-ai/sdk-go/commit/00458c35da62064878fe23ad9d49e9c918852b0a))
* **types:** rename AffiliateProduct to ReportingProduct ([bca0c2e](https://github.com/channel3-ai/sdk-go/commit/bca0c2e64942a0639edc0e1fbf379c3c189c9fc7))


### Chores

* **api:** remove deprecated brands.Find and enrich.EnrichURL methods ([051b0a1](https://github.com/channel3-ai/sdk-go/commit/051b0a1de89db24620270651f224d763f5de8f2f))
* **internal:** regenerate SDK with no functional changes ([d3b530e](https://github.com/channel3-ai/sdk-go/commit/d3b530e8c997674d12d3839fadb791c43abbd414))
* **internal:** regenerate SDK with no functional changes ([eb3e7c4](https://github.com/channel3-ai/sdk-go/commit/eb3e7c4949628be66136f0f18e7f6ddfce0101a5))
* **internal:** regenerate SDK with no functional changes ([199c691](https://github.com/channel3-ai/sdk-go/commit/199c6916a387335b9b1644f305fcd9cd47321a77))


### Documentation

* **api:** clarify Attributes parameter documentation in search filters ([66c38bd](https://github.com/channel3-ai/sdk-go/commit/66c38bd991f5edd24f59cf04cd33d3823b2987ca))
* **api:** clarify max_commission_rate as decimal fraction in ProductOffer ([f9616fd](https://github.com/channel3-ai/sdk-go/commit/f9616fd92d81896749a651f88e1c3aaec7e1dd27))
* **api:** clarify page_token satisfies required parameter in search ([8b536f2](https://github.com/channel3-ai/sdk-go/commit/8b536f22b16e126c4046763d2dbeb88697479ac9))
* **types:** mark is_cleaned_image as deprecated in image types ([ff0cf13](https://github.com/channel3-ai/sdk-go/commit/ff0cf1398a2fa298fa20e5610094a06fb3f930ea))

## 0.4.0 (2026-05-25)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/channel3-ai/sdk-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** add Category field, deprecate Categories in enrich/product responses ([0a919b8](https://github.com/channel3-ai/sdk-go/commit/0a919b822ad841452f59a5adde414edf6d2938d3))


### Documentation

* **types:** expand is_cleaned_image description in enrich and product models ([7e367af](https://github.com/channel3-ai/sdk-go/commit/7e367afda0694632cb351baf9ae1595e0df5fd51))

## 0.3.0 (2026-05-21)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/channel3-ai/sdk-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** add attributes filter, extracted_attributes to enrich/product responses ([0e47ba6](https://github.com/channel3-ai/sdk-go/commit/0e47ba6c7e1290dbc82e829fef4362d7dd4bd18c))
* **api:** add variants to product/enrich responses and AvailabilityStatus type ([35c1904](https://github.com/channel3-ai/sdk-go/commit/35c19045aac97f92b251774609d213676a6a1cf8))

## 0.2.0 (2026-05-14)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/channel3-ai/sdk-go/compare/v0.1.0...v0.2.0)

### Features

* **client:** optimize json encoder for internal types ([cd657ce](https://github.com/channel3-ai/sdk-go/commit/cd657ceee39301275170581546a5c1a494df2b6b))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([9ea7571](https://github.com/channel3-ai/sdk-go/commit/9ea7571d3ca916faa15cf869824c785586f96278))


### Chores

* **internal:** codegen related update ([9250531](https://github.com/channel3-ai/sdk-go/commit/925053187d5702998de2716ef86a2bbf043f0160))
* redact api-key headers in debug logs ([1868003](https://github.com/channel3-ai/sdk-go/commit/18680038957f3ac1e9d67dcb25d2dbf44a25f7e6))

## 0.1.0 (2026-05-01)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/channel3-ai/sdk-go/compare/v0.0.1...v0.1.0)

### Features

* Set up Go SDK to back our CLI ([b6d7178](https://github.com/channel3-ai/sdk-go/commit/b6d7178a57fea8aa987192027abcf45226ab6e07))


### Chores

* configure new SDK language ([a8d080b](https://github.com/channel3-ai/sdk-go/commit/a8d080b6f608cb7e526af16794298afcff279f81))
* **internal:** rename package from channel3publicsdk to channel3go ([ed2d269](https://github.com/channel3-ai/sdk-go/commit/ed2d269eed5911ce6c79cd55259acb431d5c0e93))
