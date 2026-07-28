---
title: Create SCIM User
page_id: operation-post-accounts-account-id-scim-v2-users-72df3846
path: operations/scim-users
description: Provisions a new account member via SCIM. The `userName` field must be a valid email address and must match the primary email in `emails`. The account must be an Enterprise account with SCIM entitlements enabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/scim/v2/Users
operation_ids:
    - scim-users-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create SCIM User

`POST /accounts/{account_id}/scim/v2/Users`

Operation ID: `scim-users-create`

Provisions a new account member via SCIM. The `userName` field must be a valid email address and must match the primary email in `emails`. The account must be an Enterprise account with SCIM entitlements enabled.

## Definition

```yaml
{"operationId": "scim-users-create", "summary": "Create SCIM User", "description": "Provisions a new account member via SCIM. The `userName` field must be a valid email address and must match the primary email in `emails`. The account must be an Enterprise account with SCIM entitlements enabled.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_user_create_request"}}}}, "responses": {"201": {"description": "Create SCIM User response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_user"}}}}, "4XX": {"description": "Create SCIM User response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Users"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.create"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
