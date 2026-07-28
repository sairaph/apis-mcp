---
title: magic-visibility-pcaps_api-response-common
page_id: schema-magic-visibility-pcaps-api-response-common-ee39ad0b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-visibility-pcaps_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/magic-visibility-pcaps_messages"}, "messages": {"$ref": "#/components/schemas/magic-visibility-pcaps_messages"}, "result": {"anyOf": [{"type": "object"}, {"items": {}, "type": "array"}, {"type": "string"}]}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
