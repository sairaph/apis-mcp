---
title: List build tokens
page_id: operation-get-accounts-account-id-builds-tokens-939a1b31
path: operations/build-tokens
description: Get all build tokens with pagination
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/tokens
operation_ids:
    - listBuildTokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List build tokens

`GET /accounts/{account_id}/builds/tokens`

Operation ID: `listBuildTokens`

Get all build tokens with pagination

## Definition

```yaml
{"operationId": "listBuildTokens", "summary": "List build tokens", "description": "Get all build tokens with pagination", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}, {"$ref": "#/components/parameters/builds_Page"}, {"$ref": "#/components/parameters/builds_PerPage"}], "responses": {"200": {"description": "Build tokens retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/builds_ListTokensResponse"}}, "result_info": {"$ref": "#/components/schemas/builds_PaginationInfo"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Build Tokens"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.tokens", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
