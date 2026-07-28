---
title: List Roles
page_id: operation-get-accounts-account-id-roles-75508020
path: operations/account-roles
description: Get all available roles for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/roles
operation_ids:
    - account-roles-list-roles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Roles

`GET /accounts/{account_id}/roles`

Operation ID: `account-roles-list-roles`

Get all available roles for an account.

## Definition

```yaml
{"operationId": "account-roles-list-roles", "summary": "List Roles", "description": "Get all available roles for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of roles per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}], "responses": {"200": {"description": "List Roles response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_collection_role_response"}}}}, "4XX": {"description": "List Roles response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}], "tags": ["Account Roles"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-stainless-deprecation-message": "Use /accounts/{account_id}/iam/permission_groups instead."}
```
