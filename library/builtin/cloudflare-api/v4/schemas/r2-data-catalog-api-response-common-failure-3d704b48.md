---
title: r2-data-catalog_api-response-common-failure
page_id: schema-r2-data-catalog-api-response-common-failure-3d704b48
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_api-response-common-failure

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer"}, "message": {"type": "string"}}, "type": "object"}, "minItems": 1}, "messages": {"type": "array", "items": {"type": "object"}}, "success": {"type": "boolean", "enum": [false]}}}
```
