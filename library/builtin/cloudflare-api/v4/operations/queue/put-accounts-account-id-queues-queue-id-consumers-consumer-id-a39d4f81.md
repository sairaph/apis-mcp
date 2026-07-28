---
title: Update Queue Consumer
page_id: operation-put-accounts-account-id-queues-queue-id-consumers-consumer-id-04052f02
path: operations/queue
description: Updates the consumer for a queue, or creates one if it does not exist.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/consumers/{consumer_id}
operation_ids:
    - queues-update-consumer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Queue Consumer

`PUT /accounts/{account_id}/queues/{queue_id}/consumers/{consumer_id}`

Operation ID: `queues-update-consumer`

Updates the consumer for a queue, or creates one if it does not exist.

## Definition

```yaml
{"operationId": "queues-update-consumer", "summary": "Update Queue Consumer", "description": "Updates the consumer for a queue, or creates one if it does not exist.", "parameters": [{"name": "consumer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_consumer-request"}}}}, "responses": {"200": {"description": "Update Queue Consumer response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/mq_consumer-response"}}, "type": "object"}]}}}}, "4XX": {"description": "Update Queue Consumer response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.update"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.consumers", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
