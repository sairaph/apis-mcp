---
title: List organization members
page_id: operation-get-organizations-organization-id-members-7ea2fdef
path: operations/organizationmembers
description: List memberships for an Organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organizations/{organization_id}/members
operation_ids:
    - Members_list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List organization members

`GET /organizations/{organization_id}/members`

Operation ID: `Members_list`

List memberships for an Organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Members_list", "summary": "List organization members", "description": "List memberships for an Organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"name": "organization_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}, {"name": "status", "in": "query", "description": "Filter the list of memberships by membership status.", "schema": {"type": "array", "items": {"enum": ["active", "pending", "rejected", "canceled"], "type": "string"}}}, {"name": "user.email", "in": "query", "description": "Filter the list of memberships for a specific email.", "schema": {"type": "string"}, "explode": false}, {"name": "user.email.contains", "in": "query", "description": "Filter the list of memberships for a specific email that contains a substring.", "schema": {"type": "string"}, "explode": false}, {"name": "user.email.startsWith", "in": "query", "description": "Filter the list of memberships for a specific email that starts with a substring.", "schema": {"type": "string"}, "explode": false}, {"name": "user.email.endsWith", "in": "query", "description": "Filter the list of memberships for a specific email that ends with a substring.", "schema": {"type": "string"}, "explode": false}, {"$ref": "#/components/parameters/organizations-api_PageTokenParamsPageToken"}, {"$ref": "#/components/parameters/organizations-api_PageTokenParamsPageSize"}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_Member"}}, "result_info": {"$ref": "#/components/schemas/organizations-api_PageTokenResultInfo"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result", "result_info"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["OrganizationMembers"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations.members", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
