---
title: List Permission Groups
page_id: operation-get-accounts-account-id-tokens-permission-groups-538df428
path: operations/account-owned-api-tokens
description: Find all available permission groups for Account Owned API Tokens
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tokens/permission_groups
operation_ids:
    - account-api-tokens-list-permission-groups
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Permission Groups

`GET /accounts/{account_id}/tokens/permission_groups`

Operation ID: `account-api-tokens-list-permission-groups`

Find all available permission groups for Account Owned API Tokens

## Definition

```yaml
{"operationId": "account-api-tokens-list-permission-groups", "summary": "List Permission Groups", "description": "Find all available permission groups for Account Owned API Tokens", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "name", "in": "query", "description": "Filter by the name of the permission group.\nThe value must be URL-encoded.", "schema": {"type": "string", "example": "Account%20Settings%20Write"}}, {"name": "scope", "in": "query", "description": "Filter by the scope of the permission group.\nThe value must be URL-encoded.", "schema": {"type": "string", "example": "com.cloudflare.api.account.zone"}}], "responses": {"200": {"description": "List Account Owned API Token Permission Groups response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_permissions_group_response_collection"}}}}, "4XX": {"description": "List Account Owned API Token Permission Groups response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Account Owned API Tokens"], "x-api-token-group": ["Account API Tokens Write", "Account API Tokens Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.token.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
