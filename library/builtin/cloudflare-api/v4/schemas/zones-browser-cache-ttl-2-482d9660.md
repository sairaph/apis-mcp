---
title: zones_browser_cache_ttl-2
page_id: schema-zones-browser-cache-ttl-2-482d9660
path: schemas
description: Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources will remain on your visitors' computers. Cloudflare will honor any larger times specified by your server. (https://support.cloudflare.com/hc/en-us/articles/200168276).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_browser_cache_ttl-2

Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources will remain on your visitors' computers. Cloudflare will honor any larger times specified by your server. (https://support.cloudflare.com/hc/en-us/articles/200168276).

```yaml
{"description": "Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources will remain on your visitors' computers. Cloudflare will honor any larger times specified by your server. (https://support.cloudflare.com/hc/en-us/articles/200168276).", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "browser_cache_ttl", "enum": ["browser_cache_ttl"]}, "value": {"$ref": "#/components/schemas/zones_browser_cache_ttl_value"}}}], "title": "Browser Cache TTL"}
```
