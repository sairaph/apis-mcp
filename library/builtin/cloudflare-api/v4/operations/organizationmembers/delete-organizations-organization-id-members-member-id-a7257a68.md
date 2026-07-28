---
title: Delete organization member
page_id: operation-delete-organizations-organization-id-members-member-id-6b2124fa
path: operations/organizationmembers
description: Delete a membership to a particular Organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /organizations/{organization_id}/members/{member_id}
operation_ids:
    - Members_delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete organization member

`DELETE /organizations/{organization_id}/members/{member_id}`

Operation ID: `Members_delete`

Delete a membership to a particular Organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Members_delete", "summary": "Delete organization member", "description": "Delete a membership to a particular Organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"name": "organization_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}, {"name": "member_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_MemberID"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"member_id": {"$ref": "#/components/schemas/organizations-api_MemberID"}}, "required": ["member_id"]}}}}, "responses": {"204": {"description": "There is no content to send for this request, but the headers may be useful."}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["OrganizationMembers"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations.members", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
