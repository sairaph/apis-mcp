---
title: Modify organization profile.
page_id: operation-put-organizations-organization-id-profile-64ff590a
path: operations/organizations
description: Modify organization profile. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /organizations/{organization_id}/profile
operation_ids:
    - Organizations_modifyProfile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Modify organization profile.

`PUT /organizations/{organization_id}/profile`

Operation ID: `Organizations_modifyProfile`

Modify organization profile. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Organizations_modifyProfile", "summary": "Modify organization profile.", "description": "Modify organization profile. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"name": "organization_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_Profile"}}}}, "responses": {"204": {"description": "There is no content to send for this request, but the headers may be useful."}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Organizations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations.organization-profile", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
