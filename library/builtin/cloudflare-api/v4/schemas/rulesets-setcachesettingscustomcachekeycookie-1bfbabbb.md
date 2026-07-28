---
title: rulesets_SetCacheSettingsCustomCacheKeyCookie
page_id: schema-rulesets-setcachesettingscustomcachekeycookie-1bfbabbb
path: schemas
description: Which cookies to include in the cache key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsCustomCacheKeyCookie

Which cookies to include in the cache key.

```yaml
{"description": "Which cookies to include in the cache key.", "type": "object", "properties": {"check_presence": {"description": "A list of cookies to check for the presence of. The presence of these cookies is included in the cache key.", "type": "array", "items": {"description": "The name of the cookie to check for the presence of.", "example": "myCookie", "minLength": 1, "title": "Cookie Name", "type": "string"}, "minItems": 1, "title": "Check Presence", "uniqueItems": true}, "include": {"description": "A list of cookies to include in the cache key.", "type": "array", "items": {"description": "The name of the cookie to include.", "example": "myCookie", "minLength": 1, "title": "Cookie Name", "type": "string"}, "minItems": 1, "title": "Include", "uniqueItems": true}}, "title": "Cookies"}
```
