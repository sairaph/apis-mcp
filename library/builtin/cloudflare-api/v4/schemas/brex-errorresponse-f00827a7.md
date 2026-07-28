---
title: brex_ErrorResponse
page_id: schema-brex-errorresponse-f00827a7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# brex_ErrorResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/brex_ApiMessage"}, "minItems": 1}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/brex_ApiMessage"}}, "result": {"nullable": true}, "success": {"type": "boolean", "enum": [false]}}, "required": ["result", "success", "errors", "messages"]}
```
