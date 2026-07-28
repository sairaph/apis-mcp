---
title: workers-kv_bulk-result
page_id: schema-workers-kv-bulk-result-5c7088b6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-kv_bulk-result

```yaml
{"type": "object", "properties": {"successful_key_count": {"description": "Number of keys successfully updated.", "type": "number", "example": 100}, "unsuccessful_keys": {"description": "Name of the keys that failed to be fully updated. They should be retried.", "type": "array", "items": {"type": "string"}}}}
```
