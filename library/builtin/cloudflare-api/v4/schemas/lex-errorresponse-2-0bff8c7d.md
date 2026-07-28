---
title: lex_ErrorResponse-2
page_id: schema-lex-errorresponse-2-0bff8c7d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lex_ErrorResponse-2

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/lex_V4Error-2"}}, "messages": {"type": "array", "items": {"type": "string", "x-auditable": true}}, "result": {"type": "object", "additionalProperties": true, "nullable": true, "x-auditable": true}, "success": {"type": "boolean", "enum": [false], "x-auditable": true}}, "required": ["result", "success", "errors", "messages"]}
```
