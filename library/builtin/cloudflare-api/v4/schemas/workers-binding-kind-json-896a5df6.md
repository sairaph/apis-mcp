---
title: workers_binding_kind_json
page_id: schema-workers-binding-kind-json-896a5df6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_json

```yaml
{"type": "object", "properties": {"json": {"description": "JSON data to use.", "type": "object", "x-stainless-any": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["json"], "x-auditable": true}}, "required": ["name", "type", "json"]}
```
