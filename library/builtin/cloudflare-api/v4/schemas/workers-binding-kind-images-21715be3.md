---
title: workers_binding_kind_images
page_id: schema-workers-binding-kind-images-21715be3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_images

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["images"], "x-auditable": true}}, "required": ["name", "type"]}
```
