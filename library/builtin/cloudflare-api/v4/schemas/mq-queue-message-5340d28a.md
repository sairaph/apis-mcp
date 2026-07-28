---
title: mq_queue-message
page_id: schema-mq-queue-message-5340d28a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_queue-message

```yaml
{"allOf": [{"properties": {"delay_seconds": {"description": "The number of seconds to wait for attempting to deliver this message to consumers", "type": "number", "example": "text", "x-auditable": true}}, "type": "object"}, {"oneOf": [{"$ref": "#/components/schemas/mq_queue-message-text"}, {"$ref": "#/components/schemas/mq_queue-message-json"}]}]}
```
