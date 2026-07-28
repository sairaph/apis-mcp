---
title: Get PayGo Account Billable Usage (Version 1, Alpha)
page_id: operation-get-accounts-account-id-paygo-usage-60ede788
path: operations/billable-usage
description: |-
    Returns billable usage data for PayGo (self-serve) accounts.
    When no query parameters are provided, returns usage for the current
    billing period.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/paygo-usage
operation_ids:
    - billable-usage-get-paygo-account-usage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get PayGo Account Billable Usage (Version 1, Alpha)

`GET /accounts/{account_id}/paygo-usage`

Operation ID: `billable-usage-get-paygo-account-usage`

Returns billable usage data for PayGo (self-serve) accounts.
When no query parameters are provided, returns usage for the current
billing period.

## Definition

```yaml
{"operationId": "billable-usage-get-paygo-account-usage", "summary": "Get PayGo Account Billable Usage (Version 1, Alpha)", "description": "Returns billable usage data for PayGo (self-serve) accounts.\nWhen no query parameters are provided, returns usage for the current\nbilling period.\n", "parameters": [{"$ref": "#/components/parameters/billable-usage-api_account_id"}, {"$ref": "#/components/parameters/billable-usage-api_from"}, {"$ref": "#/components/parameters/billable-usage-api_to"}], "responses": {"200": {"description": "Indicates PayGo account usage data was successfully retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billable-usage-api_usage_response"}}}}, "4XX": {"description": "Indicates the request failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billable-usage-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Billable Usage"], "x-api-token-group": ["Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": false, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "billing.usage", "x-fern-sdk-method-name": "paygo", "x-forge-hidden": true}
```
