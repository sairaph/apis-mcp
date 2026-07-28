---
title: spectrum-analytics_query-response-aggregate
page_id: schema-spectrum-analytics-query-response-aggregate-31070126
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-analytics_query-response-aggregate

```yaml
{"allOf": [{"$ref": "#/components/schemas/spectrum-analytics_api-response-single"}, {"properties": {"result": {"type": "array", "items": {"properties": {"appID": {"allOf": [{"description": "Application identifier."}, {"$ref": "#/components/schemas/spectrum-analytics_identifier"}]}, "bytesEgress": {"description": "Number of bytes sent.", "type": "number"}, "bytesIngress": {"description": "Number of bytes received.", "type": "number"}, "connections": {"description": "Number of connections.", "type": "number"}, "durationAvg": {"description": "Average duration of connections.", "type": "number"}}, "required": ["appID", "bytesIngress", "bytesEgress", "connections", "durationAvg"], "type": "object"}}}}]}
```
