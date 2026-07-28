---
title: Check top-up status
page_id: operation-post-accounts-account-id-ai-gateway-billing-topup-status-c572cbd6
path: operations/ai-gateway
description: Get the payment processing status of a top-up by its invoice ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/topup/status
operation_ids:
    - aig-billing-check-topup-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Check top-up status

`POST /accounts/{account_id}/ai-gateway/billing/topup/status`

Operation ID: `aig-billing-check-topup-status`

Get the payment processing status of a top-up by its invoice ID.

## Definition

```yaml
{"operationId": "aig-billing-check-topup-status", "summary": "Check top-up status", "description": "Get the payment processing status of a top-up by its invoice ID.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"payment_intent_id": {"description": "Stripe invoice ID to check status for.", "type": "string", "example": "in_1abc"}}, "additionalProperties": false, "required": ["payment_intent_id"]}}}}, "responses": {"200": {"description": "Top-up status retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_TopupStatusResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing.topup", "x-fern-sdk-method-name": "status", "x-forge-hidden": true}
```
