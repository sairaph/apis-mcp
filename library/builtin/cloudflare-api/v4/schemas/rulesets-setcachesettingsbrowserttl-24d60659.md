---
title: rulesets_SetCacheSettingsBrowserTTL
page_id: schema-rulesets-setcachesettingsbrowserttl-24d60659
path: schemas
description: How long client browsers should cache the response. Cloudflare cache purge will not purge content cached on client browsers, so high browser TTLs may lead to stale content.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsBrowserTTL

How long client browsers should cache the response. Cloudflare cache purge will not purge content cached on client browsers, so high browser TTLs may lead to stale content.

```yaml
{"description": "How long client browsers should cache the response. Cloudflare cache purge will not purge content cached on client browsers, so high browser TTLs may lead to stale content.", "type": "object", "properties": {"default": {"description": "The browser TTL (in seconds) if you choose the \"override_origin\" mode.", "type": "integer", "example": 60, "minimum": 0, "title": "Default TTL"}, "mode": {"description": "The browser TTL mode.", "type": "string", "example": "override_origin", "enum": ["respect_origin", "bypass_by_default", "override_origin", "bypass"], "title": "TTL Mode"}}, "required": ["mode"], "title": "Browser TTL"}
```
