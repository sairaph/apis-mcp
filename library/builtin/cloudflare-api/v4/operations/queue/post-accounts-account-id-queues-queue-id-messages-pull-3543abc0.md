---
title: Pull Queue Messages
page_id: operation-post-accounts-account-id-queues-queue-id-messages-pull-ebe411bf
path: operations/queue
description: Pull a batch of messages from a Queue
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/messages/pull
operation_ids:
    - queues-pull-messages
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Pull Queue Messages

`POST /accounts/{account_id}/queues/{queue_id}/messages/pull`

Operation ID: `queues-pull-messages`

Pull a batch of messages from a Queue

## Definition

```yaml
{"operationId": "queues-pull-messages", "summary": "Pull Queue Messages", "description": "Pull a batch of messages from a Queue", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"batch_size": {"$ref": "#/components/schemas/mq_batch-size"}, "visibility_timeout_ms": {"$ref": "#/components/schemas/mq_visibility-timeout"}}}}}}, "responses": {"200": {"description": "A batch of messages in the Queue", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "object", "properties": {"message_backlog_count": {"description": "The number of unacknowledged messages in the queue.", "type": "number", "example": 5}, "messages": {"$ref": "#/components/schemas/mq_queue-pull-batch"}, "metadata": {"type": "object", "properties": {"metrics": {"$ref": "#/components/schemas/mq_queue-metrics"}}}}}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-auditable": false, "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.pull"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.messages", "x-fern-sdk-method-name": "pull", "x-forge-hidden": true}
```
