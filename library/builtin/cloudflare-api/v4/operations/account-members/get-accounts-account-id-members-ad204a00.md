---
title: List Members
page_id: operation-get-accounts-account-id-members-edb67f61
path: operations/account-members
description: List all members of an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/members
operation_ids:
    - account-members-list-members
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Members

`GET /accounts/{account_id}/members`

Operation ID: `account-members-list-members`

List all members of an account.

## Definition

```yaml
{"operationId": "account-members-list-members", "summary": "List Members", "description": "List all members of an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "order", "in": "query", "schema": {"description": "Field to order results by.", "example": "status", "enum": ["user.first_name", "user.last_name", "user.email", "status"]}}, {"name": "status", "in": "query", "schema": {"description": "A member's status in the account.", "type": "string", "example": "accepted", "enum": ["accepted", "pending", "rejected"]}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order results.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "List Members response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_collection_member_response_with_policies"}}}}, "4XX": {"description": "List Members response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Account Members"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["#organization:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
