---
title: rulesets_SetCacheSettingsStatusCodeTTL
page_id: schema-rulesets-setcachesettingsstatuscodettl-e69ab687
path: schemas
description: A list of TTLs to apply to specific status codes or status code ranges.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsStatusCodeTTL

A list of TTLs to apply to specific status codes or status code ranges.

```yaml
{"description": "A list of TTLs to apply to specific status codes or status code ranges.", "type": "array", "items": {"maxProperties": 2, "minProperties": 2, "properties": {"status_code": {"description": "A single status code to apply the TTL to.", "type": "integer", "example": 200, "maximum": 999, "minimum": 100, "title": "Status Code"}, "status_code_range": {"description": "A range of status codes to apply the TTL to.", "type": "object", "minProperties": 1, "properties": {"from": {"description": "The lower bound of the range.", "type": "integer", "example": 200, "maximum": 999, "minimum": 100, "title": "From"}, "to": {"description": "The upper bound of the range.", "type": "integer", "example": 299, "maximum": 999, "minimum": 100, "title": "To"}}, "title": "Status Code Range"}, "value": {"description": "The time to cache the response for (in seconds). A value of 0 is equivalent to setting the cache control header with the value \"no-cache\". A value of -1 is equivalent to setting the cache control header with the value of \"no-store\".", "type": "integer", "example": 0, "title": "TTL Value"}}, "required": ["value"], "type": "object"}, "minItems": 1, "title": "Status Code TTLs", "uniqueItems": true}
```
