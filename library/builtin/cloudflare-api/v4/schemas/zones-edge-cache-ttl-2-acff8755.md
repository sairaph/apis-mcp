---
title: zones_edge_cache_ttl-2
page_id: schema-zones-edge-cache-ttl-2-acff8755
path: schemas
description: Time (in seconds) that a resource will be ensured to remain on Cloudflare's cache servers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_edge_cache_ttl-2

Time (in seconds) that a resource will be ensured to remain on Cloudflare's cache servers.

```yaml
{"description": "Time (in seconds) that a resource will be ensured to remain on Cloudflare's cache servers.", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "edge_cache_ttl", "enum": ["edge_cache_ttl"]}, "value": {"$ref": "#/components/schemas/zones_edge_cache_ttl_value"}}}], "title": "Edge Cache TTL"}
```
