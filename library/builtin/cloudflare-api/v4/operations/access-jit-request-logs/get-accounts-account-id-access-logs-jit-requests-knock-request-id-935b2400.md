---
title: Get an Access JIT request log
page_id: operation-get-accounts-account-id-access-logs-jit-requests-knock-request-id-c663c41f
path: operations/access-jit-request-logs
description: Gets an account-scoped Access JIT request summary and its lifecycle events in chronological order.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/logs/jit_requests/{knock_request_id}
operation_ids:
    - access-jit-request-logs-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access JIT request log

`GET /accounts/{account_id}/access/logs/jit_requests/{knock_request_id}`

Operation ID: `access-jit-request-logs-get`

Gets an account-scoped Access JIT request summary and its lifecycle events in chronological order.

## Definition

```yaml
{"operationId": "access-jit-request-logs-get", "summary": "Get an Access JIT request log", "description": "Gets an account-scoped Access JIT request summary and its lifecycle events in chronological order.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "knock_request_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Access JIT request log detail response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_single"}}}}, "404": {"description": "Access JIT request log not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}, "4XX": {"description": "Access JIT request log detail response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Access JIT request logs"]}
```
