---
title: Delete deploy hook
page_id: operation-delete-accounts-account-id-builds-workers-script-name-deploy-hooks-deplo-ccbc9645
path: operations/deploy-hooks
description: Delete a deploy hook.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks/{deploy_hook_uuid}
operation_ids:
    - deleteDeployHook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete deploy hook

`DELETE /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks/{deploy_hook_uuid}`

Operation ID: `deleteDeployHook`

Delete a deploy hook.

## Definition

```yaml
{"operationId": "deleteDeployHook", "summary": "Delete deploy hook", "description": "Delete a deploy hook.", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_ScriptName"}, {"$ref": "#/components/parameters/builds_DeployHookUuid"}], "responses": {"200": {"description": "Deploy hook deleted successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_DeployHookResponse"}}, "type": "object"}]}}}}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Deploy Hooks"], "x-api-token-group": ["Workers CI Write"]}
```
