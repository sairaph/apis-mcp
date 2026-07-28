---
title: Get SCIM Schema
page_id: operation-get-accounts-account-id-scim-v2-schemas-schema-id-5db5efa0
path: operations/scim-discovery
description: Returns a single SCIM schema definition by schema URI ID (RFC 7643 Section 7). Valid IDs are `urn:ietf:params:scim:schemas:core:2.0:User` and `urn:ietf:params:scim:schemas:core:2.0:Group`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/scim/v2/Schemas/{schema_id}
operation_ids:
    - scim-schemas-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get SCIM Schema

`GET /accounts/{account_id}/scim/v2/Schemas/{schema_id}`

Operation ID: `scim-schemas-get`

Returns a single SCIM schema definition by schema URI ID (RFC 7643 Section 7). Valid IDs are `urn:ietf:params:scim:schemas:core:2.0:User` and `urn:ietf:params:scim:schemas:core:2.0:Group`.

## Definition

```yaml
{"operationId": "scim-schemas-get", "summary": "Get SCIM Schema", "description": "Returns a single SCIM schema definition by schema URI ID (RFC 7643 Section 7). Valid IDs are `urn:ietf:params:scim:schemas:core:2.0:User` and `urn:ietf:params:scim:schemas:core:2.0:Group`.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "schema_id", "in": "path", "required": true, "schema": {"description": "The schema URI identifier.", "type": "string", "example": "urn:ietf:params:scim:schemas:core:2.0:User"}}], "responses": {"200": {"description": "Get SCIM Schema response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_schema"}}}}, "4XX": {"description": "Get SCIM Schema response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Discovery"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
