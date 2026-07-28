---
title: Update User Group Members
page_id: operation-put-accounts-account-id-iam-user-groups-user-group-id-members-83a6e2d7
path: operations/account-user-group-members
description: Replace the set of members attached to a User Group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/iam/user_groups/{user_group_id}/members
operation_ids:
    - account-user-group-members-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update User Group Members

`PUT /accounts/{account_id}/iam/user_groups/{user_group_id}/members`

Operation ID: `account-user-group-members-update`

Replace the set of members attached to a User Group.

## Definition

```yaml
{"operationId": "account-user-group-members-update", "summary": "Update User Group Members", "description": "Replace the set of members attached to a User Group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"description": "Set/Replace members to a user group.", "type": "array", "items": {"properties": {"id": {"$ref": "#/components/schemas/iam_user_group_member_identifier"}}, "required": ["id"], "type": "object"}, "title": "Update User Group Members"}}}}, "responses": {"200": {"description": "Update User Group Members response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-single"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/iam_user_group_member"}}}, "type": "object"}]}}}}, "4XX": {"description": "Update User Group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account User Group Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
