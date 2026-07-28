---
title: magic_apps-response-array
page_id: schema-magic-apps-response-array-974d4923
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_apps-response-array

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/magic_messages"}, "messages": {"$ref": "#/components/schemas/magic_messages"}, "result": {"type": "array", "items": {}, "nullable": true}, "success": {"description": "Whether the API call was successful", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
