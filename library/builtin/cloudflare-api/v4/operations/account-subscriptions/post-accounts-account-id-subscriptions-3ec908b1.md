---
title: Create Subscription
page_id: operation-post-accounts-account-id-subscriptions-ca693ab5
path: operations/account-subscriptions
description: Creates an account subscription.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/subscriptions
operation_ids:
    - account-subscriptions-create-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Subscription

`POST /accounts/{account_id}/subscriptions`

Operation ID: `account-subscriptions-create-subscription`

Creates an account subscription.

## Definition

```yaml
{"operationId": "account-subscriptions-create-subscription", "summary": "Create Subscription", "description": "Creates an account subscription.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_subscription-v2"}}}}, "responses": {"200": {"description": "Create Subscription response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_account_subscription_response_single"}}}}, "4XX": {"description": "Create Subscription response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_account_subscription_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Subscriptions"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:read", "#billing:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.subscriptions", "x-fern-sdk-method-name": "create"}
```
