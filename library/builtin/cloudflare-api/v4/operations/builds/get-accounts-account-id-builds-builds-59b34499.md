---
title: Get builds by version IDs
page_id: operation-get-accounts-account-id-builds-builds-cbeb1d12
path: operations/builds
description: Retrieve builds for specific version IDs
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/builds
operation_ids:
    - getBuildsByVersionIds
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get builds by version IDs

`GET /accounts/{account_id}/builds/builds`

Operation ID: `getBuildsByVersionIds`

Retrieve builds for specific version IDs

## Definition

```yaml
{"operationId": "getBuildsByVersionIds", "summary": "Get builds by version IDs", "description": "Retrieve builds for specific version IDs", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"name": "version_ids", "in": "query", "required": true, "schema": {"$ref": "#/components/schemas/builds_version_ids"}}], "responses": {"200": {"description": "Builds retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_BuildsByVersionResponse"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Builds"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds", "x-fern-sdk-method-name": "get-builds-by-version", "x-forge-hidden": true}
```
