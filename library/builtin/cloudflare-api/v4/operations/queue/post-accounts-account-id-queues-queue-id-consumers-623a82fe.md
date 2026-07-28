---
title: Create a Queue Consumer
page_id: operation-post-accounts-account-id-queues-queue-id-consumers-c84f6926
path: operations/queue
description: Creates a new consumer for a Queue
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/consumers
operation_ids:
    - queues-create-consumer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Queue Consumer

`POST /accounts/{account_id}/queues/{queue_id}/consumers`

Operation ID: `queues-create-consumer`

Creates a new consumer for a Queue

## Definition

```yaml
{"operationId": "queues-create-consumer", "summary": "Create a Queue Consumer", "description": "Creates a new consumer for a Queue", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_consumer-request"}}}}, "responses": {"200": {"description": "Create Queue Consumer response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/mq_consumer-response"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.create"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.consumers", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
