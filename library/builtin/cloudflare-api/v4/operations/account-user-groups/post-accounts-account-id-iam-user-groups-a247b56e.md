---
title: Create User Group
page_id: operation-post-accounts-account-id-iam-user-groups-cdae3c76
path: operations/account-user-groups
description: Create a new user group under the specified account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/iam/user_groups
operation_ids:
    - account-user-group-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create User Group

`POST /accounts/{account_id}/iam/user_groups`

Operation ID: `account-user-group-create`

Create a new user group under the specified account.

## Definition

```yaml
{"operationId": "account-user-group-create", "summary": "Create User Group", "description": "Create a new user group under the specified account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_create_user_group_body"}}}}, "responses": {"200": {"description": "Add User Group response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/iam_user_group"}}, "type": "object"}]}}}}, "4XX": {"description": "Add User Group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account User Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
