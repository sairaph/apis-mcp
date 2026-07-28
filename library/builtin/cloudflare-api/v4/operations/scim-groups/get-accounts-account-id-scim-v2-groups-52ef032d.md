---
title: List SCIM Groups
page_id: operation-get-accounts-account-id-scim-v2-groups-5d1488ce
path: operations/scim-groups
description: Lists SCIM Group resources for the account. Returns both system groups (backed by Cloudflare permission groups, prefixed `cloudflare-v1-`) and custom user groups. Supports filtering by `displayName` using SCIM filter syntax.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/scim/v2/Groups
operation_ids:
    - scim-groups-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List SCIM Groups

`GET /accounts/{account_id}/scim/v2/Groups`

Operation ID: `scim-groups-list`

Lists SCIM Group resources for the account. Returns both system groups (backed by Cloudflare permission groups, prefixed `cloudflare-v1-`) and custom user groups. Supports filtering by `displayName` using SCIM filter syntax.

## Definition

```yaml
{"operationId": "scim-groups-list", "summary": "List SCIM Groups", "description": "Lists SCIM Group resources for the account. Returns both system groups (backed by Cloudflare permission groups, prefixed `cloudflare-v1-`) and custom user groups. Supports filtering by `displayName` using SCIM filter syntax.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "startIndex", "in": "query", "schema": {"description": "The 1-based index of the first result in the current set of list results. Values less than 1 are treated as 1.\n", "type": "integer", "default": 1, "minimum": 1}}, {"name": "count", "in": "query", "schema": {"description": "Specifies the desired maximum number of query results per page.\n", "type": "integer", "minimum": 0}}, {"name": "filter", "in": "query", "schema": {"description": "SCIM filter expression (RFC 7644 Section 3.4.2.2). Only `displayName eq \"value\"` is supported.\n", "type": "string", "example": "displayName eq \"My Group\""}}], "responses": {"200": {"description": "List SCIM Groups response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_group_list_response"}}}}, "4XX": {"description": "List SCIM Groups response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
