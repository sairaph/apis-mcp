---
title: zones_cache_level
page_id: schema-zones-cache-level-1870d922
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache_level

```yaml
{"type": "object", "properties": {"id": {"description": "Apply custom caching based on the option selected.\n", "type": "string", "enum": ["cache_level"], "x-auditable": true}, "value": {"description": "* `bypass`: Cloudflare does not cache.\n* `basic`: Delivers resources from cache when there is no query\n  string.\n* `simplified`: Delivers the same resource to everyone independent\n  of the query string.\n* `aggressive`: Caches all static content that has a query string.\n* `cache_everything`: Treats all content as static and caches all\n  file types beyond the [Cloudflare default cached\n  content](https://developers.cloudflare.com/cache/concepts/default-cache-behavior/#default-cached-file-extensions).\n", "type": "string", "example": "bypass", "enum": ["bypass", "basic", "simplified", "aggressive", "cache_everything"], "x-auditable": true}}, "title": "Cache Level", "x-stainless-skip": ["terraform"]}
```
