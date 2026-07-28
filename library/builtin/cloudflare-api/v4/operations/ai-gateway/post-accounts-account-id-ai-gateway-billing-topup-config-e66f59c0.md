---
title: Set auto top-up configuration
page_id: operation-post-accounts-account-id-ai-gateway-billing-topup-config-d10b0eaf
path: operations/ai-gateway
description: Configure auto top-up with a balance threshold and top-up amount.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/topup/config
operation_ids:
    - aig-billing-set-topup-config
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set auto top-up configuration

`POST /accounts/{account_id}/ai-gateway/billing/topup/config`

Operation ID: `aig-billing-set-topup-config`

Configure auto top-up with a balance threshold and top-up amount.

## Definition

```yaml
{"operationId": "aig-billing-set-topup-config", "summary": "Set auto top-up configuration", "description": "Configure auto top-up with a balance threshold and top-up amount.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"amount": {"description": "Auto top-up amount in cents (min 1000).", "type": "integer", "example": 5000, "minimum": 1000}, "threshold": {"description": "Balance threshold in cents that triggers auto top-up (min 500).", "type": "integer", "example": 500, "minimum": 500}}, "additionalProperties": false, "required": ["threshold", "amount"]}}}}, "responses": {"200": {"description": "Auto top-up configuration saved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_SetTopupConfigResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing.topup.config", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
