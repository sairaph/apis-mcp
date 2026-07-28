---
title: List user tenants
page_id: operation-get-user-tenants-b9cacc8e
path: operations/user
description: Retrieves list of tenants the authenticated user / method has access to.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/tenants
operation_ids:
    - User_listUserTenants
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List user tenants

`GET /user/tenants`

Operation ID: `User_listUserTenants`

Retrieves list of tenants the authenticated user / method has access to.

## Definition

```yaml
{"operationId": "User_listUserTenants", "summary": "List user tenants", "description": "Retrieves list of tenants the authenticated user / method has access to.", "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_Organization"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["User"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "user.tenants", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
