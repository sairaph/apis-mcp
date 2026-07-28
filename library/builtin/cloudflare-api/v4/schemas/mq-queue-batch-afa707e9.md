---
title: mq_queue-batch
page_id: schema-mq-queue-batch-afa707e9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_queue-batch

```yaml
{"type": "object", "properties": {"delay_seconds": {"description": "The number of seconds to wait for attempting to deliver this batch to consumers", "type": "number", "example": "text", "x-auditable": true}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/mq_queue-message"}}}}
```
