---
title: Update User Subscription
page_id: operation-put-user-subscriptions-identifier-b2d680ee
path: operations/user-subscription
description: Updates a user's subscriptions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /user/subscriptions/{identifier}
operation_ids:
    - user-subscription-update-user-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update User Subscription

`PUT /user/subscriptions/{identifier}`

Operation ID: `user-subscription-update-user-subscription`

Updates a user's subscriptions.

## Definition

```yaml
{"operationId": "user-subscription-update-user-subscription", "summary": "Update User Subscription", "description": "Updates a user's subscriptions.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_subscription-v2"}}}}, "responses": {"200": {"description": "Update User Subscription response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_user_subscription_response_single"}}}}, "4XX": {"description": "Update User Subscription response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_user_subscription_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Subscription"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:read", "#billing:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "user.subscriptions", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
