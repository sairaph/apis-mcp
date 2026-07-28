---
title: Get Deployment
page_id: operation-get-accounts-account-id-workers-scripts-script-name-deployments-deployme-5390c517
path: operations/worker-deployments
description: Get information about a Worker Deployment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/deployments/{deployment_id}
operation_ids:
    - worker-deployments-get-deployment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Deployment

`GET /accounts/{account_id}/workers/scripts/{script_name}/deployments/{deployment_id}`

Operation ID: `worker-deployments-get-deployment`

Get information about a Worker Deployment.

## Definition

```yaml
{"operationId": "worker-deployments-get-deployment", "summary": "Get Deployment", "description": "Get information about a Worker Deployment.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "deployment_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Get Deployment response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_deployment"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get Deployment response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Deployments"], "x-api-token-group": ["Workers Tail Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.deployments", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
