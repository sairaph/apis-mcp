---
title: Start Tail
page_id: operation-post-accounts-account-id-workers-scripts-script-name-tails-060809ae
path: operations/worker-tail-logs
description: Starts a tail that receives logs and exception from a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/tails
operation_ids:
    - worker-tail-logs-start-tail
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start Tail

`POST /accounts/{account_id}/workers/scripts/{script_name}/tails`

Operation ID: `worker-tail-logs-start-tail`

Starts a tail that receives logs and exception from a Worker.

## Definition

```yaml
{"operationId": "worker-tail-logs-start-tail", "summary": "Start Tail", "description": "Starts a tail that receives logs and exception from a Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Start Tail response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_tail"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Start Tail response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Tail Logs"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.tail", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
