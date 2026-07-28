---
title: zones_explicit_cache_control
page_id: schema-zones-explicit-cache-control-de8af5b6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_explicit_cache_control

```yaml
{"type": "object", "properties": {"id": {"description": "Origin Cache Control is enabled by default for Free, Pro, and\nBusiness domains and disabled by default for Enterprise domains.\n", "type": "string", "enum": ["explicit_cache_control"], "x-auditable": true}, "value": {"description": "The status of Origin Cache Control.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Origin Cache Control", "x-stainless-skip": ["terraform"]}
```
