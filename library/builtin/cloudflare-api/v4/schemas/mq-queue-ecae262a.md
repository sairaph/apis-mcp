---
title: mq_queue
page_id: schema-mq-queue-ecae262a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_queue

```yaml
{"type": "object", "properties": {"consumers": {"type": "array", "items": {"$ref": "#/components/schemas/mq_consumer-response"}, "readOnly": true}, "consumers_total_count": {"type": "number", "readOnly": true}, "created_on": {"type": "string", "readOnly": true, "x-auditable": true}, "modified_on": {"type": "string", "readOnly": true, "x-auditable": true}, "producers": {"type": "array", "items": {"$ref": "#/components/schemas/mq_producer"}, "readOnly": true}, "producers_total_count": {"type": "number", "readOnly": true, "x-auditable": true}, "queue_id": {"type": "string", "readOnly": true, "x-auditable": true}, "queue_name": {"$ref": "#/components/schemas/mq_queue-name"}, "settings": {"$ref": "#/components/schemas/mq_queue-settings"}}}
```
