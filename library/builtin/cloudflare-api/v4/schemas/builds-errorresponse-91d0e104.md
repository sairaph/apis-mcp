---
title: builds_ErrorResponse
page_id: schema-builds-errorresponse-91d0e104
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_ErrorResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "example": 12000}, "message": {"type": "string", "example": "Not found"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "example": false}}, "required": ["success", "errors", "messages", "result"]}
```
