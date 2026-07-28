---
title: Delete Deployment
page_id: operation-delete-accounts-account-id-workers-scripts-script-name-deployments-deplo-8958b514
path: operations/worker-deployments
description: Delete a Worker Deployment. The latest deployment, which is actively serving traffic, cannot be deleted. All other deployments can be deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/deployments/{deployment_id}
operation_ids:
    - worker-deployments-delete-deployment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Deployment

`DELETE /accounts/{account_id}/workers/scripts/{script_name}/deployments/{deployment_id}`

Operation ID: `worker-deployments-delete-deployment`

Delete a Worker Deployment. The latest deployment, which is actively serving traffic, cannot be deleted. All other deployments can be deleted.

## Definition

```yaml
{"operationId": "worker-deployments-delete-deployment", "summary": "Delete Deployment", "description": "Delete a Worker Deployment. The latest deployment, which is actively serving traffic, cannot be deleted. All other deployments can be deleted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "deployment_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Delete Deployment response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common"}}}}, "4XX": {"description": "Delete Deployment response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Deployments"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.deployments", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
