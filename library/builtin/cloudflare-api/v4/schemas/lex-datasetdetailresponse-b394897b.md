---
title: lex_DatasetDetailResponse
page_id: schema-lex-datasetdetailresponse-b394897b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lex_DatasetDetailResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/lex_V4Error-2"}}, "messages": {"type": "array", "items": {"type": "string", "x-auditable": true}}, "result": {"$ref": "#/components/schemas/lex_DatasetResponse"}, "success": {"type": "boolean", "x-auditable": true}}, "required": ["success", "errors", "messages"]}
```
