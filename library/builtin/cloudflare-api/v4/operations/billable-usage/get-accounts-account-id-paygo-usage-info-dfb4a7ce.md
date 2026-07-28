---
title: Get PayGo Account Billable Usage Info (Version 1, Alpha)
page_id: operation-get-accounts-account-id-paygo-usage-info-6f6c70af
path: operations/billable-usage
description: |-
    Returns high-level usage information for the account, including coverage,
    and subscription metadata.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/paygo-usage-info
operation_ids:
    - billable-usage-get-paygo-account-usage-info
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get PayGo Account Billable Usage Info (Version 1, Alpha)

`GET /accounts/{account_id}/paygo-usage-info`

Operation ID: `billable-usage-get-paygo-account-usage-info`

Returns high-level usage information for the account, including coverage,
and subscription metadata.

## Definition

```yaml
{"operationId": "billable-usage-get-paygo-account-usage-info", "summary": "Get PayGo Account Billable Usage Info (Version 1, Alpha)", "description": "Returns high-level usage information for the account, including coverage,\nand subscription metadata.\n", "parameters": [{"$ref": "#/components/parameters/billable-usage-api_account_id"}], "responses": {"200": {"description": "Indicates PayGo account usage data was successfully retrieved.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billable-usage-api_usage_info_response"}}}}, "4XX": {"description": "Indicates the request failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/billable-usage-api_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Billable Usage"], "x-api-token-group": ["Billing Read"], "x-cfPlanAvailability": {"business": true, "enterprise": false, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "billing.usage", "x-fern-sdk-method-name": "paygo_info", "x-forge-hidden": true}
```
