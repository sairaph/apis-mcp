---
title: List Token Permission Groups
page_id: operation-get-user-tokens-permission-groups-6439e4bd
path: operations/user-api-tokens
description: Find all available permission groups for API Tokens.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/tokens/permission_groups
operation_ids:
    - permission-groups-list-permission-groups
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Token Permission Groups

`GET /user/tokens/permission_groups`

Operation ID: `permission-groups-list-permission-groups`

Find all available permission groups for API Tokens.

## Definition

```yaml
{"operationId": "permission-groups-list-permission-groups", "summary": "List Token Permission Groups", "description": "Find all available permission groups for API Tokens.", "parameters": [{"name": "name", "in": "query", "description": "Filter by the name of the permission group.\nThe value must be URL-encoded.", "schema": {"type": "string", "example": "Account%20Settings%20Write"}}, {"name": "scope", "in": "query", "description": "Filter by the scope of the permission group.\nThe value must be URL-encoded.", "schema": {"type": "string", "example": "com.cloudflare.api.account.zone"}}], "responses": {"200": {"description": "List Token Permission Groups response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_permissions_group_response_collection"}}}}, "4XX": {"description": "List Token Permission Groups response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["User API Tokens"], "x-api-token-group": ["API Tokens Write", "API Tokens Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.token.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
