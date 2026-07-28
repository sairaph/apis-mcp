---
title: Push Message
page_id: operation-post-accounts-account-id-queues-queue-id-messages-6d987d40
path: operations/queue
description: Push a message to a Queue
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/messages
operation_ids:
    - queues-push-message
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Push Message

`POST /accounts/{account_id}/queues/{queue_id}/messages`

Operation ID: `queues-push-message`

Push a message to a Queue

## Definition

```yaml
{"operationId": "queues-push-message", "summary": "Push Message", "description": "Push a message to a Queue", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_queue-message"}}}}, "responses": {"200": {"description": "Successful message ingestion.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "object", "properties": {"metadata": {"type": "object", "properties": {"metrics": {"$ref": "#/components/schemas/mq_queue-metrics"}}}}}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-auditable": false, "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.push"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.messages", "x-fern-sdk-method-name": "push", "x-forge-hidden": true}
```
