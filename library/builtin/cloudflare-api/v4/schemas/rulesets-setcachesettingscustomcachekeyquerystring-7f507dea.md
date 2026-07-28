---
title: rulesets_SetCacheSettingsCustomCacheKeyQueryString
page_id: schema-rulesets-setcachesettingscustomcachekeyquerystring-7f507dea
path: schemas
description: Which query string parameters to include in or exclude from the cache key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SetCacheSettingsCustomCacheKeyQueryString

Which query string parameters to include in or exclude from the cache key.

```yaml
{"description": "Which query string parameters to include in or exclude from the cache key.", "type": "object", "properties": {"exclude": {"description": "Which query string parameters to exclude from the cache key.", "type": "object", "maxProperties": 1, "minProperties": 1, "properties": {"all": {"description": "Whether to exclude all query string parameters from the cache key.", "type": "boolean", "enum": [true], "title": "Exclude All"}, "list": {"description": "A list of query string parameters to exclude from the cache key.", "type": "array", "items": {"description": "The name of the query string parameter to exclude.", "example": "foo", "minLength": 1, "title": "Parameter Name", "type": "string"}, "minItems": 1, "title": "Exclude List", "uniqueItems": true}}, "title": "Exclude"}, "include": {"description": "Which query string parameters to include in the cache key.", "type": "object", "maxProperties": 1, "minProperties": 1, "properties": {"all": {"description": "Whether to include all query string parameters in the cache key.", "type": "boolean", "enum": [true], "title": "Include All"}, "list": {"description": "A list of query string parameters to include in the cache key.", "type": "array", "items": {"description": "The name of the query string parameter to include.", "example": "foo", "minLength": 1, "title": "Parameter Name", "type": "string"}, "minItems": 1, "title": "Include List", "uniqueItems": true}}, "title": "Include"}}, "maxProperties": 1, "title": "Query String Parameters"}
```
