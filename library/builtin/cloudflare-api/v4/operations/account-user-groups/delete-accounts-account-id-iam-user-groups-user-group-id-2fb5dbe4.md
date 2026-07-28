---
title: Remove User Group
page_id: operation-delete-accounts-account-id-iam-user-groups-user-group-id-32d470fe
path: operations/account-user-groups
description: Remove a user group from an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/iam/user_groups/{user_group_id}
operation_ids:
    - account-user-group-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove User Group

`DELETE /accounts/{account_id}/iam/user_groups/{user_group_id}`

Operation ID: `account-user-group-delete`

Remove a user group from an account.

## Definition

```yaml
{"operationId": "account-user-group-delete", "summary": "Remove User Group", "description": "Remove a user group from an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_identifier"}}], "responses": {"200": {"description": "Remove User Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single-id"}}}}, "4XX": {"description": "Remove User Group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account User Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
