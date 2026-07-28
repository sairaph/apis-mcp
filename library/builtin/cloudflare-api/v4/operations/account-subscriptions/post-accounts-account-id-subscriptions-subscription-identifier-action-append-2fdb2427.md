---
title: Append Subscription Action
page_id: operation-post-accounts-account-id-subscriptions-subscription-identifier-action-ap-97ad4359
path: operations/account-subscriptions
description: Smartly applies the incoming subscription into the lifecycle of the subscription.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/subscriptions/{subscription_identifier}/action/append
operation_ids:
    - account-subscriptions-action-append-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Append Subscription Action

`POST /accounts/{account_id}/subscriptions/{subscription_identifier}/action/append`

Operation ID: `account-subscriptions-action-append-subscription`

Smartly applies the incoming subscription into the lifecycle of the subscription.

## Definition

```yaml
{"operationId": "account-subscriptions-action-append-subscription", "summary": "Append Subscription Action", "description": "Smartly applies the incoming subscription into the lifecycle of the subscription.", "parameters": [{"name": "subscription_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_schemas-identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_subscription-v2"}}}}, "responses": {"200": {"description": "Append Subscription Action response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_account_subscription_response_single"}}}}, "4XX": {"description": "Append Subscription Action response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_account_subscription_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Subscriptions"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:read", "#billing:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true, "x-forge-hidden": true}
```
