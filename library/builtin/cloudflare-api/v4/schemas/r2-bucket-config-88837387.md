---
title: r2_bucket-config
page_id: schema-r2-bucket-config-88837387
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_bucket-config

```yaml
{"type": "object", "properties": {"bucketName": {"description": "Name of the bucket.", "type": "string", "x-auditable": true}, "queues": {"description": "List of queues associated with the bucket.", "type": "array", "items": {"$ref": "#/components/schemas/r2_queues-config"}}}}
```
