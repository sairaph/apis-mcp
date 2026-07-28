---
title: rulesets_SetCacheSettingsVaryDefault
page_id: schema-rulesets-setcachesettingsvarydefault-c6e099ab
path: schemas
description: Controls how response Vary headers without a per-header override contribute to the cache key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsVaryDefault

Controls how response Vary headers without a per-header override contribute to the cache key.

```yaml
{"description": "Controls how response Vary headers without a per-header override contribute to the cache key.", "type": "object", "properties": {"action": {"description": "How the header value is treated when building the cache key.", "type": "string", "example": "normalize", "enum": ["bypass", "passthrough", "normalize"], "title": "Action"}}, "required": ["action"], "title": "Vary Default"}
```
