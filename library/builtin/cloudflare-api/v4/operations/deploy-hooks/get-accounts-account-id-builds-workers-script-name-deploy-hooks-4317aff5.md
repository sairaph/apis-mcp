---
title: List deploy hooks
page_id: operation-get-accounts-account-id-builds-workers-script-name-deploy-hooks-a2b51bb5
path: operations/deploy-hooks
description: Get all deploy hooks for a specific worker script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks
operation_ids:
    - listDeployHooks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List deploy hooks

`GET /accounts/{account_id}/builds/workers/{script_name}/deploy_hooks`

Operation ID: `listDeployHooks`

Get all deploy hooks for a specific worker script.

## Definition

```yaml
{"operationId": "listDeployHooks", "summary": "List deploy hooks", "description": "Get all deploy hooks for a specific worker script.", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_ScriptName"}], "responses": {"200": {"description": "Deploy hooks retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/builds_DetailedDeployHookResponse"}}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Deploy Hooks"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"]}
```
