---
title: d1_raw-result-response
page_id: schema-d1-raw-result-response-fa3b09aa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_raw-result-response

```yaml
{"type": "object", "properties": {"meta": {"$ref": "#/components/schemas/d1_query-meta"}, "results": {"type": "object", "properties": {"columns": {"type": "array", "items": {"type": "string"}}, "rows": {"type": "array", "items": {"items": {"anyOf": [{"type": "number"}, {"type": "string"}, {"type": "object"}]}, "type": "array"}}}}, "success": {"type": "boolean"}}}
```
