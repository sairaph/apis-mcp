---
title: Get usage history
page_id: operation-get-accounts-account-id-ai-gateway-billing-usage-history-c1d00b7d
path: operations/ai-gateway
description: Retrieve aggregated usage meter event summaries for the given time range.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/ai-gateway/billing/usage-history
operation_ids:
    - aig-billing-get-usage-history
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get usage history

`GET /accounts/{account_id}/ai-gateway/billing/usage-history`

Operation ID: `aig-billing-get-usage-history`

Retrieve aggregated usage meter event summaries for the given time range.

## Definition

```yaml
{"operationId": "aig-billing-get-usage-history", "summary": "Get usage history", "description": "Retrieve aggregated usage meter event summaries for the given time range.", "parameters": [{"name": "account_id", "in": "path", "description": "Cloudflare account ID.", "required": true, "schema": {"type": "string"}}, {"name": "value_grouping_window", "in": "query", "description": "Grouping window for usage data.", "required": true, "schema": {"description": "Grouping window for usage data.", "type": "string", "example": "day", "enum": ["day", "hour"]}}, {"name": "start_time", "in": "query", "description": "Start time as Unix timestamp in milliseconds.", "schema": {"description": "Start time as Unix timestamp in milliseconds.", "type": "number", "example": 1700000000000, "nullable": true}}, {"name": "end_time", "in": "query", "description": "End time as Unix timestamp in milliseconds.", "schema": {"description": "End time as Unix timestamp in milliseconds.", "type": "number", "example": 1700086400000, "nullable": true}}], "responses": {"200": {"description": "Usage history retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_GetUsageHistoryResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aig-billing_ErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["AI Gateway"], "x-api-token-group": ["AI Gateway Write", "AI Gateway Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.aig"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "ai_gateway.billing", "x-fern-sdk-method-name": "usage_history", "x-forge-hidden": true}
```
