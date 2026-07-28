---
title: workers_binding_kind_plain_text
page_id: schema-workers-binding-kind-plain-text-3914938c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_plain_text

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "text": {"description": "The text value to use.", "type": "string", "example": "Hello, world!", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["plain_text"], "x-auditable": true}}, "required": ["name", "type", "text"]}
```
