---
title: zones_edge_cache_ttl
page_id: schema-zones-edge-cache-ttl-80001f16
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_edge_cache_ttl

```yaml
{"type": "object", "properties": {"id": {"description": "Specify how long to cache a resource in the Cloudflare global\nnetwork. *Edge Cache TTL* is not visible in response headers.\n", "type": "string", "enum": ["edge_cache_ttl"], "x-auditable": true}, "value": {"type": "integer", "maximum": 31536000, "minimum": 1, "x-auditable": true}}, "title": "Edge Cache TTL", "x-stainless-skip": ["terraform"]}
```
