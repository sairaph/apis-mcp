---
title: workers_binding_kind_flagship
page_id: schema-workers-binding-kind-flagship-07bfc50b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_flagship

```yaml
{"type": "object", "properties": {"app_id": {"description": "ID of the Flagship app to bind to for feature flag evaluation.", "type": "string", "example": "app-12345678-1234-1234-1234-123456789012", "maxLength": 128, "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["flagship"], "x-auditable": true}}, "required": ["name", "type", "app_id"]}
```
