---
title: workers_binding_kind_queue
page_id: schema-workers-binding-kind-queue-0daf55e1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_queue

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "queue_name": {"description": "Name of the Queue to bind to.", "type": "string", "example": "my-queue"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["queue"], "x-auditable": true}}, "required": ["name", "type", "queue_name"]}
```
