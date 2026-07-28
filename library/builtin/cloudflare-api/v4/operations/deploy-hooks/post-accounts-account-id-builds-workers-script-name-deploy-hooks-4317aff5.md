---
title: Create deploy hook
page_id: operation-post-accounts-account-id-builds-workers-script-name-deploy-hooks-e2826e7f
path: operations/deploy-hooks
description: Create a new deploy hook for a worker script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks
operation_ids:
    - createDeployHook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create deploy hook

`POST /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks`

Operation ID: `createDeployHook`

Create a new deploy hook for a worker script.

## Definition

```yaml
{"operationId": "createDeployHook", "summary": "Create deploy hook", "description": "Create a new deploy hook for a worker script.", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_ScriptName"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/builds_CreateDeployHookRequest"}}}}, "responses": {"200": {"description": "Deploy hook created successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_DeployHookResponse"}}, "type": "object"}]}}}}, "400": {"$ref": "#/components/responses/builds_BadRequest"}, "404": {"$ref": "#/components/responses/builds_NotFound"}, "409": {"$ref": "#/components/responses/builds_Conflict"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Deploy Hooks"], "x-api-token-group": ["Workers CI Write"]}
```
