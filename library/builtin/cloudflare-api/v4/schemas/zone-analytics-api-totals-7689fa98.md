---
title: zone-analytics-api_totals
page_id: schema-zone-analytics-api-totals-7689fa98
path: schemas
description: Breakdown of totals by data type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_totals

Breakdown of totals by data type.

```yaml
{"description": "Breakdown of totals by data type.", "type": "object", "properties": {"bandwidth": {"$ref": "#/components/schemas/zone-analytics-api_bandwidth"}, "pageviews": {"$ref": "#/components/schemas/zone-analytics-api_pageviews"}, "requests": {"$ref": "#/components/schemas/zone-analytics-api_requests"}, "since": {"$ref": "#/components/schemas/zone-analytics-api_since"}, "threats": {"$ref": "#/components/schemas/zone-analytics-api_threats"}, "uniques": {"$ref": "#/components/schemas/zone-analytics-api_uniques"}, "until": {"$ref": "#/components/schemas/zone-analytics-api_until"}}}
```
