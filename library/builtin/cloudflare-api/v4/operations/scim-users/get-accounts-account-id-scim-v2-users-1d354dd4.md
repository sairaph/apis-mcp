---
title: List SCIM Users
page_id: operation-get-accounts-account-id-scim-v2-users-e4874093
path: operations/scim-users
description: Lists account members as SCIM User resources. Supports optional filtering by `userName` (email) using the SCIM filter syntax (e.g. `userName eq "user@example.com"`). Pagination is controlled via `startIndex` and `count` query parameters per RFC 7644 Section 3.4.2.4.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/scim/v2/Users
operation_ids:
    - scim-users-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List SCIM Users

`GET /accounts/{account_id}/scim/v2/Users`

Operation ID: `scim-users-list`

Lists account members as SCIM User resources. Supports optional filtering by `userName` (email) using the SCIM filter syntax (e.g. `userName eq "user@example.com"`). Pagination is controlled via `startIndex` and `count` query parameters per RFC 7644 Section 3.4.2.4.

## Definition

```yaml
{"operationId": "scim-users-list", "summary": "List SCIM Users", "description": "Lists account members as SCIM User resources. Supports optional filtering by `userName` (email) using the SCIM filter syntax (e.g. `userName eq \"user@example.com\"`). Pagination is controlled via `startIndex` and `count` query parameters per RFC 7644 Section 3.4.2.4.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "startIndex", "in": "query", "schema": {"description": "The 1-based index of the first result in the current set of list results. Values less than 1 are treated as 1.\n", "type": "integer", "default": 1, "minimum": 1}}, {"name": "count", "in": "query", "schema": {"description": "Specifies the desired maximum number of query results per page.\n", "type": "integer", "minimum": 0}}, {"name": "filter", "in": "query", "schema": {"description": "SCIM filter expression (RFC 7644 Section 3.4.2.2). Only `userName eq \"value\"` is supported.\n", "type": "string", "example": "userName eq \"user@example.com\""}}], "responses": {"200": {"description": "List SCIM Users response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_user_list_response"}}}}, "4XX": {"description": "List SCIM Users response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Users"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
