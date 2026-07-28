---
title: cache-purge_FlexPurgeByTags
page_id: schema-cache-purge-flexpurgebytags-0fbeea04
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-purge_FlexPurgeByTags

```yaml
{"type": "object", "properties": {"tags": {"description": "For more information on cache tags and purging by tags, please refer to [purge by cache-tags documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-tags/).", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["a-cache-tag", "another-cache-tag"]}}, "title": "Purge by tags"}
```
