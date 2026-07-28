---
title: mq_queue-pull-batch
page_id: schema-mq-queue-pull-batch-ead94e66
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_queue-pull-batch

```yaml
{"type": "array", "items": {"properties": {"attempts": {"type": "number", "example": 1, "readOnly": true, "x-auditable": true}, "body": {"type": "string", "example": "hello world", "readOnly": true}, "id": {"type": "string", "example": "b01b5594f784d0165c2985833f5660dd", "readOnly": true, "x-auditable": true}, "lease_id": {"$ref": "#/components/schemas/mq_lease-id"}, "metadata": {"type": "object", "example": {"CF-Content-Type": "text", "CF-sourceMessageSource": "dash"}, "readOnly": true}, "timestamp_ms": {"type": "number", "example": 1710950954154, "readOnly": true, "x-auditable": true}}, "type": "object"}}
```
