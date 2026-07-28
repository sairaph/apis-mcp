---
title: workers_exports_reconciliation_result
page_id: schema-workers-exports-reconciliation-result-471a310c
path: schemas
description: |-
    Summary of what the declarative exports reconciliation did on this
    upload. Present only on uploads that included an `exports` block.
    Reconciliation acts on Durable Object entries; `type: worker`
    entries in the same map do not participate. Every array is always
    present (possibly empty) so clients can iterate unconditionally.
    Clients should surface `warnings` prominently, render `info` at
    lower visibility, and use `removable_entries` as the "safe to
    delete from your config" hint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_reconciliation_result

Summary of what the declarative exports reconciliation did on this
upload. Present only on uploads that included an `exports` block.
Reconciliation acts on Durable Object entries; `type: worker`
entries in the same map do not participate. Every array is always
present (possibly empty) so clients can iterate unconditionally.
Clients should surface `warnings` prominently, render `info` at
lower visibility, and use `removable_entries` as the "safe to
delete from your config" hint.

```yaml
{"description": "Summary of what the declarative exports reconciliation did on this\nupload. Present only on uploads that included an `exports` block.\nReconciliation acts on Durable Object entries; `type: worker`\nentries in the same map do not participate. Every array is always\npresent (possibly empty) so clients can iterate unconditionally.\nClients should surface `warnings` prominently, render `info` at\nlower visibility, and use `removable_entries` as the \"safe to\ndelete from your config\" hint.\n", "type": "object", "properties": {"created": {"description": "Class names for which a new namespace was provisioned.", "type": "array", "items": {"type": "string"}}, "deleted": {"description": "Class names whose namespace was deleted by a `deleted` tombstone.", "type": "array", "items": {"type": "string"}}, "info": {"description": "Non-blocking info entries (stale tombstones, tombstone applied\nwith class still in code). See `exports_reconciliation_info`.\n", "type": "array", "items": {"$ref": "#/components/schemas/workers_exports_reconciliation_info"}}, "removable_entries": {"description": "Source class names whose tombstone entry is now stale and safe\nto delete from `exports` (no remaining referencing scripts).\n", "type": "array", "items": {"type": "string"}}, "renamed": {"description": "Applied `renamed` tombstones.", "type": "array", "items": {"$ref": "#/components/schemas/workers_exports_reconciliation_rename"}}, "transfer_pending": {"description": "Phase-1 transfer hints recorded on the target side.", "type": "array", "items": {"$ref": "#/components/schemas/workers_exports_reconciliation_transfer_pending"}}, "transferred": {"description": "Committed `transferred` tombstones (phase-2).", "type": "array", "items": {"$ref": "#/components/schemas/workers_exports_reconciliation_transfer"}}, "updated": {"description": "Class names whose provisioned namespace was mutated in place.", "type": "array", "items": {"type": "string"}}, "warnings": {"description": "Non-blocking warnings. See `exports_reconciliation_warning`.", "type": "array", "items": {"$ref": "#/components/schemas/workers_exports_reconciliation_warning"}}}, "readOnly": true, "required": ["created", "updated", "deleted", "renamed", "transferred", "transfer_pending", "warnings", "info", "removable_entries"]}
```
