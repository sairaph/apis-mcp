---
title: workers_binding_kind_ai_search_namespace
page_id: schema-workers-binding-kind-ai-search-namespace-a4fdfcda
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_ai_search_namespace

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "namespace": {"description": "The user-chosen namespace name. Must exist before deploy -- Wrangler handles auto-creation on deploy failure (R2 bucket pattern). The \"default\" namespace is auto-created by config-api for new accounts. Grants full access (CRUD + search + chat) to all instances within the namespace.", "type": "string", "example": "production", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["ai_search_namespace"], "x-auditable": true}}, "required": ["name", "type", "namespace"]}
```
