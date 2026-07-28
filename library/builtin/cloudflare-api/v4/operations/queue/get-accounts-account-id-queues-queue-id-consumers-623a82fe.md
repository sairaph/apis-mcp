---
title: List Queue Consumers
page_id: operation-get-accounts-account-id-queues-queue-id-consumers-f6400e07
path: operations/queue
description: Returns the consumers for a Queue
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/consumers
operation_ids:
    - queues-list-consumers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Queue Consumers

`GET /accounts/{account_id}/queues/{queue_id}/consumers`

Operation ID: `queues-list-consumers`

Returns the consumers for a Queue

## Definition

```yaml
{"operationId": "queues-list-consumers", "summary": "List Queue Consumers", "description": "Returns the consumers for a Queue", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "responses": {"200": {"description": "All consumers attached to this Queue", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/mq_consumer-response"}}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Queues Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.list"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.consumers", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
