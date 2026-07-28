---
title: Delete spending limit
page_id: operation-delete-accounts-account-id-ai-gateway-billing-spending-limit-a756a905
path: operations/ai-gateway
description: Remove the spending limit for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/spending-limit
operation_ids:
    - aig-billing-delete-spending-limit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete spending limit

`DELETE /accounts/{account_id}/ai-gateway/billing/spending-limit`

Operation ID: `aig-billing-delete-spending-limit`

Remove the spending limit for the account.

## Definition

```yaml
{"operationId": "aig-billing-delete-spending-limit", "summary": "Delete spending limit", "description": "Remove the spending limit for the account.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Spending limit deleted.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_DeleteSpendingLimitResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing.spending_limit", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
