---
title: List SCIM Schemas
page_id: operation-get-accounts-account-id-scim-v2-schemas-b412d1e3
path: operations/scim-discovery
description: Returns the list of SCIM schemas supported by the Cloudflare SCIM service (RFC 7643 Section 7, RFC 7644 Section 4). Clients use this to introspect the attributes of each resource type. Query parameters are not supported on this endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/scim/v2/Schemas
operation_ids:
    - scim-schemas-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List SCIM Schemas

`GET /accounts/{account_id}/scim/v2/Schemas`

Operation ID: `scim-schemas-list`

Returns the list of SCIM schemas supported by the Cloudflare SCIM service (RFC 7643 Section 7, RFC 7644 Section 4). Clients use this to introspect the attributes of each resource type. Query parameters are not supported on this endpoint.

## Definition

```yaml
{"operationId": "scim-schemas-list", "summary": "List SCIM Schemas", "description": "Returns the list of SCIM schemas supported by the Cloudflare SCIM service (RFC 7643 Section 7, RFC 7644 Section 4). Clients use this to introspect the attributes of each resource type. Query parameters are not supported on this endpoint.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "responses": {"200": {"description": "List SCIM Schemas response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_schema_list_response"}}}}, "4XX": {"description": "List SCIM Schemas response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Discovery"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
