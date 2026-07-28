---
title: lex_QueryResponse
page_id: schema-lex-queryresponse-523f2d0f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lex_QueryResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/lex_V4Error"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/lex_QueryRow"}, "nullable": true}, "success": {"type": "boolean"}}, "required": ["success", "errors", "messages"]}
```
