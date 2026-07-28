---
title: Token Details
page_id: operation-get-accounts-account-id-tokens-token-id-35266a0c
path: operations/account-owned-api-tokens
description: Get information about a specific Account Owned API token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tokens/{token_id}
operation_ids:
    - account-api-tokens-token-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Token Details

`GET /accounts/{account_id}/tokens/{token_id}`

Operation ID: `account-api-tokens-token-details`

Get information about a specific Account Owned API token.

## Definition

```yaml
{"operationId": "account-api-tokens-token-details", "summary": "Token Details", "description": "Get information about a specific Account Owned API token.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_token_identifier"}}], "responses": {"200": {"description": "Token Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_token_response"}}}}, "4XX": {"description": "Token Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Account Owned API Tokens"], "x-api-token-group": ["Account API Tokens Write", "Account API Tokens Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.token.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
