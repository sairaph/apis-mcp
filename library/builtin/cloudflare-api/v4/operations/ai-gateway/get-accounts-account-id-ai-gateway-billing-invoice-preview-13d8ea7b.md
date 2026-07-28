---
title: Get invoice preview
page_id: operation-get-accounts-account-id-ai-gateway-billing-invoice-preview-d2e367a4
path: operations/ai-gateway
description: Retrieve a preview of the upcoming invoice including line items and tax.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/invoice-preview
operation_ids:
    - aig-billing-get-invoice-preview
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get invoice preview

`GET /accounts/{account_id}/ai-gateway/billing/invoice-preview`

Operation ID: `aig-billing-get-invoice-preview`

Retrieve a preview of the upcoming invoice including line items and tax.

## Definition

```yaml
{"operationId": "aig-billing-get-invoice-preview", "summary": "Get invoice preview", "description": "Retrieve a preview of the upcoming invoice including line items and tax.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Invoice preview retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_GetInvoicePreviewResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing", "x-fern-sdk-method-name": "invoice_preview", "x-forge-hidden": true}
```
