---
title: mq_event-destination-queue
page_id: schema-mq-event-destination-queue-b14a7361
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_event-destination-queue

```yaml
{"type": "object", "properties": {"queue_id": {"description": "ID of the target queue", "type": "string", "x-auditable": true}, "type": {"description": "Type of destination", "type": "string", "enum": ["queues.queue"], "x-auditable": true}}, "required": ["type", "queue_id"]}
```
