---
title: zone-analytics-api_datacenters
page_id: schema-zone-analytics-api-datacenters-341f595c
path: schemas
description: A breakdown of all dashboard analytics data by co-locations. This is limited to Enterprise zones only.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_datacenters

A breakdown of all dashboard analytics data by co-locations. This is limited to Enterprise zones only.

```yaml
{"description": "A breakdown of all dashboard analytics data by co-locations. This is limited to Enterprise zones only.", "type": "array", "items": {"properties": {"colo_id": {"description": "The airport code identifer for the co-location.", "type": "string", "example": "SFO"}, "timeseries": {"$ref": "#/components/schemas/zone-analytics-api_timeseries_by_colo"}, "totals": {"$ref": "#/components/schemas/zone-analytics-api_totals_by_colo"}}, "type": "object"}, "title": "Analytics data by datacenter"}
```
