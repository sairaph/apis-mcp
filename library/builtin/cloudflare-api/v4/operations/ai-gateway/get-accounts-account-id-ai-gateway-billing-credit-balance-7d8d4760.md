---
title: Get credit balance
page_id: operation-get-accounts-account-id-ai-gateway-billing-credit-balance-ba58dcdb
path: operations/ai-gateway
description: Retrieve the current credit balance, payment method info, and top-up configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/credit-balance
operation_ids:
    - aig-billing-get-credit-balance
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get credit balance

`GET /accounts/{account_id}/ai-gateway/billing/credit-balance`

Operation ID: `aig-billing-get-credit-balance`

Retrieve the current credit balance, payment method info, and top-up configuration.

## Definition

```yaml
{"operationId": "aig-billing-get-credit-balance", "summary": "Get credit balance", "description": "Retrieve the current credit balance, payment method info, and top-up configuration.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Credit balance retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_GetCreditBalanceResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing", "x-fern-sdk-method-name": "credit_balance", "x-forge-hidden": true}
```
