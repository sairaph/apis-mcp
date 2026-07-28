---
title: magic_apps-response-object
page_id: schema-magic-apps-response-object-0a3313d3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_apps-response-object

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/magic_messages"}, "messages": {"$ref": "#/components/schemas/magic_messages"}, "result": {"type": "object", "nullable": true}, "success": {"description": "Whether the API call was successful", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
