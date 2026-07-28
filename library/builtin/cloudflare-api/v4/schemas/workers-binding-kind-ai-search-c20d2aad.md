---
title: workers_binding_kind_ai_search
page_id: schema-workers-binding-kind-ai-search-c20d2aad
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_ai_search

```yaml
{"type": "object", "properties": {"instance_name": {"description": "The user-chosen instance name. Must exist at deploy time. The worker can search, chat, update, and manage items/jobs on this instance.", "type": "string", "example": "cloudflare-blog", "x-auditable": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "namespace": {"description": "The namespace the instance belongs to. Defaults to \"default\" if omitted. Customers who don't use namespaces can simply omit this field.", "type": "string", "example": "production", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["ai_search"], "x-auditable": true}}, "required": ["name", "type", "instance_name"]}
```
