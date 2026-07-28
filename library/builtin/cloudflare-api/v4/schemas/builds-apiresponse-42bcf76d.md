---
title: builds_APIResponse
page_id: schema-builds-apiresponse-42bcf76d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_APIResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "example": 12000}, "message": {"type": "string", "example": "Not found"}}, "type": "object"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "object", "nullable": true}, "result_info": {"$ref": "#/components/schemas/builds_PaginationInfo"}, "success": {"type": "boolean", "example": true}}, "required": ["success", "errors", "messages", "result"]}
```
