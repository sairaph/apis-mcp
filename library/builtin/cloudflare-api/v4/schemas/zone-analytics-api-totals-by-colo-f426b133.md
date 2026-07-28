---
title: zone-analytics-api_totals_by_colo
page_id: schema-zone-analytics-api-totals-by-colo-f426b133
path: schemas
description: Breakdown of totals by data type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_totals_by_colo

Breakdown of totals by data type.

```yaml
{"description": "Breakdown of totals by data type.", "type": "object", "properties": {"bandwidth": {"$ref": "#/components/schemas/zone-analytics-api_bandwidth_by_colo"}, "requests": {"$ref": "#/components/schemas/zone-analytics-api_requests_by_colo"}, "since": {"$ref": "#/components/schemas/zone-analytics-api_since"}, "threats": {"$ref": "#/components/schemas/zone-analytics-api_threats"}, "until": {"$ref": "#/components/schemas/zone-analytics-api_until"}}}
```
