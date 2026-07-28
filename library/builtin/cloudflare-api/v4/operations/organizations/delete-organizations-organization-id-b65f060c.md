---
title: Delete organization.
page_id: operation-delete-organizations-organization-id-55cd9e5f
path: operations/organizations
description: |-
    Delete an organization. The organization MUST be empty before deleting.
    It must not contain any sub-organizations, accounts, members or users. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

    **Access Control:** Restricted to enterprise organizations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /organizations/{organization_id}
operation_ids:
    - Organizations_delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete organization.

`DELETE /organizations/{organization_id}`

Operation ID: `Organizations_delete`

Delete an organization. The organization MUST be empty before deleting.
It must not contain any sub-organizations, accounts, members or users. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

**Access Control:** Restricted to enterprise organizations.

## Definition

```yaml
{"operationId": "Organizations_delete", "summary": "Delete organization.", "description": "Delete an organization. The organization MUST be empty before deleting.\nIt must not contain any sub-organizations, accounts, members or users. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)\n\n**Access Control:** Restricted to enterprise organizations.", "parameters": [{"name": "organization_id", "in": "path", "description": "The ID of the organization to delete.", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_DeleteOrganizationResponse"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Organizations"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
