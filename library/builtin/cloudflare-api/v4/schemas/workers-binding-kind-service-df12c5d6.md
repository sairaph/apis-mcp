---
title: workers_binding_kind_service
page_id: schema-workers-binding-kind-service-df12c5d6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_service

```yaml
{"type": "object", "properties": {"entrypoint": {"description": "Entrypoint to invoke on the target Worker.", "type": "string", "example": "MyHandler", "x-auditable": true}, "environment": {"description": "Optional environment if the Worker utilizes one.", "type": "string", "example": "production", "default": "production", "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "service": {"description": "Name of Worker to bind to.", "type": "string", "example": "my-worker", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["service"], "x-auditable": true}}, "required": ["name", "type", "service"]}
```
