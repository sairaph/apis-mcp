---
title: art_ErrorResponse
page_id: schema-art-errorresponse-04a9d22b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_ErrorResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/art_APIError"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/art_APIMessage"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "example": false}}, "required": ["success", "errors", "messages"]}
```
