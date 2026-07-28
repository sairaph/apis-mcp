---
title: images_api-response-common
page_id: schema-images-api-response-common-df96a00b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/images_messages"}, "messages": {"$ref": "#/components/schemas/images_messages"}, "result": {"anyOf": [{"type": "object"}, {"items": {}, "type": "array"}, {"type": "string"}]}, "success": {"description": "Whether the API call was successful", "type": "boolean", "example": true, "enum": [true], "x-auditable": true}}, "required": ["success", "errors", "messages", "result"]}
```
