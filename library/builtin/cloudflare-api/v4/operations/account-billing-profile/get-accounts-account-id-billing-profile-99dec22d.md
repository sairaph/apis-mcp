---
title: Billing Profile Details
page_id: operation-get-accounts-account-id-billing-profile-f1708805
path: operations/account-billing-profile
description: Gets the current billing profile for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/billing/profile
operation_ids:
    - account-billing-profile-(-deprecated)-billing-profile-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Billing Profile Details

`GET /accounts/{account_id}/billing/profile`

Operation ID: `account-billing-profile-(-deprecated)-billing-profile-details`

Gets the current billing profile for the account.

## Definition

```yaml
{"operationId": "account-billing-profile-(-deprecated)-billing-profile-details", "summary": "Billing Profile Details", "description": "Gets the current billing profile for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "responses": {"200": {"description": "Billing Profile Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_billing_response_single"}}}}, "4XX": {"description": "Billing Profile Details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_billing_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Billing Profile"], "x-api-token-group": ["Billing Write", "Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "billing.profiles", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
