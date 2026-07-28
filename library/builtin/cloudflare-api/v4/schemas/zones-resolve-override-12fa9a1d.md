---
title: zones_resolve_override
page_id: schema-zones-resolve-override-12fa9a1d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_resolve_override

```yaml
{"type": "object", "properties": {"id": {"description": "Change the origin address to the value specified in this setting.\n", "type": "string", "example": "resolve_override", "enum": ["resolve_override"], "x-auditable": true}, "value": {"description": "The origin address you want to override with.\n", "type": "string", "example": "example.com", "x-auditable": true}}, "title": "Resolve Override", "x-stainless-skip": ["terraform"]}
```
