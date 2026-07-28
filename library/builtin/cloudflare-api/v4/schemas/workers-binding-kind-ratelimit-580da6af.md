---
title: workers_binding_kind_ratelimit
page_id: schema-workers-binding-kind-ratelimit-580da6af
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_ratelimit

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "namespace_id": {"description": "Identifier of the rate limit namespace to bind to.", "type": "string", "example": "1234", "x-auditable": true}, "simple": {"description": "The rate limit configuration.", "type": "object", "properties": {"limit": {"description": "The limit (requests per period).", "type": "number", "example": 100, "x-auditable": true}, "mitigation_timeout": {"description": "Duration in seconds to apply the mitigation action after the rate limit is exceeded. Valid values are 0 (disabled), 10, or multiples of 60 up to 86400. Must be greater than or equal to the period when non-zero.\n", "type": "integer", "example": 60, "x-auditable": true}, "period": {"description": "The period in seconds.", "type": "integer", "example": 60, "x-auditable": true}}, "required": ["limit", "period"]}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["ratelimit"], "x-auditable": true}}, "required": ["name", "type", "namespace_id", "simple"]}
```
