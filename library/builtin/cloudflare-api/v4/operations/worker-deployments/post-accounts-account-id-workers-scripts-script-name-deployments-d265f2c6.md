---
title: Create Deployment
page_id: operation-post-accounts-account-id-workers-scripts-script-name-deployments-0864cbf5
path: operations/worker-deployments
description: Deployments configure how [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions) are deployed to traffic. A deployment can consist of one or two versions of a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/scripts/{script_name}/deployments
operation_ids:
    - worker-deployments-create-deployment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Deployment

`POST /accounts/{account_id}/workers/scripts/{script_name}/deployments`

Operation ID: `worker-deployments-create-deployment`

Deployments configure how [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions) are deployed to traffic. A deployment can consist of one or two versions of a Worker.

## Definition

```yaml
{"operationId": "worker-deployments-create-deployment", "summary": "Create Deployment", "description": "Deployments configure how [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions) are deployed to traffic. A deployment can consist of one or two versions of a Worker.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "script_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_script_name"}}, {"name": "force", "in": "query", "description": "If set to true, the deployment will be created even if normally blocked by something such rolling back to an older version when a secret has changed.", "schema": {"type": "boolean"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_deployment"}}}}, "responses": {"200": {"description": "Create Deployment response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_deployment"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Create Deployment response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Deployments"], "x-api-token-group": ["Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.deployments", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
