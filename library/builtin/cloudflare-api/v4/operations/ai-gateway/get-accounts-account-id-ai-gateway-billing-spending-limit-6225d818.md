---
title: Get spending limit
page_id: operation-get-accounts-account-id-ai-gateway-billing-spending-limit-c8aaca23
path: operations/ai-gateway
description: Retrieve the current spending limit configuration for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/spending-limit
operation_ids:
    - aig-billing-get-spending-limit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get spending limit

`GET /accounts/{account_id}/ai-gateway/billing/spending-limit`

Operation ID: `aig-billing-get-spending-limit`

Retrieve the current spending limit configuration for the account.

## Definition

```yaml
{"operationId": "aig-billing-get-spending-limit", "summary": "Get spending limit", "description": "Retrieve the current spending limit configuration for the account.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Spending limit configuration retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_GetSpendingLimitResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing.spending_limit", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
