---
title: dns-analytics_report_bytime
page_id: schema-dns-analytics-report-bytime-02892bed
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-analytics_report_bytime

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-analytics_result"}, {"properties": {"data": {"type": "array", "items": {"properties": {"metrics": {"description": "Array with one item per requested metric. Each item is an array of values, broken down by time interval.", "type": "array", "items": {"description": "Nominal metric values, broken down by time interval.", "items": {"description": "Nominal metric value.", "type": "number"}, "type": "array"}}}, "required": ["metrics"], "type": "object"}}, "query": {"properties": {"time_delta": {"$ref": "#/components/schemas/dns-analytics_time_delta"}}, "required": ["time_delta"], "type": "object"}, "time_intervals": {"description": "Array of time intervals in the response data. Each interval is represented as an array containing two values: the start time, and the end time.\n", "type": "array", "items": {"description": "Array with exactly two items, representing the start and end time (respectively) of this time interval.", "items": {"description": "Time value.", "example": "2023-11-11T12:00:00Z", "format": "date-time", "type": "string"}, "type": "array"}}}, "required": ["time_intervals", "query", "data"], "type": "object"}]}
```
