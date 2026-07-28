---
title: workers_exports_reconciliation_transfer
page_id: schema-workers-exports-reconciliation-transfer-9842cc57
path: schemas
description: A single committed `transferred` tombstone (phase-2 commit).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_reconciliation_transfer

A single committed `transferred` tombstone (phase-2 commit).

```yaml
{"description": "A single committed `transferred` tombstone (phase-2 commit).", "type": "object", "properties": {"class": {"description": "The source class name that was transferred.", "type": "string"}, "phase": {"description": "The transfer phase. Currently always `committed`.", "type": "string", "example": "committed", "enum": ["committed"]}, "to": {"description": "The destination script that now owns the namespace.", "type": "string"}}, "required": ["class", "to", "phase"]}
```
