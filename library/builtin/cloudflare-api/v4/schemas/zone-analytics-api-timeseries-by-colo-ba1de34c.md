---
title: zone-analytics-api_timeseries_by_colo
page_id: schema-zone-analytics-api-timeseries-by-colo-ba1de34c
path: schemas
description: Time deltas containing metadata about each bucket of time. The number of buckets (resolution) is determined by the amount of time between the since and until parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_timeseries_by_colo

Time deltas containing metadata about each bucket of time. The number of buckets (resolution) is determined by the amount of time between the since and until parameters.

```yaml
{"description": "Time deltas containing metadata about each bucket of time. The number of buckets (resolution) is determined by the amount of time between the since and until parameters.", "type": "array", "items": {"properties": {"bandwidth": {"$ref": "#/components/schemas/zone-analytics-api_bandwidth_by_colo"}, "requests": {"$ref": "#/components/schemas/zone-analytics-api_requests_by_colo"}, "since": {"$ref": "#/components/schemas/zone-analytics-api_since"}, "threats": {"$ref": "#/components/schemas/zone-analytics-api_threats"}, "until": {"$ref": "#/components/schemas/zone-analytics-api_until"}}, "type": "object"}}
```
