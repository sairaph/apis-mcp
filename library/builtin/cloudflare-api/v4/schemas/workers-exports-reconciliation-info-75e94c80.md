---
title: workers_exports_reconciliation_info
page_id: schema-workers-exports-reconciliation-info-75e94c80
path: schemas
description: |-
    A non-blocking reconciliation info entry. Emitted for stale
    tombstones (a no-op on this deploy) and for tombstones applied
    with the source class still in code (the supported zero-downtime
    rollout pattern).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_reconciliation_info

A non-blocking reconciliation info entry. Emitted for stale
tombstones (a no-op on this deploy) and for tombstones applied
with the source class still in code (the supported zero-downtime
rollout pattern).

```yaml
{"description": "A non-blocking reconciliation info entry. Emitted for stale\ntombstones (a no-op on this deploy) and for tombstones applied\nwith the source class still in code (the supported zero-downtime\nrollout pattern).\n", "type": "object", "properties": {"class": {"description": "The class name the info entry is about.", "type": "string"}, "message": {"description": "Human-readable explanation.", "type": "string"}, "namespace_id": {"description": "The provisioned namespace the entry relates to, when applicable.", "type": "string", "format": "uuid"}, "referencing_scripts": {"description": "Other Workers in the account that still bind to the affected\nclass. Advisory: while non-empty the tombstone is not yet safe\nto remove — redeploy these Workers with bindings re-pointed\nfirst.\n", "type": "array", "items": {"type": "string"}}, "scenario": {"$ref": "#/components/schemas/workers_exports_reconciliation_scenario"}}, "required": ["class", "scenario", "message"]}
```
