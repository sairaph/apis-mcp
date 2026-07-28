---
title: rulesets_SetCacheSettingsServeStale
page_id: schema-rulesets-setcachesettingsservestale-a7b213c7
path: schemas
description: When to serve stale content from cache.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsServeStale

When to serve stale content from cache.

```yaml
{"description": "When to serve stale content from cache.", "type": "object", "properties": {"disable_stale_while_updating": {"description": "Whether Cloudflare should disable serving stale content while getting the latest content from the origin.", "type": "boolean", "example": true, "title": "Disable Stale While Updating"}}, "title": "Serve Stale"}
```
