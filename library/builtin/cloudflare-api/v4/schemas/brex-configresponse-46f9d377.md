---
title: brex_ConfigResponse
page_id: schema-brex-configresponse-46f9d377
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# brex_ConfigResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/brex_ApiMessage"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/brex_ApiMessage"}}, "result": {"description": "Browser extension configuration payload.", "type": "string", "example": "Hello World!"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["result", "success", "errors", "messages"]}
```
