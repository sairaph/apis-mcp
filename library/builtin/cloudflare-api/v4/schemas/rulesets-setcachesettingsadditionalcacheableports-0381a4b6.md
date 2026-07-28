---
title: rulesets_SetCacheSettingsAdditionalCacheablePorts
page_id: schema-rulesets-setcachesettingsadditionalcacheableports-0381a4b6
path: schemas
description: A list of additional ports that caching should be enabled on.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsAdditionalCacheablePorts

A list of additional ports that caching should be enabled on.

```yaml
{"description": "A list of additional ports that caching should be enabled on.", "type": "array", "items": {"description": "A port to enable caching on.", "example": 8080, "maximum": 65535, "minimum": 1, "title": "Additional Cacheable Port", "type": "integer"}, "minItems": 1, "title": "Additional Cacheable Ports (Enterprise-Only)", "uniqueItems": true}
```
