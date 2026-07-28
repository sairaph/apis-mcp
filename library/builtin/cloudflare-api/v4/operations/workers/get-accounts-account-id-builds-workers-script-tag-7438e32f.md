---
title: Get worker build configuration
page_id: operation-get-accounts-account-id-builds-workers-script-tag-e929768b
path: operations/workers
description: Retrieve the build configuration for a specific Worker script, including git repository details and production settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/workers/{script_tag}
operation_ids:
    - getWorkerBuild
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get worker build configuration

`GET /accounts/{account_id}/builds/workers/{script_tag}`

Operation ID: `getWorkerBuild`

Retrieve the build configuration for a specific Worker script, including git repository details and production settings.

## Definition

```yaml
{"operationId": "getWorkerBuild", "summary": "Get worker build configuration", "description": "Retrieve the build configuration for a specific Worker script, including git repository details and production settings.", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_ScriptTag"}], "responses": {"200": {"description": "Worker build configuration retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_WorkerResponse"}}, "type": "object"}]}}}}, "401": {"$ref": "#/components/responses/builds_Unauthorized"}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"]}
```
