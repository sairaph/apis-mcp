---
title: Create User Subscription
page_id: operation-post-user-subscriptions-d906453a
path: operations/user-subscription
description: Creates a user subscription.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /user/subscriptions
operation_ids:
    - user-subscription-create-user-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create User Subscription

`POST /user/subscriptions`

Operation ID: `user-subscription-create-user-subscription`

Creates a user subscription.

## Definition

```yaml
{"operationId": "user-subscription-create-user-subscription", "summary": "Create User Subscription", "description": "Creates a user subscription.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_subscription-v2"}}}}, "responses": {"200": {"description": "Create User Subscription response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_user_subscription_response_single"}}}}, "4XX": {"description": "Create User Subscription response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_user_subscription_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Subscription"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:read", "#billing:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true, "x-forge-hidden": true}
```
