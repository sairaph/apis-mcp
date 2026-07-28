---
title: rulesets_SetCacheSettingsCustomCacheKeyHeader
page_id: schema-rulesets-setcachesettingscustomcachekeyheader-4675445a
path: schemas
description: Which headers to include in the cache key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsCustomCacheKeyHeader

Which headers to include in the cache key.

```yaml
{"description": "Which headers to include in the cache key.", "type": "object", "properties": {"check_presence": {"description": "A list of headers to check for the presence of. The presence of these headers is included in the cache key.", "type": "array", "items": {"description": "The name of the header to check for the presence of.", "example": "my-header", "minLength": 1, "title": "Header", "type": "string"}, "minItems": 1, "title": "Check Presence", "uniqueItems": true}, "contains": {"description": "A mapping of header names to a list of values. If a header is present in the request and contains any of the values provided, its value is included in the cache key.", "type": "object", "example": {"my-header": ["my-header-value-1", "my-header-value-2"]}, "additionalProperties": {"description": "A list of values to match the header against.", "items": {"description": "The header value to match against.", "example": "my-header-value", "minLength": 1, "title": "Header Value", "type": "string"}, "minItems": 1, "title": "Header Values", "type": "array", "uniqueItems": true}, "minProperties": 1, "title": "Contains"}, "exclude_origin": {"description": "Whether to exclude the origin header in the cache key.", "type": "boolean", "example": true, "title": "Exclude Origin"}, "include": {"description": "A list of headers to include in the cache key.", "type": "array", "items": {"description": "The name of the header to include.", "example": "my-header", "minLength": 1, "title": "Header", "type": "string"}, "minItems": 1, "title": "Include", "uniqueItems": true}}, "title": "Headers"}
```
