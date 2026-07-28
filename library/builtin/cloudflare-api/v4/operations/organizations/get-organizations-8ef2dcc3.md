---
title: List organizations the user has access to
page_id: operation-get-organizations-44c62ccf
path: operations/organizations
description: Retrieve a list of organizations a particular user has access to. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organizations
operation_ids:
    - Organization_listOrganizations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List organizations the user has access to

`GET /organizations`

Operation ID: `Organization_listOrganizations`

Retrieve a list of organizations a particular user has access to. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Organization_listOrganizations", "summary": "List organizations the user has access to", "description": "Retrieve a list of organizations a particular user has access to. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"$ref": "#/components/parameters/organizations-api_OrganizationListFiltersId"}, {"$ref": "#/components/parameters/organizations-api_OrganizationListFiltersName"}, {"$ref": "#/components/parameters/organizations-api_OrganizationListFiltersNameStartsWith"}, {"$ref": "#/components/parameters/organizations-api_OrganizationListFiltersNameEndsWith"}, {"$ref": "#/components/parameters/organizations-api_OrganizationListFiltersNameContains"}, {"$ref": "#/components/parameters/organizations-api_OrganizationListFiltersContainingAccount"}, {"$ref": "#/components/parameters/organizations-api_OrganizationListFiltersContainingUser"}, {"$ref": "#/components/parameters/organizations-api_OrganizationListFiltersContainingOrganization"}, {"$ref": "#/components/parameters/organizations-api_OrganizationListFiltersParentId"}, {"$ref": "#/components/parameters/organizations-api_PageTokenParamsPageToken"}, {"$ref": "#/components/parameters/organizations-api_PageTokenParamsPageSize"}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_Organization"}}, "result_info": {"$ref": "#/components/schemas/organizations-api_PageTokenResultInfo"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result", "result_info"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Organizations"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
