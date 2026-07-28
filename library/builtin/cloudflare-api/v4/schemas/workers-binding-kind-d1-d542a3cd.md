---
title: workers_binding_kind_d1
page_id: schema-workers-binding-kind-d1-d542a3cd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_d1

```yaml
{"type": "object", "properties": {"database_id": {"description": "Identifier of the D1 database to bind to.", "type": "string", "example": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", "x-auditable": true}, "id": {"description": "Identifier of the D1 database to bind to.", "type": "string", "example": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", "deprecated": true, "x-auditable": true, "x-stainless-deprecation-message": "This property has been renamed to `database_id`."}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["d1"], "x-auditable": true}}, "required": ["name", "type", "database_id"]}
```
