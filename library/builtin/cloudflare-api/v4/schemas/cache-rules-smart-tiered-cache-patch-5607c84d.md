---
title: cache-rules_smart_tiered_cache_patch
page_id: schema-cache-rules-smart-tiered-cache-patch-5607c84d
path: schemas
description: Update enablement of Smart Tiered Cache.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_smart_tiered_cache_patch

Update enablement of Smart Tiered Cache.

```yaml
{"description": "Update enablement of Smart Tiered Cache.", "type": "object", "properties": {"value": {"description": "Enable or disable the Smart Tiered Cache.", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "required": ["value"]}
```
