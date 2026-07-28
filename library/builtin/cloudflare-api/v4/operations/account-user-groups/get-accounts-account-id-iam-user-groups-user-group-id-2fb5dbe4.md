---
title: User Group Details
page_id: operation-get-accounts-account-id-iam-user-groups-user-group-id-f49ba2b5
path: operations/account-user-groups
description: Get information about a specific user group in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/iam/user_groups/{user_group_id}
operation_ids:
    - account-user-group-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# User Group Details

`GET /accounts/{account_id}/iam/user_groups/{user_group_id}`

Operation ID: `account-user-group-details`

Get information about a specific user group in an account.

## Definition

```yaml
{"operationId": "account-user-group-details", "summary": "User Group Details", "description": "Get information about a specific user group in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_identifier"}}], "responses": {"200": {"description": "User Group Details response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/iam_user_group"}}, "type": "object"}]}}}}, "4XX": {"description": "User Group Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account User Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
