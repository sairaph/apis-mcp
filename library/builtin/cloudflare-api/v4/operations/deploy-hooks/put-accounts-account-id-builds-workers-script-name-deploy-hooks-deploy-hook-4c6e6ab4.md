---
title: Update deploy hook
page_id: operation-put-accounts-account-id-builds-workers-script-name-deploy-hooks-deploy-h-2b290c0a
path: operations/deploy-hooks
description: Update an existing deploy hook.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks/{deploy_hook_uuid}
operation_ids:
    - updateDeployHook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update deploy hook

`PUT /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks/{deploy_hook_uuid}`

Operation ID: `updateDeployHook`

Update an existing deploy hook.

## Definition

```yaml
{"operationId": "updateDeployHook", "summary": "Update deploy hook", "description": "Update an existing deploy hook.", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_ScriptName"}, {"$ref": "#/components/parameters/builds_DeployHookUuid"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/builds_CreateDeployHookRequest"}}}}, "responses": {"200": {"description": "Deploy hook updated successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_DeployHookResponse"}}, "type": "object"}]}}}}, "400": {"$ref": "#/components/responses/builds_BadRequest"}, "404": {"$ref": "#/components/responses/builds_NotFound"}, "409": {"$ref": "#/components/responses/builds_Conflict"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Deploy Hooks"], "x-api-token-group": ["Workers CI Write"]}
```
