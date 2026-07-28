---
title: List tenant accounts
page_id: operation-get-tenants-tenant-id-accounts-11fc3a59
path: operations/tenants
description: List of accounts for the Tenant.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /tenants/{tenant_id}/accounts
operation_ids:
    - Tenants_listAccounts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tenant accounts

`GET /tenants/{tenant_id}/accounts`

Operation ID: `Tenants_listAccounts`

List of accounts for the Tenant.

## Definition

```yaml
{"operationId": "Tenants_listAccounts", "summary": "List tenant accounts", "description": "List of accounts for the Tenant.", "parameters": [{"name": "tenant_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_Account"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Tenants"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "tenants.accounts", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
