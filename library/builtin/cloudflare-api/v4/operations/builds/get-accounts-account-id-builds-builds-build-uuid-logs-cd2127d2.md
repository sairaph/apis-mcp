---
title: Get build logs
page_id: operation-get-accounts-account-id-builds-builds-build-uuid-logs-bee8ce6a
path: operations/builds
description: Retrieve logs for a specific build with cursor-based pagination
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/builds/{build_uuid}/logs
operation_ids:
    - getBuildLogs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get build logs

`GET /accounts/{account_id}/builds/builds/{build_uuid}/logs`

Operation ID: `getBuildLogs`

Retrieve logs for a specific build with cursor-based pagination

## Definition

```yaml
{"operationId": "getBuildLogs", "summary": "Get build logs", "description": "Retrieve logs for a specific build with cursor-based pagination", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_BuildUuid"}, {"name": "cursor", "in": "query", "schema": {"$ref": "#/components/schemas/builds_cursor"}}], "responses": {"200": {"description": "Build logs retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_BuildLogsResponse"}}, "type": "object"}]}}}}, "404": {"$ref": "#/components/responses/builds_NotFound"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Builds"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.builds.logs", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
