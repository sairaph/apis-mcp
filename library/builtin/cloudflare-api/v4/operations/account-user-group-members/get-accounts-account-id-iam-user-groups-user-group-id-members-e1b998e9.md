---
title: List User Group Members
page_id: operation-get-accounts-account-id-iam-user-groups-user-group-id-members-85366194
path: operations/account-user-group-members
description: List all the members attached to a user group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/iam/user_groups/{user_group_id}/members
operation_ids:
    - account-user-group-member-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List User Group Members

`GET /accounts/{account_id}/iam/user_groups/{user_group_id}/members`

Operation ID: `account-user-group-member-list`

List all the members attached to a user group.

## Definition

```yaml
{"operationId": "account-user-group-member-list", "summary": "List User Group Members", "description": "List all the members attached to a user group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_user_group_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 100, "maximum": 500, "minimum": 1}}, {"name": "fuzzyEmail", "in": "query", "schema": {"description": "A string used for filtering members by partial email match.", "type": "string", "example": "user@"}}, {"name": "direction", "in": "query", "schema": {"description": "The sort order of returned user group members by email.", "type": "string", "default": "asc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "List User Group Members", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/iam_user_group_member"}}}, "type": "object"}], "title": "List of members attached to a user group"}}}}, "4XX": {"description": "User Group Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account User Group Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
