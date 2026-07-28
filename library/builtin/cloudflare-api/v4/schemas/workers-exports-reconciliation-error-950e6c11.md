---
title: workers_exports_reconciliation_error
page_id: schema-workers-exports-reconciliation-error-950e6c11
path: schemas
description: |-
    A single blocking reconciliation error. Returned in the v4 error
    envelope under `meta.details` on the exports reconciliation
    failure error (code 100402). All per-class failures are reported
    together, sorted by class name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_reconciliation_error

A single blocking reconciliation error. Returned in the v4 error
envelope under `meta.details` on the exports reconciliation
failure error (code 100402). All per-class failures are reported
together, sorted by class name.

```yaml
{"description": "A single blocking reconciliation error. Returned in the v4 error\nenvelope under `meta.details` on the exports reconciliation\nfailure error (code 100402). All per-class failures are reported\ntogether, sorted by class name.\n", "type": "object", "properties": {"class": {"description": "The class name the error is about.", "type": "string"}, "message": {"description": "Human-readable explanation.", "type": "string"}, "namespace_id": {"description": "The provisioned namespace the error relates to, when applicable.", "type": "string", "format": "uuid"}, "referencing_scripts": {"description": "Other Workers in the account whose bindings block the\noperation (e.g. a `deleted` tombstone blocked by external\nbindings).\n", "type": "array", "items": {"type": "string"}}, "scenario": {"$ref": "#/components/schemas/workers_exports_reconciliation_scenario"}, "suggestion": {"description": "An actionable hint for resolving the error, when available.", "type": "string"}}, "required": ["class", "scenario", "message"]}
```
