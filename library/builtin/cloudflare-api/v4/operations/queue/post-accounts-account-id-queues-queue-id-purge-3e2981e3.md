---
title: Purge Queue
page_id: operation-post-accounts-account-id-queues-queue-id-purge-b9a2cbf5
path: operations/queue
description: Deletes all messages from the Queue.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/purge
operation_ids:
    - queues-purge
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Purge Queue

`POST /accounts/{account_id}/queues/{queue_id}/purge`

Operation ID: `queues-purge`

Deletes all messages from the Queue.

## Definition

```yaml
{"operationId": "queues-purge", "summary": "Purge Queue", "description": "Deletes all messages from the Queue.", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"delete_messages_permanently": {"description": "Confimation that all messages will be deleted permanently.", "type": "boolean", "example": true, "x-auditable": true}}}}}}, "responses": {"200": {"description": "Updated Queue", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_queue"}]}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.ack"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.purge", "x-fern-sdk-method-name": "start", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation drops every message in the queue."}
```
