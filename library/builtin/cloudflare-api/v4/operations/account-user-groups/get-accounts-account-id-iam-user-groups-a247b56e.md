---
title: List User Groups
page_id: operation-get-accounts-account-id-iam-user-groups-d55bd4d3
path: operations/account-user-groups
description: List all the user groups for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/iam/user_groups
operation_ids:
    - account-user-group-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List User Groups

`GET /accounts/{account_id}/iam/user_groups`

Operation ID: `account-user-group-list`

List all the user groups for an account.

## Definition

```yaml
{"operationId": "account-user-group-list", "summary": "List User Groups", "description": "List all the user groups for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "id", "in": "query", "description": "ID of the user group to be fetched.", "schema": {"$ref": "#/components/schemas/iam_user_group_identifier"}}, {"name": "name", "in": "query", "schema": {"description": "Name of the user group to be fetched.", "type": "string", "example": "NameOfTheUserGroup"}}, {"name": "fuzzyName", "in": "query", "schema": {"description": "A string used for searching for user groups containing that substring.", "type": "string", "example": "Foo"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 100, "maximum": 500, "minimum": 1}}, {"name": "direction", "in": "query", "schema": {"description": "The sort order of returned user groups by name (ascending or descending).", "type": "string", "example": "desc", "default": "asc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "List User Group response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/iam_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/iam_user_groups"}}, "type": "object"}]}}}}, "4XX": {"description": "List User Group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_email": [], "api_key": []}], "tags": ["Account User Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
