---
title: Get organization
page_id: operation-get-organizations-organization-id-4fb1ab4b
path: operations/organizations
description: Retrieve the details of a certain organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organizations/{organization_id}
operation_ids:
    - Organizations_retrieve
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get organization

`GET /organizations/{organization_id}`

Operation ID: `Organizations_retrieve`

Retrieve the details of a certain organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Organizations_retrieve", "summary": "Get organization", "description": "Retrieve the details of a certain organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"name": "organization_id", "in": "path", "description": "The ID of the organization to retrieve.", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_Organization"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Organizations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
