---
title: Preview Queue Messages
page_id: operation-post-accounts-account-id-queues-queue-id-messages-preview-9d4ee0d5
path: operations/queue
description: Preview messages from a Queue without leasing them. Messages remain available for subsequent preview or pull operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/messages/preview
operation_ids:
    - queues-preview-messages
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Preview Queue Messages

`POST /accounts/{account_id}/queues/{queue_id}/messages/preview`

Operation ID: `queues-preview-messages`

Preview messages from a Queue without leasing them. Messages remain available for subsequent preview or pull operations.

## Definition

```yaml
{"operationId": "queues-preview-messages", "summary": "Preview Queue Messages", "description": "Preview messages from a Queue without leasing them. Messages remain available for subsequent preview or pull operations.", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"batch_size": {"$ref": "#/components/schemas/mq_batch-size"}}}}}}, "responses": {"200": {"description": "A batch of previewed messages from the Queue", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "object", "properties": {"messages": {"$ref": "#/components/schemas/mq_queue-pull-batch"}}}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Queues Read", "Workers Scripts Write", "Workers Scripts Read"], "x-auditable": false, "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.peek"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.messages", "x-fern-sdk-method-name": "preview", "x-forge-hidden": true}
```
