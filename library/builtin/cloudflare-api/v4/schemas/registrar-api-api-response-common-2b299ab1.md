---
title: registrar-api_api-response-common
page_id: schema-registrar-api-api-response-common-2b299ab1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/registrar-api_messages"}, "messages": {"$ref": "#/components/schemas/registrar-api_messages"}, "result": {"anyOf": [{"type": "object"}, {"items": {"type": "object"}, "type": "array"}, {"type": "string"}]}, "success": {"description": "Whether the API call was successful", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
