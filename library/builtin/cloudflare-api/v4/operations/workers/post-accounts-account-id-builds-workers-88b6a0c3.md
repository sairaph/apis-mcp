---
title: Create worker build configuration
page_id: operation-post-accounts-account-id-builds-workers-20a80b4b
path: operations/workers
description: Create a new build configuration for a Worker script, linking it to a git repository with CI/CD triggers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/builds/workers
operation_ids:
    - createWorkerBuild
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create worker build configuration

`POST /accounts/{account_id}/builds/workers`

Operation ID: `createWorkerBuild`

Create a new build configuration for a Worker script, linking it to a git repository with CI/CD triggers.

## Definition

```yaml
{"operationId": "createWorkerBuild", "summary": "Create worker build configuration", "description": "Create a new build configuration for a Worker script, linking it to a git repository with CI/CD triggers.", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}], "requestBody": {"required": true, "content": {"application/json": {"example": {"git_repository": {"branch": "main", "provider_account_id": "cloudflare", "provider_account_name": "Cloudflare", "provider_type": "github", "repo_id": "workers-sdk", "repo_name": "workers-sdk"}, "production_settings": {"build_caching_enabled": true, "build_command": "npm run build", "build_token_uuid": "your-build-token-uuid", "deploy_command": "npx wrangler deploy", "environment_variables": {}, "path_excludes": ["*.md"], "path_includes": ["*"], "root_directory": "/"}, "script_tag": "my-worker"}, "schema": {"$ref": "#/components/schemas/builds_CreateWorkerRequest"}}}}, "responses": {"201": {"description": "Worker build configuration created successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_WorkerResponse"}}, "type": "object"}]}}}}, "400": {"$ref": "#/components/responses/builds_BadRequest"}, "401": {"$ref": "#/components/responses/builds_Unauthorized"}, "404": {"$ref": "#/components/responses/builds_NotFound"}, "409": {"$ref": "#/components/responses/builds_Conflict"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers CI Write"]}
```
