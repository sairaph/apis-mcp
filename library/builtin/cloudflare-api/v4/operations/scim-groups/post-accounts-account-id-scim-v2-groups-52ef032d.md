---
title: Create SCIM Group
page_id: operation-post-accounts-account-id-scim-v2-groups-2a9e2ab0
path: operations/scim-groups
description: Creates a new SCIM Group (user group) for the account. The `displayName` must not be empty and must not begin with `CF` (reserved for system groups).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/scim/v2/Groups
operation_ids:
    - scim-groups-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create SCIM Group

`POST /accounts/{account_id}/scim/v2/Groups`

Operation ID: `scim-groups-create`

Creates a new SCIM Group (user group) for the account. The `displayName` must not be empty and must not begin with `CF` (reserved for system groups).

## Definition

```yaml
{"operationId": "scim-groups-create", "summary": "Create SCIM Group", "description": "Creates a new SCIM Group (user group) for the account. The `displayName` must not be empty and must not begin with `CF` (reserved for system groups).\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_group_create_request"}}}}, "responses": {"201": {"description": "Create SCIM Group response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_group"}}}}, "4XX": {"description": "Create SCIM Group response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Groups"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.create"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
