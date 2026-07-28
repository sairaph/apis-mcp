---
title: Set spending limit (deprecated)
page_id: operation-post-accounts-account-id-ai-gateway-billing-spending-limit-a1633819
path: operations/ai-gateway
description: 'Deprecated: spending limits can no longer be created, enabled, or modified and this endpoint always responds 403. Use the new AI Gateway spend limits instead: https://developers.cloudflare.com/ai-gateway/features/spend-limits/. Existing limits can be removed via DELETE /spending-limit.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/spending-limit
operation_ids:
    - aig-billing-set-spending-limit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set spending limit (deprecated)

`POST /accounts/{account_id}/ai-gateway/billing/spending-limit`

Operation ID: `aig-billing-set-spending-limit`

Deprecated: spending limits can no longer be created, enabled, or modified and this endpoint always responds 403. Use the new AI Gateway spend limits instead: https://developers.cloudflare.com/ai-gateway/features/spend-limits/. Existing limits can be removed via DELETE /spending-limit.

## Definition

```yaml
{"operationId": "aig-billing-set-spending-limit", "summary": "Set spending limit (deprecated)", "description": "Deprecated: spending limits can no longer be created, enabled, or modified and this endpoint always responds 403. Use the new AI Gateway spend limits instead: https://developers.cloudflare.com/ai-gateway/features/spend-limits/. Existing limits can be removed via DELETE /spending-limit.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"amount": {"description": "Spending limit amount in cents (min 100).", "type": "integer", "example": 10000, "minimum": 100}, "duration": {"description": "Spending limit duration.", "type": "string", "example": "monthly", "enum": ["daily", "weekly", "monthly"]}, "strategy": {"description": "Spending limit strategy.", "type": "string", "example": "fixed", "enum": ["fixed", "sliding"]}}, "additionalProperties": false, "required": ["amount", "strategy", "duration"]}}}}, "responses": {"201": {"description": "Spending limit created.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_SetSpendingLimitResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing.spending_limit", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
