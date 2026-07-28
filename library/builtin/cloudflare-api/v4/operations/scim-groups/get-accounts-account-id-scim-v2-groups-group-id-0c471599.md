---
title: Get SCIM Group
page_id: operation-get-accounts-account-id-scim-v2-groups-group-id-3cccc9af
path: operations/scim-groups
description: Retrieves a single SCIM Group resource by group ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/scim/v2/Groups/{group_id}
operation_ids:
    - scim-groups-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get SCIM Group

`GET /accounts/{account_id}/scim/v2/Groups/{group_id}`

Operation ID: `scim-groups-get`

Retrieves a single SCIM Group resource by group ID.

## Definition

```yaml
{"operationId": "scim-groups-get", "summary": "Get SCIM Group", "description": "Retrieves a single SCIM Group resource by group ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_scim_group_identifier"}}], "responses": {"200": {"description": "Get SCIM Group response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_group"}}}}, "4XX": {"description": "Get SCIM Group response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
