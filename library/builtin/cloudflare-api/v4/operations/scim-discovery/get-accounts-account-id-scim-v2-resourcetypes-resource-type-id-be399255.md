---
title: Get SCIM Resource Type
page_id: operation-get-accounts-account-id-scim-v2-resourcetypes-resource-type-id-132cf26d
path: operations/scim-discovery
description: Returns a single SCIM resource type by ID (RFC 7643 Section 6). Valid IDs are `User` and `Group`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/scim/v2/ResourceTypes/{resource_type_id}
operation_ids:
    - scim-resource-types-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get SCIM Resource Type

`GET /accounts/{account_id}/scim/v2/ResourceTypes/{resource_type_id}`

Operation ID: `scim-resource-types-get`

Returns a single SCIM resource type by ID (RFC 7643 Section 6). Valid IDs are `User` and `Group`.

## Definition

```yaml
{"operationId": "scim-resource-types-get", "summary": "Get SCIM Resource Type", "description": "Returns a single SCIM resource type by ID (RFC 7643 Section 6). Valid IDs are `User` and `Group`.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "resource_type_id", "in": "path", "required": true, "schema": {"description": "The resource type identifier.", "type": "string", "example": "User", "enum": ["User", "Group"]}}], "responses": {"200": {"description": "Get SCIM Resource Type response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_resource_type"}}}}, "4XX": {"description": "Get SCIM Resource Type response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Discovery"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
