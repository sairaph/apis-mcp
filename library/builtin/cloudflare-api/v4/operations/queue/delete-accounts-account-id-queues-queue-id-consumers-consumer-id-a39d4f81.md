---
title: Delete Queue Consumer
page_id: operation-delete-accounts-account-id-queues-queue-id-consumers-consumer-id-3e6d52a2
path: operations/queue
description: Deletes the consumer for a queue.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/consumers/{consumer_id}
operation_ids:
    - queues-delete-consumer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Queue Consumer

`DELETE /accounts/{account_id}/queues/{queue_id}/consumers/{consumer_id}`

Operation ID: `queues-delete-consumer`

Deletes the consumer for a queue.

## Definition

```yaml
{"operationId": "queues-delete-consumer", "summary": "Delete Queue Consumer", "description": "Deletes the consumer for a queue.", "parameters": [{"name": "consumer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "responses": {"200": {"description": "Successful consumer delete", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-success"}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.delete"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.consumers", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
