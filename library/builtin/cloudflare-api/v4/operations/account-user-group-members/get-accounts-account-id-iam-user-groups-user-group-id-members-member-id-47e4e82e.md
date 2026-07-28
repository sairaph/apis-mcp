---
title: Get User Group Member
page_id: operation-get-accounts-account-id-iam-user-groups-user-group-id-members-member-id-471bc50e
path: operations/account-user-group-members
description: Get details of a specific member in a user group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/iam/user_groups/{user_group_id}/members/{member_id}
operation_ids:
    - account-user-group-member-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get User Group Member

`GET /accounts/{account_id}/iam/user_groups/{user_group_id}/members/{member_id}`

Operation ID: `account-user-group-member-get`

Get details of a specific member in a user group.

## Definition

```yaml
{"operationId": "account-user-group-member-get", "summary": "Get User Group Member", "description": "Get details of a specific member in a user group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_identifier"}}, {"name": "member_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_member_identifier"}}], "responses": {"200": {"description": "Get User Group Member response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/iam_user_group_member_detailed"}}, "type": "object"}]}}}}, "4XX": {"description": "Get User Group Member response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account User Group Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
