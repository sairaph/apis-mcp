---
title: Roll Token
page_id: operation-put-accounts-account-id-tokens-token-id-value-358bd928
path: operations/account-owned-api-tokens
description: Roll the Account Owned API token secret.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/tokens/{token_id}/value
operation_ids:
    - account-api-tokens-roll-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Roll Token

`PUT /accounts/{account_id}/tokens/{token_id}/value`

Operation ID: `account-api-tokens-roll-token`

Roll the Account Owned API token secret.

## Definition

```yaml
{"operationId": "account-api-tokens-roll-token", "summary": "Roll Token", "description": "Roll the Account Owned API token secret.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_token_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object"}}}}, "responses": {"200": {"description": "Roll Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_response_single_value"}}}}, "4XX": {"description": "Roll Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Account Owned API Tokens"], "x-api-token-group": ["Account API Tokens Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.token.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
