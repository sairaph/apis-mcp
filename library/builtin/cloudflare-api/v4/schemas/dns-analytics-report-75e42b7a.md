---
title: dns-analytics_report
page_id: schema-dns-analytics-report-75e42b7a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-analytics_report

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-analytics_result"}, {"properties": {"data": {"type": "array", "items": {"properties": {"metrics": {"description": "Array with one item per requested metric. Each item is a single value.", "type": "array", "items": {"description": "Nominal metric value.", "type": "number"}}}, "required": ["metrics"], "type": "object"}}}, "required": ["data"], "type": "object"}]}
```
