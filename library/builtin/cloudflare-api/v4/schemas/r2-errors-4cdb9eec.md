---
title: r2_errors
page_id: schema-r2-errors-4cdb9eec
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_errors

```yaml
{"type": "array", "items": {"properties": {"code": {"type": "integer", "minimum": 1000}, "message": {"type": "string"}}, "required": ["code", "message"], "type": "object", "uniqueItems": true}}
```
