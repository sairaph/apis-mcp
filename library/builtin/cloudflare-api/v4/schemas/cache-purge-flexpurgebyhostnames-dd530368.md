---
title: cache-purge_FlexPurgeByHostnames
page_id: schema-cache-purge-flexpurgebyhostnames-dd530368
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-purge_FlexPurgeByHostnames

```yaml
{"type": "object", "properties": {"hosts": {"description": "For more information purging by hostnames, please refer to [purge by hostname documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-hostname/).", "type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["www.example.com", "images.example.com"]}}, "title": "Purge by hostnames"}
```
