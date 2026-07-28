---
title: List Tokens
page_id: operation-get-accounts-account-id-tokens-db54ce59
path: operations/account-owned-api-tokens
description: List all Account Owned API tokens created for this account. Results include active, disabled, and recently-expired tokens when include_expired is set to true.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tokens
operation_ids:
    - account-api-tokens-list-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Tokens

`GET /accounts/{account_id}/tokens`

Operation ID: `account-api-tokens-list-tokens`

List all Account Owned API tokens created for this account. Results include active, disabled, and recently-expired tokens when include_expired is set to true.

## Definition

```yaml
{"operationId": "account-api-tokens-list-tokens", "summary": "List Tokens", "description": "List all Account Owned API tokens created for this account. Results include active, disabled, and recently-expired tokens when include_expired is set to true.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order results.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}, {"$ref": "#/components/parameters/iam_include_expired"}], "responses": {"200": {"description": "List Tokens response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_collection_tokens_response"}}}}, "4XX": {"description": "List Tokens response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Account Owned API Tokens"], "x-api-token-group": ["Account API Tokens Write", "Account API Tokens Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.token.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
