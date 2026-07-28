---
title: Get User Subscriptions
page_id: operation-get-user-subscriptions-45796097
path: operations/user-subscription
description: Lists all of a user's subscriptions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/subscriptions
operation_ids:
    - user-subscription-get-user-subscriptions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get User Subscriptions

`GET /user/subscriptions`

Operation ID: `user-subscription-get-user-subscriptions`

Lists all of a user's subscriptions.

## Definition

```yaml
{"operationId": "user-subscription-get-user-subscriptions", "summary": "Get User Subscriptions", "description": "Lists all of a user's subscriptions.", "responses": {"200": {"description": "Get User Subscriptions response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_user_subscription_response_collection"}}}}, "4XX": {"description": "Get User Subscriptions response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_user_subscription_response_collection"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Subscription"], "x-api-token-group": ["Billing Write", "Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "user.subscriptions", "x-fern-sdk-method-name": "get"}
```
