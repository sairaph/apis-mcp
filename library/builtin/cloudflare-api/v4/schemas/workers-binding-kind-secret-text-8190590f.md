---
title: workers_binding_kind_secret_text
page_id: schema-workers-binding-kind-secret-text-8190590f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_secret_text

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "text": {"description": "The secret value to use.", "type": "string", "example": "My secret.", "writeOnly": true, "x-sensitive": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["secret_text"], "x-auditable": true}}, "required": ["name", "type", "text"]}
```
