---
title: List Deployments
page_id: operation-get-accounts-account-id-workers-scripts-script-name-deployments-27c0f7c6
path: operations/worker-deployments
description: List of Worker Deployments. The first deployment in the list is the latest deployment actively serving traffic.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/deployments
operation_ids:
    - worker-deployments-list-deployments
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Deployments

`GET /accounts/{account_id}/workers/scripts/{script_name}/deployments`

Operation ID: `worker-deployments-list-deployments`

List of Worker Deployments. The first deployment in the list is the latest deployment actively serving traffic.

## Definition

```yaml
{"operationId": "worker-deployments-list-deployments", "summary": "List Deployments", "description": "List of Worker Deployments. The first deployment in the list is the latest deployment actively serving traffic.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}], "responses": {"200": {"description": "List Deployments response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"type": "object", "properties": {"deployments": {"type": "array", "items": {"$ref": "#/components/schemas/workers_deployment"}}}, "required": ["deployments"]}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "List Deployments response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Deployments"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.deployments", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
