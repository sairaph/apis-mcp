---
title: mq_http-consumer-response
page_id: schema-mq-http-consumer-response-77cbb0a7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_http-consumer-response

```yaml
{"type": "object", "properties": {"consumer_id": {"$ref": "#/components/schemas/mq_identifier"}, "created_on": {"type": "string", "format": "date-time", "x-auditable": true}, "dead_letter_queue": {"description": "Name of the dead letter queue, or empty string if not configured", "type": "string", "x-auditable": true}, "queue_name": {"$ref": "#/components/schemas/mq_queue-name"}, "settings": {"type": "object", "properties": {"batch_size": {"$ref": "#/components/schemas/mq_batch-size"}, "max_retries": {"$ref": "#/components/schemas/mq_max-retries"}, "retry_delay": {"$ref": "#/components/schemas/mq_retry-delay"}, "visibility_timeout_ms": {"$ref": "#/components/schemas/mq_visibility-timeout"}}}, "type": {"type": "string", "enum": ["http_pull"], "x-auditable": true}}}
```
