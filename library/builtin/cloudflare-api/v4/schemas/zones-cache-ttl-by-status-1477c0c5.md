---
title: zones_cache_ttl_by_status
page_id: schema-zones-cache-ttl-by-status-1477c0c5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache_ttl_by_status

```yaml
{"type": "object", "properties": {"id": {"description": "Enterprise customers can set cache time-to-live (TTL) based on the\nresponse status from the origin web server. Cache TTL refers to the\nduration of a resource in the Cloudflare network before being\nmarked as stale or discarded from cache. Status codes are returned\nby a resource's origin. Setting cache TTL based on response status\noverrides the default cache behavior (standard caching) for static\nfiles and overrides cache instructions sent by the origin web\nserver. To cache non-static assets, set a Cache Level of Cache\nEverything using a Page Rule. Setting no-store Cache-Control or a\nlow TTL (using `max-age`/`s-maxage`) increases requests to origin\nweb servers and decreases performance.\n", "type": "string", "enum": ["cache_ttl_by_status"], "x-auditable": true}, "value": {"description": "A JSON object containing status codes and their corresponding TTLs.\nEach key-value pair in the cache TTL by status cache rule has the\nfollowing syntax\n- `status_code`: An integer value such as 200 or 500. status_code\n  matches the exact status code from the origin web server. Valid\n  status codes are between 100-999.\n- `status_code_range`: Integer values for from and to.\n  status_code_range matches any status code from the origin web\n  server within the specified range.\n- `value`: An integer value that defines the duration an asset is\n  valid in seconds or one of the following strings: no-store\n  (equivalent to -1), no-cache (equivalent to 0).\n", "type": "object", "example": {"200-299": 86400, "300-499": "no-cache", "500-599": "no-store"}, "additionalProperties": {"anyOf": [{"description": "`no-store` (equivalent to -1), `no-cache` (equivalent to 0)\n", "enum": ["no-cache", "no-store"], "example": "no-cache", "type": "string"}, {"description": "An integer value that defines the duration an asset is valid in\nseconds.\n", "example": 86400, "type": "integer"}]}}}, "title": "Cache TTL by Status Code", "x-stainless-skip": ["terraform"]}
```
