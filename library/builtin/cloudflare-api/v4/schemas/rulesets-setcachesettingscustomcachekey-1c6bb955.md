---
title: rulesets_SetCacheSettingsCustomCacheKey
page_id: schema-rulesets-setcachesettingscustomcachekey-1c6bb955
path: schemas
description: Which components of the request are included or excluded from the cache key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsCustomCacheKey

Which components of the request are included or excluded from the cache key.

```yaml
{"description": "Which components of the request are included or excluded from the cache key.", "type": "object", "properties": {"cookie": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsCustomCacheKeyCookie"}, "header": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsCustomCacheKeyHeader"}, "host": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsCustomCacheKeyHost"}, "query_string": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsCustomCacheKeyQueryString"}, "user": {"$ref": "#/components/schemas/rulesets_SetCacheSettingsCustomCacheKeyUser"}}, "title": "Custom Cache Key"}
```
