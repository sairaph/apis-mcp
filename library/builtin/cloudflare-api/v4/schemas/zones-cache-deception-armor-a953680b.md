---
title: zones_cache_deception_armor
page_id: schema-zones-cache-deception-armor-a953680b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache_deception_armor

```yaml
{"type": "object", "properties": {"id": {"description": "Protect from web cache deception attacks while still allowing static\nassets to be cached. This setting verifies that the URL's extension\nmatches the returned `Content-Type`.\n", "type": "string", "enum": ["cache_deception_armor"], "x-auditable": true}, "value": {"description": "The status of Cache Deception Armor.\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Cache Deception Armor", "x-stainless-skip": ["terraform"]}
```
