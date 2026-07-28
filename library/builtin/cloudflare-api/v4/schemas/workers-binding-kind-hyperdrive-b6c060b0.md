---
title: workers_binding_kind_hyperdrive
page_id: schema-workers-binding-kind-hyperdrive-b6c060b0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_hyperdrive

```yaml
{"type": "object", "properties": {"id": {"description": "Identifier of the Hyperdrive connection to bind to.", "type": "string", "example": "57b7076f58be42419276f058a8968187", "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["hyperdrive"], "x-auditable": true}}, "required": ["name", "type", "id"]}
```
