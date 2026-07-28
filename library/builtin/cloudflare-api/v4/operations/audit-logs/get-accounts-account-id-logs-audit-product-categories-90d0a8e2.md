---
title: List account audit log product categories (Version 2)
page_id: operation-get-accounts-account-id-logs-audit-product-categories-6ce4b487
path: operations/audit-logs
description: Lists the available audit log product categories and the resource products each one expands to. Use these values with the product_category filter on the account audit logs endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/logs/audit/product_categories
operation_ids:
    - audit-logs-v2-list-account-product-categories
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account audit log product categories (Version 2)

`GET /accounts/{account_id}/logs/audit/product_categories`

Operation ID: `audit-logs-v2-list-account-product-categories`

Lists the available audit log product categories and the resource products each one expands to. Use these values with the product_category filter on the account audit logs endpoint.

## Definition

```yaml
{"operationId": "audit-logs-v2-list-account-product-categories", "summary": "List account audit log product categories (Version 2)", "description": "Lists the available audit log product categories and the resource products each one expands to. Use these values with the product_category filter on the account audit logs endpoint.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"description": "The unique id that identifies the account.", "type": "string", "example": "a67e14daa5f8dceeb91fe5449ba496ef"}}], "responses": {"200": {"description": "List account audit log product categories successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_audit-logs-v2-product-categories-response-collection"}}}}, "4XX": {"description": "List account audit log product categories failed response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-failure-2"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Audit Logs"], "x-api-token-group": ["Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
