---
title: Get tenant account types
page_id: operation-get-tenants-tenant-id-account-types-ffdbcb8c
path: operations/tenants
description: List of account types available for the Tenant to provision accounts.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /tenants/{tenant_id}/account_types
operation_ids:
    - Tenants_validAccountTypes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get tenant account types

`GET /tenants/{tenant_id}/account_types`

Operation ID: `Tenants_validAccountTypes`

List of account types available for the Tenant to provision accounts.

## Definition

```yaml
{"operationId": "Tenants_validAccountTypes", "summary": "Get tenant account types", "description": "List of account types available for the Tenant to provision accounts.", "parameters": [{"name": "tenant_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"type": "array", "items": {"type": "string"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Tenants"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "tenants.account-types", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
