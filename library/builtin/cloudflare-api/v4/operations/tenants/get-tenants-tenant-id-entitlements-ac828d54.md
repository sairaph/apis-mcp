---
title: List tenant entitlements
page_id: operation-get-tenants-tenant-id-entitlements-548eaef7
path: operations/tenants
description: List of innate entitlements available for the Tenant.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /tenants/{tenant_id}/entitlements
operation_ids:
    - Tenants_listEntitlements
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tenant entitlements

`GET /tenants/{tenant_id}/entitlements`

Operation ID: `Tenants_listEntitlements`

List of innate entitlements available for the Tenant.

## Definition

```yaml
{"operationId": "Tenants_listEntitlements", "summary": "List tenant entitlements", "description": "List of innate entitlements available for the Tenant.", "parameters": [{"name": "tenant_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_InnateEntitlements"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Tenants"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "tenants.entitlements", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
