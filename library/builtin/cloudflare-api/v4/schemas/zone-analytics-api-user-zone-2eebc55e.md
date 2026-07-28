---
title: zone-analytics-api_user_zone
page_id: schema-zone-analytics-api-user-zone-2eebc55e
path: schemas
description: Analytics data for a single zone owned by the user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_user_zone

Analytics data for a single zone owned by the user.

```yaml
{"description": "Analytics data for a single zone owned by the user.", "type": "object", "properties": {"timeseries": {"description": "Time deltas containing analytics data for each bucket. The number of\nbuckets (resolution) is determined by the time range between since and until.", "type": "array", "items": {"$ref": "#/components/schemas/zone-analytics-api_user_period"}}, "totals": {"$ref": "#/components/schemas/zone-analytics-api_user_period"}, "zone_id": {"description": "The zone identifier.", "type": "string", "example": "023e105f4ecef8ad9ca31a8372d0c353", "maxLength": 32, "readOnly": true}}}
```
