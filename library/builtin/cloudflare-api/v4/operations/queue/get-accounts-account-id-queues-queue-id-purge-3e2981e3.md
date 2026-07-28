---
title: Get Queue Purge Status
page_id: operation-get-accounts-account-id-queues-queue-id-purge-64f97d61
path: operations/queue
description: Get details about a Queue's purge status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/queues/{queue_id}/purge
operation_ids:
    - queues-purge-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Queue Purge Status

`GET /accounts/{account_id}/queues/{queue_id}/purge`

Operation ID: `queues-purge-get`

Get details about a Queue's purge status.

## Definition

```yaml
{"operationId": "queues-purge-get", "summary": "Get Queue Purge Status", "description": "Get details about a Queue's purge status.", "parameters": [{"name": "queue_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "responses": {"200": {"description": "Details of the requested Queue", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "object", "properties": {"completed": {"description": "Indicates if the last purge operation completed successfully.", "type": "string", "readOnly": true, "x-auditable": true}, "started_at": {"description": "Timestamp when the last purge operation started.", "type": "string", "readOnly": true, "x-auditable": true}}}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Queues Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.read"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.purge", "x-fern-sdk-method-name": "status", "x-forge-hidden": true}
```
