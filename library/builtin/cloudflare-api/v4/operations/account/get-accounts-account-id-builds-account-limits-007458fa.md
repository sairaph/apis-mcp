---
title: Get account limits
page_id: operation-get-accounts-account-id-builds-account-limits-0692d4f1
path: operations/account
description: Retrieve account limits and usage information
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/builds/account/limits
operation_ids:
    - getAccountLimits
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account limits

`GET /accounts/{account_id}/builds/account/limits`

Operation ID: `getAccountLimits`

Retrieve account limits and usage information

## Definition

```yaml
{"operationId": "getAccountLimits", "summary": "Get account limits", "description": "Retrieve account limits and usage information", "parameters": [{"$ref": "#/components/parameters/builds_AccountId"}], "responses": {"200": {"description": "Account limits retrieved successfully", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/builds_APIResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/builds_GetAccountLimitResponse"}}, "type": "object"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account"], "x-api-token-group": ["Workers CI Write", "Workers CI Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers-builds", "x-fern-sdk-method-name": "get-account-limits", "x-forge-hidden": true}
```
