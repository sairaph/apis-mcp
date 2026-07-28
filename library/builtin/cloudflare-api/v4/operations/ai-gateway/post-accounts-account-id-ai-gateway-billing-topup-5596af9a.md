---
title: Create a top-up
page_id: operation-post-accounts-account-id-ai-gateway-billing-topup-af234705
path: operations/ai-gateway
description: Create a credit top-up via Stripe PaymentIntent for the given account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/topup
operation_ids:
    - aig-billing-create-topup
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a top-up

`POST /accounts/{account_id}/ai-gateway/billing/topup`

Operation ID: `aig-billing-create-topup`

Create a credit top-up via Stripe PaymentIntent for the given account.

## Definition

```yaml
{"operationId": "aig-billing-create-topup", "summary": "Create a top-up", "description": "Create a credit top-up via Stripe PaymentIntent for the given account.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"amount": {"description": "Top-up amount in cents (min 1000).", "type": "integer", "example": 5000, "minimum": 1000}}, "additionalProperties": false, "required": ["amount"]}}}}, "responses": {"200": {"description": "Top-up initiated successfully.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_CreateTopupResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing.topup", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
