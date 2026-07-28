---
title: lex_CreateDatasetRequest
page_id: schema-lex-createdatasetrequest-9b8725b6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lex_CreateDatasetRequest

```yaml
{"type": "object", "properties": {"dataset": {"description": "Dataset type name to create (e.g. `http_requests`).", "type": "string", "x-auditable": true}, "fields": {"description": "Controls which fields the API ingests. Defaults to all available\nfields when absent.\n", "type": "array", "items": {"$ref": "#/components/schemas/lex_LogField"}}}, "required": ["dataset"]}
```
