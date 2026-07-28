---
title: Update worker build configuration
page_id: operation-patch-accounts-account-id-builds-workers-script-tag-2ad8dd4f
path: operations/workers
description: Update the build configuration for a Worker script. Supports partial updates to git repository settings and production build settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/builds/workers/{script_tag}
operation_ids:
    - updateWorkerBuild
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update worker build configuration

`PATCH /accounts/{account_id}/builds/workers/{script_tag}`

Operation ID: `updateWorkerBuild`

Update the build configuration for a Worker script. Supports partial updates to git repository settings and production build settings.

## Definition

```yaml
{"operationId": "updateWorkerBuild", "summary": "Update worker build configuration", "description": "Update the build configuration for a Worker script. Supports partial updates to git repository settings and production build settings.", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_ScriptTag"}], "requestBody": {"required": true, "content": {"application/json": {"example": {"production_settings": {"build_command": "npm run build:production", "deploy_command": "npx wrangler deploy --env production"}}, "schema": {"$ref": "#/components/schemas/builds_UpdateWorkerRequest"}}}}, "responses": {"200": {"description": "Worker build configuration updated successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_WorkerResponse"}}, "type": "object"}]}}}}, "400": {"$ref": "#/components/responses/builds_BadRequest"}, "401": {"$ref": "#/components/responses/builds_Unauthorized"}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers CI Write"]}
```
