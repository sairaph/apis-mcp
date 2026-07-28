---
title: Create build token
page_id: operation-post-accounts-account-id-builds-tokens-389e2e0b
path: operations/build-tokens
description: Create a new build authentication token
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/builds/tokens
operation_ids:
    - createBuildToken
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create build token

`POST /accounts/{account_id}/builds/tokens`

Operation ID: `createBuildToken`

Create a new build authentication token

## Definition

```yaml
{"operationId": "createBuildToken", "summary": "Create build token", "description": "Create a new build authentication token", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/builds_CreateBuildTokenRequest"}}}}, "responses": {"200": {"description": "Build token created successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_CreateBuildTokenResponse"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Build Tokens"], "x-api-token-group": ["Workers CI Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds.tokens", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
