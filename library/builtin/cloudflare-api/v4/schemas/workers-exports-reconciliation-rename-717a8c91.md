---
title: workers_exports_reconciliation_rename
page_id: schema-workers-exports-reconciliation-rename-717a8c91
path: schemas
description: A single applied `renamed` tombstone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_reconciliation_rename

A single applied `renamed` tombstone.

```yaml
{"description": "A single applied `renamed` tombstone.", "type": "object", "properties": {"from": {"description": "The original (source) class name.", "type": "string"}, "to": {"description": "The new class name (`renamed_to`).", "type": "string"}}, "required": ["from", "to"]}
```
