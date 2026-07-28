---
title: Create Subscriptions
page_id: operation-post-accounts-account-id-bulk-subscriptions-c0fe5157
path: operations/account-subscriptions
description: Creates multiple subscriptions for an account in a single request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/bulk/subscriptions
operation_ids:
    - account-subscriptions-bulk-create-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Subscriptions

`POST /accounts/{account_id}/bulk/subscriptions`

Operation ID: `account-subscriptions-bulk-create-subscription`

Creates multiple subscriptions for an account in a single request.

## Definition

```yaml
{"operationId": "account-subscriptions-bulk-create-subscription", "summary": "Create Subscriptions", "description": "Creates multiple subscriptions for an account in a single request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}, {"name": "idemp_key", "in": "query", "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"coupon_code": {"type": "string"}, "payment_hold_id": {"type": "integer", "format": "int64"}, "subscriptions": {"type": "array", "items": {"$ref": "#/components/schemas/bill-subs-api_subscription-v2"}}, "user_is_on_session": {"type": "boolean"}}}}}}, "responses": {"200": {"description": "Create Subscriptions response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_bulk_account_subscription_response"}}}}, "4XX": {"description": "Create Subscriptions response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_bulk_account_subscription_response"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Account Subscriptions"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:read", "#billing:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true, "x-forge-hidden": true}
```
