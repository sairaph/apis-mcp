---
title: workers_script-settings-item
page_id: schema-workers-script-settings-item-6aae4447
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_script-settings-item

```yaml
{"type": "object", "properties": {"logpush": {"$ref": "#/components/schemas/workers_logpush"}, "observability": {"allOf": [{"$ref": "#/components/schemas/workers_observability"}, {"nullable": true, "type": "object", "x-auditable": true}]}, "tags": {"allOf": [{"$ref": "#/components/schemas/workers_tags"}, {"items": {"$ref": "#/components/schemas/workers_tag"}, "nullable": true, "type": "array", "x-auditable": true}]}, "tail_consumers": {"description": "List of Workers that will consume logs from the attached Worker.", "type": "array", "items": {"$ref": "#/components/schemas/workers_tail_consumers_script"}, "nullable": true, "x-stainless-collection-type": "set"}}}
```
