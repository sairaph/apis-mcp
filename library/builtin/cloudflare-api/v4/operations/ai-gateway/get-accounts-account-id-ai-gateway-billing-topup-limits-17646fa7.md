---
title: Get account top-up limits
page_id: operation-get-accounts-account-id-ai-gateway-billing-topup-limits-4208e11b
path: operations/ai-gateway
description: Retrieve the minimum and maximum allowed top-up amounts (in cents) for this account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/topup/limits
operation_ids:
    - aig-billing-get-topup-limits
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account top-up limits

`GET /accounts/{account_id}/ai-gateway/billing/topup/limits`

Operation ID: `aig-billing-get-topup-limits`

Retrieve the minimum and maximum allowed top-up amounts (in cents) for this account.

## Definition

```yaml
{"operationId": "aig-billing-get-topup-limits", "summary": "Get account top-up limits", "description": "Retrieve the minimum and maximum allowed top-up amounts (in cents) for this account.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Top-up limits retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_GetTopupLimitsResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
