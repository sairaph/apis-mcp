---
title: dns-analytics_result
page_id: schema-dns-analytics-result-f9afb74b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-analytics_result

```yaml
{"type": "object", "properties": {"data": {"$ref": "#/components/schemas/dns-analytics_data"}, "data_lag": {"description": "Number of seconds between current time and last processed event, in another words how many seconds of data could be missing.", "type": "number", "example": 60, "minimum": 0}, "max": {"description": "Maximum results for each metric (object mapping metric names to values). Currently always an empty object.", "type": "object"}, "min": {"description": "Minimum results for each metric (object mapping metric names to values). Currently always an empty object.", "type": "object"}, "query": {"$ref": "#/components/schemas/dns-analytics_query"}, "rows": {"description": "Total number of rows in the result.", "type": "number", "example": 100, "minimum": 0}, "totals": {"description": "Total results for metrics across all data (object mapping metric names to values).", "type": "object"}}, "required": ["rows", "totals", "min", "max", "data_lag", "query", "data"]}
```
