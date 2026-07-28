---
title: List Tails
page_id: operation-get-accounts-account-id-workers-scripts-script-name-tails-99a78a34
path: operations/worker-tail-logs
description: Get list of tails currently deployed on a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/tails
operation_ids:
    - worker-tail-logs-list-tails
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Tails

`GET /accounts/{account_id}/workers/scripts/{script_name}/tails`

Operation ID: `worker-tail-logs-list-tails`

Get list of tails currently deployed on a Worker.

## Definition

```yaml
{"operationId": "worker-tail-logs-list-tails", "summary": "List Tails", "description": "Get list of tails currently deployed on a Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "List Tails response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_tail"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "List Tails response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Tail Logs"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write"]}
```
