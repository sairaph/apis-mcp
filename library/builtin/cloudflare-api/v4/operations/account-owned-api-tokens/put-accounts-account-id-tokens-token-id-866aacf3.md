---
title: Update Token
page_id: operation-put-accounts-account-id-tokens-token-id-78de3f7e
path: operations/account-owned-api-tokens
description: Update an existing token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/tokens/{token_id}
operation_ids:
    - account-api-tokens-update-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Token

`PUT /accounts/{account_id}/tokens/{token_id}`

Operation ID: `account-api-tokens-update-token`

Update an existing token.

## Definition

```yaml
{"operationId": "account-api-tokens-update-token", "summary": "Update Token", "description": "Update an existing token.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_token_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_token_body"}}}}, "responses": {"200": {"description": "Update Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_token_response"}}}}, "4XX": {"description": "Update Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Account Owned API Tokens"], "x-api-token-group": ["Account API Tokens Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.token.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
