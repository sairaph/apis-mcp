---
title: Update Subscription
page_id: operation-put-accounts-account-id-subscriptions-subscription-identifier-d6f9af58
path: operations/account-subscriptions
description: Updates an account subscription.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/subscriptions/{subscription_identifier}
operation_ids:
    - account-subscriptions-update-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Subscription

`PUT /accounts/{account_id}/subscriptions/{subscription_identifier}`

Operation ID: `account-subscriptions-update-subscription`

Updates an account subscription.

## Definition

```yaml
{"operationId": "account-subscriptions-update-subscription", "summary": "Update Subscription", "description": "Updates an account subscription.", "parameters": [{"name": "subscription_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_subscription-v2"}}}}, "responses": {"200": {"description": "Update Subscription response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_account_subscription_response_single"}}}}, "4XX": {"description": "Update Subscription response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_account_subscription_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Subscriptions"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:read", "#billing:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.subscriptions", "x-fern-sdk-method-name": "update"}
```
