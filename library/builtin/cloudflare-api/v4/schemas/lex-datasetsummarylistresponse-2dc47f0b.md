---
title: lex_DatasetSummaryListResponse
page_id: schema-lex-datasetsummarylistresponse-2dc47f0b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lex_DatasetSummaryListResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/lex_V4Error-2"}}, "messages": {"type": "array", "items": {"type": "string"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/lex_DatasetSummary"}, "nullable": true}, "success": {"type": "boolean"}}, "required": ["success", "errors", "messages"]}
```
