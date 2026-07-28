---
title: rulesets_SetCacheSettingsCacheReserve
page_id: schema-rulesets-setcachesettingscachereserve-827dfc04
path: schemas
description: Settings to determine whether the request's response from origin is eligible for Cache Reserve (requires a Cache Reserve add-on plan).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsCacheReserve

Settings to determine whether the request's response from origin is eligible for Cache Reserve (requires a Cache Reserve add-on plan).

```yaml
{"description": "Settings to determine whether the request's response from origin is eligible for Cache Reserve (requires a Cache Reserve add-on plan).", "type": "object", "properties": {"eligible": {"description": "Whether Cache Reserve is enabled. If this is true and a request meets eligibility criteria, Cloudflare will write the resource to Cache Reserve.", "type": "boolean", "example": true, "title": "Eligible"}, "minimum_file_size": {"description": "The minimum file size eligible for storage in Cache Reserve.", "type": "integer", "example": 1024, "minimum": 0, "title": "Minimum File Size"}}, "required": ["eligible"], "title": "Cache Reserve"}
```
