---
title: lex_DatasetSummary
page_id: schema-lex-datasetsummary-42cc4623
path: schemas
description: |-
    A Log Explorer dataset summary. List endpoints return this type and omit
    field configuration; use the single-dataset endpoint to retrieve it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# lex_DatasetSummary

A Log Explorer dataset summary. List endpoints return this type and omit
field configuration; use the single-dataset endpoint to retrieve it.

```yaml
{"description": "A Log Explorer dataset summary. List endpoints return this type and omit\nfield configuration; use the single-dataset endpoint to retrieve it.\n", "type": "object", "properties": {"created_at": {"description": "RFC3339 timestamp recording when the API created this dataset.", "type": "string", "format": "date-time", "x-auditable": true}, "dataset": {"description": "Dataset type name (e.g. `http_requests`).", "type": "string", "x-auditable": true}, "dataset_id": {"description": "Unique dataset ID.", "type": "string", "pattern": "^[a-f0-9]{32}$", "x-auditable": true}, "enabled": {"description": "Whether log ingest is currently active for this dataset.", "type": "boolean", "x-auditable": true}, "object_id": {"description": "Public ID of the account or zone that owns this dataset.", "type": "string", "x-auditable": true}, "object_type": {"description": "Whether this dataset belongs to an account or a zone.", "type": "string", "enum": ["account", "zone"], "x-auditable": true}, "updated_at": {"description": "RFC3339 timestamp recording when the API last updated this dataset.", "type": "string", "format": "date-time", "x-auditable": true}}, "required": ["dataset", "object_type", "object_id", "created_at", "updated_at", "dataset_id", "enabled"]}
```
