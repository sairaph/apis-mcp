---
title: Remove User Group Member
page_id: operation-delete-accounts-account-id-iam-user-groups-user-group-id-members-member-3a4acc8d
path: operations/account-user-group-members
description: Remove a member from User Group
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/iam/user_groups/{user_group_id}/members/{member_id}
operation_ids:
    - account-user-group-member-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove User Group Member

`DELETE /accounts/{account_id}/iam/user_groups/{user_group_id}/members/{member_id}`

Operation ID: `account-user-group-member-delete`

Remove a member from User Group

## Definition

```yaml
{"operationId": "account-user-group-member-delete", "summary": "Remove User Group Member", "description": "Remove a member from User Group", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_identifier"}}, {"name": "member_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_member_identifier"}}], "responses": {"200": {"description": "Delete User Group Member response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/iam_user_group_member"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete User Group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account User Group Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
