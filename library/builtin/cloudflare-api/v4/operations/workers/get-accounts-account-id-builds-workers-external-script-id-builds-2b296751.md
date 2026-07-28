---
title: List builds by script
page_id: operation-get-accounts-account-id-builds-workers-external-script-id-builds-090e0390
path: operations/workers
description: Get all builds for a specific worker script with pagination
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/workers/{external_script_id}/builds
operation_ids:
    - listBuildsByScript
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List builds by script

`GET /accounts/{account_id}/builds/workers/{external_script_id}/builds`

Operation ID: `listBuildsByScript`

Get all builds for a specific worker script with pagination

## Definition

```yaml
{"operationId": "listBuildsByScript", "summary": "List builds by script", "description": "Get all builds for a specific worker script with pagination", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"name": "external_script_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/builds_external_script_id"}}, {"$ref": "#/components/parameters/builds_Page"}, {"$ref": "#/components/parameters/builds_PerPage"}], "responses": {"200": {"description": "Builds retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/builds_BuildResponse"}}, "result_info": {"$ref": "#/components/schemas/builds_PaginationInfo"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Workers"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.builds", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
