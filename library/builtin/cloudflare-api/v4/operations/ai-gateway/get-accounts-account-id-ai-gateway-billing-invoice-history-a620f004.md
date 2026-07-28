---
title: Get invoice history
page_id: operation-get-accounts-account-id-ai-gateway-billing-invoice-history-03398134
path: operations/ai-gateway
description: Retrieve a list of past invoices with pagination, optionally filtered by type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/invoice-history
operation_ids:
    - aig-billing-get-invoice-history
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get invoice history

`GET /accounts/{account_id}/ai-gateway/billing/invoice-history`

Operation ID: `aig-billing-get-invoice-history`

Retrieve a list of past invoices with pagination, optionally filtered by type.

## Definition

```yaml
{"operationId": "aig-billing-get-invoice-history", "summary": "Get invoice history", "description": "Retrieve a list of past invoices with pagination, optionally filtered by type.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}, {"name": "type", "in": "query", "description": "Filter invoice type: auto, manual, or all.", "schema": {"description": "Filter invoice type: auto, manual, or all.", "type": "string", "example": "all", "default": "all", "enum": ["auto", "all", "manual"]}}], "responses": {"200": {"description": "Invoice history retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_GetInvoiceHistoryResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing", "x-fern-sdk-method-name": "invoice_history", "x-forge-hidden": true}
```
