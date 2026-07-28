---
title: Delete User Subscription
page_id: operation-delete-user-subscriptions-identifier-335eea20
path: operations/user-subscription
description: Deletes a user's subscription.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /user/subscriptions/{identifier}
operation_ids:
    - user-subscription-delete-user-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete User Subscription

`DELETE /user/subscriptions/{identifier}`

Operation ID: `user-subscription-delete-user-subscription`

Deletes a user's subscription.

## Definition

```yaml
{"operationId": "user-subscription-delete-user-subscription", "summary": "Delete User Subscription", "description": "Deletes a user's subscription.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_schemas-identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete User Subscription response", "content": {"application/json": {"schema": {"type": "object", "properties": {"subscription_id": {"$ref": "#/components/schemas/bill-subs-api_schemas-identifier"}}}}}}, "4XX": {"description": "Delete User Subscription response failure", "content": {"application/json": {"schema": {"allOf": [{"properties": {"subscription_id": {"$ref": "#/components/schemas/bill-subs-api_schemas-identifier"}}, "type": "object"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["User Subscription"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:edit"]}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "user.subscriptions", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
