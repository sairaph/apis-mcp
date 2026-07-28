---
title: mq_worker-consumer-response
page_id: schema-mq-worker-consumer-response-fa002131
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_worker-consumer-response

```yaml
{"type": "object", "properties": {"consumer_id": {"$ref": "#/components/schemas/mq_identifier"}, "created_on": {"type": "string", "format": "date-time", "x-auditable": true}, "dead_letter_queue": {"description": "Name of the dead letter queue, or empty string if not configured", "type": "string", "x-auditable": true}, "queue_name": {"$ref": "#/components/schemas/mq_queue-name"}, "script_name": {"$ref": "#/components/schemas/mq_script-name"}, "settings": {"type": "object", "properties": {"batch_size": {"$ref": "#/components/schemas/mq_batch-size"}, "max_concurrency": {"$ref": "#/components/schemas/mq_max-concurrency"}, "max_retries": {"$ref": "#/components/schemas/mq_max-retries"}, "max_wait_time_ms": {"$ref": "#/components/schemas/mq_max-wait-time"}, "retry_delay": {"$ref": "#/components/schemas/mq_retry-delay"}}}, "type": {"type": "string", "enum": ["worker"], "x-auditable": true}}}
```
