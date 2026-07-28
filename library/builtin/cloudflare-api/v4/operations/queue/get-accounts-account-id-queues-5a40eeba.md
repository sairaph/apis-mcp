---
title: List Queues
page_id: operation-get-accounts-account-id-queues-a42814e1
path: operations/queue
description: Returns the queues owned by an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/queues
operation_ids:
    - queues-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Queues

`GET /accounts/{account_id}/queues`

Operation ID: `queues-list`

Returns the queues owned by an account.

## Definition

```yaml
{"operationId": "queues-list", "summary": "List Queues", "description": "Returns the queues owned by an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "responses": {"200": {"description": "List of all Queues that belong to this account", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/mq_queue"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Total number of queues", "type": "number", "example": 1}, "page": {"description": "Current page within paginated list of queues", "type": "number", "example": 1}, "per_page": {"description": "Number of queues per page", "type": "number", "example": 20}, "total_count": {"description": "Total queues available without any search parameters", "type": "number", "example": 2000}, "total_pages": {"description": "Total pages available without any search parameters", "type": "number", "example": 100}}}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Queues Read", "Workers Scripts Write", "Workers Scripts Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.edge.worker.queue.list"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
