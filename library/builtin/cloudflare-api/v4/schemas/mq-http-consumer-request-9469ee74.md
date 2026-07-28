---
title: mq_http-consumer-request
page_id: schema-mq-http-consumer-request-9469ee74
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_http-consumer-request

```yaml
{"type": "object", "properties": {"dead_letter_queue": {"$ref": "#/components/schemas/mq_queue-name"}, "settings": {"type": "object", "properties": {"batch_size": {"$ref": "#/components/schemas/mq_batch-size"}, "max_retries": {"$ref": "#/components/schemas/mq_max-retries"}, "retry_delay": {"$ref": "#/components/schemas/mq_retry-delay"}, "visibility_timeout_ms": {"$ref": "#/components/schemas/mq_visibility-timeout"}}}, "type": {"type": "string", "enum": ["http_pull"], "x-auditable": true}}, "required": ["type"]}
```
