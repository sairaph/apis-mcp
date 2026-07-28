---
title: hyperdrive_hyperdrive-caching-enabled
page_id: schema-hyperdrive-hyperdrive-caching-enabled-8fe0760e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_hyperdrive-caching-enabled

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/hyperdrive_hyperdrive-caching-common"}, {"properties": {"max_age": {"description": "Specify the maximum duration (in seconds) items should persist in the cache. Defaults to 60 seconds if not specified.", "type": "integer", "example": 60, "x-auditable": true}, "stale_while_revalidate": {"description": "Specify the number of seconds the cache may serve a stale response. Defaults to 15 seconds if not specified.", "type": "integer", "example": 15, "x-auditable": true}}, "type": "object"}]}
```
