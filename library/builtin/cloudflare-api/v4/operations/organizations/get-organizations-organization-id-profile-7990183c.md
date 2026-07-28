---
title: Get organization profile
page_id: operation-get-organizations-organization-id-profile-155b3310
path: operations/organizations
description: Get an organizations profile if it exists. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organizations/{organization_id}/profile
operation_ids:
    - Organizations_getProfile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get organization profile

`GET /organizations/{organization_id}/profile`

Operation ID: `Organizations_getProfile`

Get an organizations profile if it exists. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Organizations_getProfile", "summary": "Get organization profile", "description": "Get an organizations profile if it exists. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"name": "organization_id", "in": "path", "description": "The ID of the organization to retrieve a profile for.", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_ProfileResponse"}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Organizations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations.organization-profile", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
