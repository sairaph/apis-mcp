---
title: Get organization member
page_id: operation-get-organizations-organization-id-members-member-id-87cbcb18
path: operations/organizationmembers
description: Retrieve a single membership from an Organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /organizations/{organization_id}/members/{member_id}
operation_ids:
    - Members_retrieve
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get organization member

`GET /organizations/{organization_id}/members/{member_id}`

Operation ID: `Members_retrieve`

Retrieve a single membership from an Organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Members_retrieve", "summary": "Get organization member", "description": "Retrieve a single membership from an Organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"name": "organization_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}, {"name": "member_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_MemberID"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_Member"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["OrganizationMembers"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations.members", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
