---
title: zones_sort_query_string_for_cache-2
page_id: schema-zones-sort-query-string-for-cache-2-297f5ff3
path: schemas
description: Cloudflare will treat files with the same query strings as the same file in cache, regardless of the order of the query strings. This is limited to Enterprise Zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_sort_query_string_for_cache-2

Cloudflare will treat files with the same query strings as the same file in cache, regardless of the order of the query strings. This is limited to Enterprise Zones.

```yaml
{"description": "Cloudflare will treat files with the same query strings as the same file in cache, regardless of the order of the query strings. This is limited to Enterprise Zones.", "default": "off", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "sort_query_string_for_cache", "enum": ["sort_query_string_for_cache"]}, "value": {"$ref": "#/components/schemas/zones_sort_query_string_for_cache_value"}}}], "title": "Get String Sort"}
```
