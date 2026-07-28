---
title: lex_UpdateDatasetRequest
page_id: schema-lex-updatedatasetrequest-c0a5544b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lex_UpdateDatasetRequest

```yaml
{"type": "object", "properties": {"enabled": {"description": "Whether to enable or disable log ingest for this dataset.", "type": "boolean", "x-auditable": true}, "fields": {"description": "Controls which fields the API ingests after the update. Defaults\nto all available fields when absent.\n", "type": "array", "items": {"$ref": "#/components/schemas/lex_LogField"}}}, "required": ["enabled"]}
```
