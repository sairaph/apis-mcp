---
title: lists_api-response-common
page_id: schema-lists-api-response-common-ac6ce22d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lists_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/lists_messages"}, "messages": {"$ref": "#/components/schemas/lists_messages"}, "result": {"anyOf": [{"type": "object"}, {"items": {"type": "object"}, "type": "array"}]}, "success": {"description": "Defines whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
