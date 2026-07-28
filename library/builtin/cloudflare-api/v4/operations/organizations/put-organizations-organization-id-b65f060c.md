---
title: Modify organization.
page_id: operation-put-organizations-organization-id-8ea70697
path: operations/organizations
description: Modify organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /organizations/{organization_id}
operation_ids:
    - Organizations_modify
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Modify organization.

`PUT /organizations/{organization_id}`

Operation ID: `Organizations_modify`

Modify organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Organizations_modify", "summary": "Modify organization.", "description": "Modify organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"name": "organization_id", "in": "path", "description": "The ID of the organization to modify.", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}], "requestBody": {"description": "The new details of the organization. Only accepts updates\nto \"name\" currently.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_Organization"}}}}, "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_Organization"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Organizations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
