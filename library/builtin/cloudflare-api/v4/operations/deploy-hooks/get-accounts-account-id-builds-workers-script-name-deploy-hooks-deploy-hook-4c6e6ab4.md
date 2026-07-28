---
title: Get deploy hook
page_id: operation-get-accounts-account-id-builds-workers-script-name-deploy-hooks-deploy-h-fe58acf9
path: operations/deploy-hooks
description: Get details of a specific deploy hook.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks/{deploy_hook_uuid}
operation_ids:
    - getDeployHook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get deploy hook

`GET /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks/{deploy_hook_uuid}`

Operation ID: `getDeployHook`

Get details of a specific deploy hook.

## Definition

```yaml
{"operationId": "getDeployHook", "summary": "Get deploy hook", "description": "Get details of a specific deploy hook.", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_ScriptName"}, {"$ref": "#/components/parameters/builds_DeployHookUuid"}], "responses": {"200": {"description": "Deploy hook retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_DeployHookResponse"}}, "type": "object"}]}}}}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Deploy Hooks"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"]}
```
