---
title: workers_binding_kind_dispatch_namespace
page_id: schema-workers-binding-kind-dispatch-namespace-0ee882ae
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_dispatch_namespace

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "namespace": {"description": "The name of the dispatch namespace.", "type": "string", "example": "my-namespace", "x-auditable": true}, "outbound": {"description": "Outbound worker.", "type": "object", "properties": {"params": {"description": "Pass information from the Dispatch Worker to the Outbound Worker through the parameters.", "type": "array", "items": {"properties": {"name": {"description": "Name of the parameter.", "type": "string", "example": "customer_name", "x-auditable": true}}, "required": ["name"], "type": "object"}}, "worker": {"description": "Outbound worker.", "type": "object", "properties": {"entrypoint": {"description": "Entrypoint to invoke on the outbound worker.", "type": "string", "x-auditable": true}, "environment": {"description": "Environment of the outbound worker.", "type": "string", "x-auditable": true}, "service": {"description": "Name of the outbound worker.", "type": "string", "x-auditable": true}}}}}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["dispatch_namespace"], "x-auditable": true}}, "required": ["name", "type", "namespace"]}
```
