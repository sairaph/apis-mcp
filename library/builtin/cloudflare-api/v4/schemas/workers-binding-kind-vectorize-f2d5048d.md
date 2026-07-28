---
title: workers_binding_kind_vectorize
page_id: schema-workers-binding-kind-vectorize-f2d5048d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_vectorize

```yaml
{"type": "object", "properties": {"index_name": {"description": "Name of the Vectorize index to bind to.", "type": "string", "example": "my-index-name", "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["vectorize"], "x-auditable": true}}, "required": ["name", "type", "index_name"]}
```
