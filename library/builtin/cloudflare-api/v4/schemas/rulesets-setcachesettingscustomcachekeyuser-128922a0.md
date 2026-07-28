---
title: rulesets_SetCacheSettingsCustomCacheKeyUser
page_id: schema-rulesets-setcachesettingscustomcachekeyuser-128922a0
path: schemas
description: How to use characteristics of the request user agent in the cache key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsCustomCacheKeyUser

How to use characteristics of the request user agent in the cache key.

```yaml
{"description": "How to use characteristics of the request user agent in the cache key.", "type": "object", "properties": {"device_type": {"description": "Whether to use the user agent's device type in the cache key.", "type": "boolean", "example": true, "title": "Device Type"}, "geo": {"description": "Whether to use the user agents's country in the cache key.", "type": "boolean", "example": true, "title": "Country"}, "lang": {"description": "Whether to use the user agent's language in the cache key.", "type": "boolean", "example": true, "title": "Language"}}, "title": "User"}
```
