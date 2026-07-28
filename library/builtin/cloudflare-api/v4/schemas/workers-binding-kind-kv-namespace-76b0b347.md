---
title: workers_binding_kind_kv_namespace
page_id: schema-workers-binding-kind-kv-namespace-76b0b347
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_kv_namespace

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "namespace_id": {"$ref": "#/components/schemas/workers_namespace_identifier"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["kv_namespace"], "x-auditable": true}}, "required": ["name", "type", "namespace_id"]}
```
