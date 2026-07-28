---
title: Role Details
page_id: operation-get-accounts-account-id-roles-role-id-2604bd91
path: operations/account-roles
description: Get information about a specific role for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/roles/{role_id}
operation_ids:
    - account-roles-role-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Role Details

`GET /accounts/{account_id}/roles/{role_id}`

Operation ID: `account-roles-role-details`

Get information about a specific role for an account.

## Definition

```yaml
{"operationId": "account-roles-role-details", "summary": "Role Details", "description": "Get information about a specific role for an account.", "parameters": [{"name": "role_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_role_components-schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "responses": {"200": {"description": "Role Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_role_response"}}}}, "4XX": {"description": "Role Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}], "tags": ["Account Roles"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-stainless-deprecation-message": "Use /accounts/{account_id}/iam/permission_groups/{permission_group_id} instead."}
```
