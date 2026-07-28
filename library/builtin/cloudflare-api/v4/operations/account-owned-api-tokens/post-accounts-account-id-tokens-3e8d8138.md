---
title: Create Token
page_id: operation-post-accounts-account-id-tokens-37687659
path: operations/account-owned-api-tokens
description: Create a new Account Owned API token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/tokens
operation_ids:
    - account-api-tokens-create-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Token

`POST /accounts/{account_id}/tokens`

Operation ID: `account-api-tokens-create-token`

Create a new Account Owned API token.

## Definition

```yaml
{"operationId": "account-api-tokens-create-token", "summary": "Create Token", "description": "Create a new Account Owned API token.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_create_payload"}}}}, "responses": {"200": {"description": "Create Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_token_create_response"}}}}, "4XX": {"description": "Create Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Account Owned API Tokens"], "x-api-token-group": ["Account API Tokens Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.token.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
