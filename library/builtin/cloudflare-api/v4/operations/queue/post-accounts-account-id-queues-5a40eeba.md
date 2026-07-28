---
title: Create Queue
page_id: operation-post-accounts-account-id-queues-25fff895
path: operations/queue
description: Create a new queue
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/queues
operation_ids:
    - queues-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Queue

`POST /accounts/{account_id}/queues`

Operation ID: `queues-create`

Create a new queue

## Definition

```yaml
{"operationId": "queues-create", "summary": "Create Queue", "description": "Create a new queue", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"queue_name": {"$ref": "#/components/schemas/mq_queue-name"}}, "required": ["queue_name"]}}}}, "responses": {"200": {"description": "Created Queue", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/mq_queue"}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.create"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
