---
title: Delete worker build configuration
page_id: operation-delete-accounts-account-id-builds-workers-script-tag-ed3c08c6
path: operations/workers
description: Delete the build configuration for a Worker script.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/builds/workers/{script_tag}
operation_ids:
    - deleteWorkerBuild
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete worker build configuration

`DELETE /accounts/{account_id}/builds/workers/{script_tag}`

Operation ID: `deleteWorkerBuild`

Delete the build configuration for a Worker script.

## Definition

```yaml
{"operationId": "deleteWorkerBuild", "summary": "Delete worker build configuration", "description": "Delete the build configuration for a Worker script.", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_ScriptTag"}], "responses": {"200": {"description": "Worker build configuration deleted successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"type": "string", "example": "ok"}}, "type": "object"}]}}}}, "401": {"$ref": "#/components/responses/builds_Unauthorized"}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers CI Write"]}
```
