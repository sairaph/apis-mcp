---
title: zone-analytics-api_pageviews
page_id: schema-zone-analytics-api-pageviews-2e6a92dc
path: schemas
description: Breakdown of totals for pageviews.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_pageviews

Breakdown of totals for pageviews.

```yaml
{"description": "Breakdown of totals for pageviews.", "type": "object", "properties": {"all": {"description": "The total number of pageviews served within the time range.", "type": "integer"}, "search_engine": {"description": "A variable list of key/value pairs representing the search engine and number of hits.", "type": "object", "example": {"baidubot": 1345, "bingbot": 5372, "googlebot": 35272, "pingdom": 13435}}}}
```
