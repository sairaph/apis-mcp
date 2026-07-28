---
title: Update Queue
page_id: operation-patch-accounts-account-id-queues-queue-id-b30e5139
path: operations/queue
description: Updates a Queue.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}
operation_ids:
    - queues-update-partial
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Queue

`PATCH /accounts/{account_id}/queues/{queue_id}`

Operation ID: `queues-update-partial`

Updates a Queue.

## Definition

```yaml
{"operationId": "queues-update-partial", "summary": "Update Queue", "description": "Updates a Queue.", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_queue"}}}}, "responses": {"200": {"description": "Updated Queue", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_queue"}]}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.update"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
