---
title: List Subscriptions
page_id: operation-get-accounts-account-id-subscriptions-d5c8b68e
path: operations/account-subscriptions
description: Lists all of an account's subscriptions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/subscriptions
operation_ids:
    - account-subscriptions-list-subscriptions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Subscriptions

`GET /accounts/{account_id}/subscriptions`

Operation ID: `account-subscriptions-list-subscriptions`

Lists all of an account's subscriptions.

## Definition

```yaml
{"operationId": "account-subscriptions-list-subscriptions", "summary": "List Subscriptions", "description": "Lists all of an account's subscriptions.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "responses": {"200": {"description": "List Subscriptions response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_account_subscription_response_collection"}}}}, "4XX": {"description": "List Subscriptions response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_account_subscription_response_collection"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Subscriptions"], "x-api-token-group": ["Billing Write", "Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.subscriptions", "x-fern-sdk-method-name": "get"}
```
