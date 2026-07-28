---
title: mq_queue-metrics
page_id: schema-mq-queue-metrics-b15e530a
path: schemas
description: Best-effort metrics for the queue. Values may be approximate due to the distributed nature of queues.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mq_queue-metrics

Best-effort metrics for the queue. Values may be approximate due to the distributed nature of queues.

```yaml
{"description": "Best-effort metrics for the queue. Values may be approximate due to the distributed nature of queues.", "type": "object", "properties": {"backlog_bytes": {"description": "The size in bytes of unacknowledged messages in the queue.", "type": "number", "example": 1024, "x-auditable": true}, "backlog_count": {"description": "The number of unacknowledged messages in the queue.", "type": "number", "example": 5, "x-auditable": true}, "oldest_message_timestamp_ms": {"description": "Unix timestamp in milliseconds of the oldest unacknowledged message in the queue. Returns 0 if unknown.", "type": "number", "example": 1710950954154, "x-auditable": true}}, "required": ["backlog_count", "backlog_bytes", "oldest_message_timestamp_ms"]}
```
