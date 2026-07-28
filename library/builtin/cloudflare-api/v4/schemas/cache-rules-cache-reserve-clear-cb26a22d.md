---
title: cache-rules_cache_reserve_clear
page_id: schema-cache-rules-cache-reserve-clear-cb26a22d
path: schemas
description: You can use Cache Reserve Clear to clear your Cache Reserve, but you must first disable Cache Reserve. In most cases, this will be accomplished within 24 hours. You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind that you cannot undo or cancel this operation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_cache_reserve_clear

You can use Cache Reserve Clear to clear your Cache Reserve, but you must first disable Cache Reserve. In most cases, this will be accomplished within 24 hours. You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind that you cannot undo or cancel this operation.

```yaml
{"description": "You can use Cache Reserve Clear to clear your Cache Reserve, but you must first disable Cache Reserve. In most cases, this will be accomplished within 24 hours. You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind that you cannot undo or cancel this operation.", "type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "type": "string", "example": "cache_reserve_clear", "enum": ["cache_reserve_clear"], "x-auditable": true}}, "type": "object"}], "title": "Cache Reserve Clear"}
```
