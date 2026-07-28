---
title: Delete Tail
page_id: operation-delete-accounts-account-id-workers-scripts-script-name-tails-id-5b61048c
path: operations/worker-tail-logs
description: Deletes a tail from a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/tails/{id}
operation_ids:
    - worker-tail-logs-delete-tail
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Tail

`DELETE /accounts/{account_id}/workers/scripts/{script_name}/tails/{id}`

Operation ID: `worker-tail-logs-delete-tail`

Deletes a tail from a Worker.

## Definition

```yaml
{"operationId": "worker-tail-logs-delete-tail", "summary": "Delete Tail", "description": "Deletes a tail from a Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Tail response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common"}}}}, "4XX": {"description": "Delete Tail response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Tail Logs"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.scripts.tail", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
