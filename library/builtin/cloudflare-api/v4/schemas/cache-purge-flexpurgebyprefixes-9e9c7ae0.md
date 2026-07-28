---
title: cache-purge_FlexPurgeByPrefixes
page_id: schema-cache-purge-flexpurgebyprefixes-9e9c7ae0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-purge_FlexPurgeByPrefixes

```yaml
{"type": "object", "properties": {"prefixes": {"description": "For more information on purging by prefixes, please refer to [purge by prefix documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge_by_prefix/).", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["www.example.com/foo", "images.example.com/bar/baz"]}}, "title": "Purge by prefixes"}
```
