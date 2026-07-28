---
title: workers_binding_kind_text_blob
page_id: schema-workers-binding-kind-text-blob-a47a855e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_text_blob

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "part": {"description": "The name of the file containing the text content. Only accepted for `service worker syntax` Workers.", "type": "string", "example": "my-module.txt", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["text_blob"], "deprecated": true, "x-auditable": true}}, "required": ["name", "type", "part"]}
```
