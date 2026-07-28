---
title: zones_sort_query_string_for_cache
page_id: schema-zones-sort-query-string-for-cache-ed2d5fae
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_sort_query_string_for_cache

```yaml
{"type": "object", "properties": {"id": {"description": "Turn on or off the reordering of query strings. When query strings have the same structure, caching improves.\n", "type": "string", "example": "sort_query_string_for_cache", "enum": ["sort_query_string_for_cache"], "x-auditable": true}, "value": {"description": "The status of Query String Sort\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "title": "Query String Sort", "x-stainless-skip": ["terraform"]}
```
