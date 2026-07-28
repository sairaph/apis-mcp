---
title: mq_queue-settings
page_id: schema-mq-queue-settings-24b92307
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_queue-settings

```yaml
{"type": "object", "properties": {"delivery_delay": {"description": "Number of seconds to delay delivery of all messages to consumers.", "type": "number", "example": 5, "x-auditable": true}, "delivery_paused": {"description": "Indicates if message delivery to consumers is currently paused.", "type": "boolean", "example": true, "x-auditable": true}, "message_retention_period": {"description": "Number of seconds after which an unconsumed message will be delayed.", "type": "number", "example": 345600, "x-auditable": true}}}
```
