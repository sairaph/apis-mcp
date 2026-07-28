---
title: Get Queue
page_id: operation-get-accounts-account-id-queues-queue-id-b1793e33
path: operations/queue
description: Get details about a specific queue.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}
operation_ids:
    - queues-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Queue

`GET /accounts/{account_id}/queues/{queue_id}`

Operation ID: `queues-get`

Get details about a specific queue.

## Definition

```yaml
{"operationId": "queues-get", "summary": "Get Queue", "description": "Get details about a specific queue.", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "responses": {"200": {"description": "Details of the requested Queue", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/mq_queue"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Queues Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.read"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
