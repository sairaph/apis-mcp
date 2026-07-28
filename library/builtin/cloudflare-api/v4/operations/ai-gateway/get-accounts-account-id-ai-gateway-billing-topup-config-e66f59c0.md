---
title: Get auto top-up configuration
page_id: operation-get-accounts-account-id-ai-gateway-billing-topup-config-a7f2d4a2
path: operations/ai-gateway
description: Retrieve the current auto top-up threshold, amount, and any error state.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/topup/config
operation_ids:
    - aig-billing-get-topup-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get auto top-up configuration

`GET /accounts/{account_id}/ai-gateway/billing/topup/config`

Operation ID: `aig-billing-get-topup-config`

Retrieve the current auto top-up threshold, amount, and any error state.

## Definition

```yaml
{"operationId": "aig-billing-get-topup-config", "summary": "Get auto top-up configuration", "description": "Retrieve the current auto top-up threshold, amount, and any error state.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Auto top-up configuration retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_GetTopupConfigResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing.topup.config", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
