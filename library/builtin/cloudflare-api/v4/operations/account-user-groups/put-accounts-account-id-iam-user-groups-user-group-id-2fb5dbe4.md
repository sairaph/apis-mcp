---
title: Update User Group
page_id: operation-put-accounts-account-id-iam-user-groups-user-group-id-b38b22ba
path: operations/account-user-groups
description: Modify an existing user group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/iam/user_groups/{user_group_id}
operation_ids:
    - account-user-group-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update User Group

`PUT /accounts/{account_id}/iam/user_groups/{user_group_id}`

Operation ID: `account-user-group-update`

Modify an existing user group.

## Definition

```yaml
{"operationId": "account-user-group-update", "summary": "Update User Group", "description": "Modify an existing user group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_update_user_group_body"}}}}, "responses": {"200": {"description": "Update User Group response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/iam_user_group"}}, "type": "object"}]}}}}, "4XX": {"description": "Update User Group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account User Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
