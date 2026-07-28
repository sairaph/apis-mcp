---
title: mq_worker-consumer-request
page_id: schema-mq-worker-consumer-request-69ce9e70
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_worker-consumer-request

```yaml
{"type": "object", "properties": {"dead_letter_queue": {"$ref": "#/components/schemas/mq_queue-name"}, "script_name": {"$ref": "#/components/schemas/mq_script-name"}, "settings": {"type": "object", "properties": {"batch_size": {"$ref": "#/components/schemas/mq_batch-size"}, "max_concurrency": {"$ref": "#/components/schemas/mq_max-concurrency"}, "max_retries": {"$ref": "#/components/schemas/mq_max-retries"}, "max_wait_time_ms": {"$ref": "#/components/schemas/mq_max-wait-time"}, "retry_delay": {"$ref": "#/components/schemas/mq_retry-delay"}}}, "type": {"type": "string", "enum": ["worker"], "x-auditable": true}}, "required": ["type", "script_name"]}
```
