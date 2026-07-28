---
title: r2-slurper_api-v4-error
page_id: schema-r2-slurper-api-v4-error-2bdb8034
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-slurper_api-v4-error

```yaml
{"type": "array", "items": {"properties": {"code": {"type": "integer", "minimum": 1000, "x-auditable": true}, "message": {"type": "string", "x-auditable": true}}, "required": ["code", "message"], "type": "object", "uniqueItems": true}, "example": [{"code": 7003, "message": "No route for the URI"}], "minLength": 1}
```
