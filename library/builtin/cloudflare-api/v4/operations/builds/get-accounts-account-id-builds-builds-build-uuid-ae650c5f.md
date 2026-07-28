---
title: Get build by UUID
page_id: operation-get-accounts-account-id-builds-builds-build-uuid-6c7f9362
path: operations/builds
description: Retrieve detailed information about a specific build
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/builds/{build_uuid}
operation_ids:
    - getBuildByUuid
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get build by UUID

`GET /accounts/{account_id}/builds/builds/{build_uuid}`

Operation ID: `getBuildByUuid`

Retrieve detailed information about a specific build

## Definition

```yaml
{"operationId": "getBuildByUuid", "summary": "Get build by UUID", "description": "Retrieve detailed information about a specific build", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_BuildUuid"}], "responses": {"200": {"description": "Build retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_BuildResponse"}}, "type": "object"}]}}}}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Builds"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.builds", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
