---
title: zone-analytics-api_user_period
page_id: schema-zone-analytics-api-user-period-8edf4ecc
path: schemas
description: Analytics data for a single time period bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_user_period

Analytics data for a single time period bucket.

```yaml
{"description": "Analytics data for a single time period bucket.", "type": "object", "properties": {"bandwidth": {"$ref": "#/components/schemas/zone-analytics-api_user_period_bandwidth"}, "requests": {"$ref": "#/components/schemas/zone-analytics-api_user_period_requests"}, "since": {"$ref": "#/components/schemas/zone-analytics-api_since"}, "until": {"$ref": "#/components/schemas/zone-analytics-api_until"}}}
```
