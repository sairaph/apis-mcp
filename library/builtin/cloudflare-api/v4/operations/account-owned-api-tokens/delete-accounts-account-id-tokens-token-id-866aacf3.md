---
title: Delete Token
page_id: operation-delete-accounts-account-id-tokens-token-id-1706ace7
path: operations/account-owned-api-tokens
description: Destroy an Account Owned API token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/tokens/{token_id}
operation_ids:
    - account-api-tokens-delete-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Token

`DELETE /accounts/{account_id}/tokens/{token_id}`

Operation ID: `account-api-tokens-delete-token`

Destroy an Account Owned API token.

## Definition

```yaml
{"operationId": "account-api-tokens-delete-token", "summary": "Delete Token", "description": "Destroy an Account Owned API token.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_token_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single-id"}}}}, "4XX": {"description": "Delete Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Account Owned API Tokens"], "x-api-token-group": ["Account API Tokens Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.token.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
