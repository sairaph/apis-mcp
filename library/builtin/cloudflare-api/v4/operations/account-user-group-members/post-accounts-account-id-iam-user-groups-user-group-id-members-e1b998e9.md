---
title: Add User Group Members
page_id: operation-post-accounts-account-id-iam-user-groups-user-group-id-members-616f71a5
path: operations/account-user-group-members
description: Add members to a User Group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/iam/user_groups/{user_group_id}/members
operation_ids:
    - account-user-group-member-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add User Group Members

`POST /accounts/{account_id}/iam/user_groups/{user_group_id}/members`

Operation ID: `account-user-group-member-create`

Add members to a User Group.

## Definition

```yaml
{"operationId": "account-user-group-member-create", "summary": "Add User Group Members", "description": "Add members to a User Group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/iam_user_group_member_identifier"}}, "required": ["id"], "type": "object"}}}}}, "responses": {"200": {"description": "Add User Group Members response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/iam_user_group_member"}}}, "type": "object"}]}}}}, "4XX": {"description": "Add User Group Members response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account User Group Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
