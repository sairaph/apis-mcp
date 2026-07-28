---
title: rulesets_SetCacheSettingsCacheKey
page_id: schema-rulesets-setcachesettingscachekey-b480d07a
path: schemas
description: Which components of the request are included in or excluded from the cache key Cloudflare uses to store the response in cache.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsCacheKey

Which components of the request are included in or excluded from the cache key Cloudflare uses to store the response in cache.

```yaml
{"description": "Which components of the request are included in or excluded from the cache key Cloudflare uses to store the response in cache.", "type": "object", "properties": {"cache_by_device_type": {"description": "Whether to separate cached content based on the visitor's device type.", "type": "boolean", "example": true, "title": "Cache by Device Type"}, "cache_deception_armor": {"description": "Whether to protect from web cache deception attacks, while allowing static assets to be cached.", "type": "boolean", "example": true, "title": "Cache Deception Armor"}, "custom_key": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsCustomCacheKey"}, "ignore_query_strings_order": {"description": "Whether to treat requests with the same query parameters the same, regardless of the order those query parameters are in.", "type": "boolean", "example": true, "title": "Ignore Query Strings Order"}}, "title": "Cache Key"}
```
