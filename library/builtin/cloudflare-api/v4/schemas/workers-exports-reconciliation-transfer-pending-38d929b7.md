---
title: workers_exports_reconciliation_transfer_pending
page_id: schema-workers-exports-reconciliation-transfer-pending-38d929b7
path: schemas
description: |-
    A single phase-1 transfer hint recorded on the target side (a live
    `expecting-transfer` entry).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_reconciliation_transfer_pending

A single phase-1 transfer hint recorded on the target side (a live
`expecting-transfer` entry).

```yaml
{"description": "A single phase-1 transfer hint recorded on the target side (a live\n`expecting-transfer` entry).\n", "type": "object", "properties": {"class": {"description": "The target-side class name awaiting transfer.", "type": "string"}, "from": {"description": "The source script the namespace will be transferred from.", "type": "string"}}, "required": ["class", "from"]}
```
