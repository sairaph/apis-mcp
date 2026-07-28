---
title: List SCIM Resource Types
page_id: operation-get-accounts-account-id-scim-v2-resourcetypes-ed64430b
path: operations/scim-discovery
description: Returns the list of SCIM resource types supported by the Cloudflare SCIM service (RFC 7643 Section 6, RFC 7644 Section 4). Clients use this to discover available resource categories (e.g. Users, Groups) and their associated schemas. Query parameters are not supported on this endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/scim/v2/ResourceTypes
operation_ids:
    - scim-resource-types-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List SCIM Resource Types

`GET /accounts/{account_id}/scim/v2/ResourceTypes`

Operation ID: `scim-resource-types-list`

Returns the list of SCIM resource types supported by the Cloudflare SCIM service (RFC 7643 Section 6, RFC 7644 Section 4). Clients use this to discover available resource categories (e.g. Users, Groups) and their associated schemas. Query parameters are not supported on this endpoint.

## Definition

```yaml
{"operationId": "scim-resource-types-list", "summary": "List SCIM Resource Types", "description": "Returns the list of SCIM resource types supported by the Cloudflare SCIM service (RFC 7643 Section 6, RFC 7644 Section 4). Clients use this to discover available resource categories (e.g. Users, Groups) and their associated schemas. Query parameters are not supported on this endpoint.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "responses": {"200": {"description": "List SCIM Resource Types response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_resource_type_list_response"}}}}, "4XX": {"description": "List SCIM Resource Types response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Discovery"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
