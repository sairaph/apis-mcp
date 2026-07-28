---
title: mcn_observation
page_id: schema-mcn-observation-879aa665
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_observation

```yaml
{"type": "object", "properties": {"first_observed_at": {"type": "string", "x-auditable": true}, "last_observed_at": {"type": "string", "x-auditable": true}, "provider_id": {"$ref": "#/components/schemas/mcn_provider_id"}, "resource_id": {"$ref": "#/components/schemas/mcn_resource_id"}}, "required": ["provider_id", "resource_id", "first_observed_at", "last_observed_at"]}
```
