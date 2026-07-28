---
title: cache-rules_cache_reserve_clear_response_value
page_id: schema-cache-rules-cache-reserve-clear-response-value-1196a036
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_cache_reserve_clear_response_value

```yaml
{"type": "object", "properties": {"result": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_cache_reserve_clear"}, {"properties": {"end_ts": {"$ref": "#/components/schemas/cache-rules_cache_reserve_clear_end_ts"}, "start_ts": {"$ref": "#/components/schemas/cache-rules_cache_reserve_clear_start_ts"}, "state": {"$ref": "#/components/schemas/cache-rules_cache_reserve_clear_state"}}, "required": ["state", "start_ts"], "type": "object"}]}}}
```
