---
title: Get tenant
page_id: operation-get-tenants-tenant-id-b2a3f36b
path: operations/tenants
description: Retrieves a Tenant by Tenant ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /tenants/{tenant_id}
operation_ids:
    - Tenants_retrieveTenant
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get tenant

`GET /tenants/{tenant_id}`

Operation ID: `Tenants_retrieveTenant`

Retrieves a Tenant by Tenant ID.

## Definition

```yaml
{"operationId": "Tenants_retrieveTenant", "summary": "Get tenant", "description": "Retrieves a Tenant by Tenant ID.", "parameters": [{"name": "tenant_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_Tenant"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Tenants"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "tenants", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
