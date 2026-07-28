---
title: Get Queue Consumer
page_id: operation-get-accounts-account-id-queues-queue-id-consumers-consumer-id-71ca9169
path: operations/queue
description: Fetches the consumer for a queue by consumer id
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/consumers/{consumer_id}
operation_ids:
    - queues-get-consumer
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Queue Consumer

`GET /accounts/{account_id}/queues/{queue_id}/consumers/{consumer_id}`

Operation ID: `queues-get-consumer`

Fetches the consumer for a queue by consumer id

## Definition

```yaml
{"operationId": "queues-get-consumer", "summary": "Get Queue Consumer", "description": "Fetches the consumer for a queue by consumer id", "parameters": [{"name": "consumer_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "responses": {"200": {"description": "Get Queue Consumer response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/mq_consumer-response"}}, "type": "object"}]}}}}, "4XX": {"description": "Get Queue Consumer response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Queues Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.read"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.consumers", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
