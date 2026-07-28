---
title: zones_browser_cache_ttl
page_id: schema-zones-browser-cache-ttl-1b932900
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_browser_cache_ttl

```yaml
{"type": "object", "properties": {"id": {"description": "Control how long resources cached by client browsers remain valid.\n", "type": "string", "enum": ["browser_cache_ttl"], "x-auditable": true}, "value": {"description": "The number of seconds to cache resources for.\nSetting this to 0 enables \"Respect Existing Headers\".\n", "type": "integer", "example": 3600, "maximum": 31536000, "minimum": 0, "x-auditable": true}}, "title": "Browser Cache TTL", "x-stainless-skip": ["terraform"]}
```
