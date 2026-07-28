---
title: lex_LogField
page_id: schema-lex-logfield-99d09941
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lex_LogField

```yaml
{"type": "object", "properties": {"enabled": {"description": "Whether the API includes this field in log ingest.", "type": "boolean", "x-auditable": true}, "name": {"description": "Field name in lowercase.", "type": "string", "x-auditable": true}}, "required": ["name", "enabled"]}
```
