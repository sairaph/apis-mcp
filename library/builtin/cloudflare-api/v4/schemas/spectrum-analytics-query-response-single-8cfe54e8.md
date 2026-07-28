---
title: spectrum-analytics_query-response-single
page_id: schema-spectrum-analytics-query-response-single-8cfe54e8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-analytics_query-response-single

```yaml
{"allOf": [{"$ref": "#/components/schemas/spectrum-analytics_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"data": {"description": "List of columns returned by the analytics query.", "type": "array", "items": {"$ref": "#/components/schemas/spectrum-analytics_column"}}, "data_lag": {"description": "Number of seconds between current time and last processed event, i.e. how many seconds of data could be missing.", "type": "number", "example": 3, "minimum": 0}, "max": {"allOf": [{"description": "Maximum result for each selected metrics across all data."}, {"$ref": "#/components/schemas/spectrum-analytics_stat"}]}, "min": {"allOf": [{"description": "Minimum result for each selected metrics across all data."}, {"$ref": "#/components/schemas/spectrum-analytics_stat"}]}, "query": {"$ref": "#/components/schemas/spectrum-analytics_query"}, "rows": {"description": "Total number of rows in the result.", "type": "number", "example": 5, "minimum": 0}, "time_intervals": {"description": "List of time interval buckets: [start, end].", "type": "array", "items": {"items": {"$ref": "#/components/schemas/spectrum-analytics_timestamp"}, "type": "array"}}, "totals": {"allOf": [{"description": "Total result for each selected metrics across all data."}, {"$ref": "#/components/schemas/spectrum-analytics_stat"}]}}, "required": ["rows", "data", "data_lag", "min", "max", "totals", "query"]}}}]}
```
