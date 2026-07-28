---
title: Batch create organization members
page_id: operation-post-organizations-organization-id-members-batchcreate-22fc774e
path: operations/organizationmembers
description: Batch create multiple memberships that grant access to a specific Organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /organizations/{organization_id}/members:batchCreate
operation_ids:
    - Members_batchCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch create organization members

`POST /organizations/{organization_id}/members:batchCreate`

Operation ID: `Members_batchCreate`

Batch create multiple memberships that grant access to a specific Organization.

## Definition

```yaml
{"operationId": "Members_batchCreate", "summary": "Batch create organization members", "description": "Batch create multiple memberships that grant access to a specific Organization.", "parameters": [{"name": "organization_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/organizations-api_OrganizationID"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_BatchCreateMembersRequest"}}}}, "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_Member"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["OrganizationMembers"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations.members-batch-create", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
