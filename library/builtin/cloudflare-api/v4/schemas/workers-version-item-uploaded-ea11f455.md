---
title: workers_version-item-uploaded
page_id: schema-workers-version-item-uploaded-ea11f455
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_version-item-uploaded

```yaml
{"allOf": [{"$ref": "#/components/schemas/workers_version-item-full"}, {"properties": {"exports_reconciliation": {"description": "Summary of the declarative exports reconciliation that\nran on this upload. Populated only when the uploaded\nmetadata included an `exports` block. Durable Object\nentries drive reconciliation; `type: worker` entries do\nnot contribute to this summary.\n", "allOf": [{"$ref": "#/components/schemas/workers_exports_reconciliation_result"}]}, "startup_time_ms": {"description": "Time in milliseconds spent on [Worker startup](https://developers.cloudflare.com/workers/platform/limits/#worker-startup-time).", "type": "integer", "example": 10}}, "required": ["resources"], "type": "object"}]}
```
