---
title: Get Event Subscription
page_id: operation-get-accounts-account-id-event-subscriptions-subscriptions-subscription-i-5c1260bc
path: operations/queue
description: Get details about an existing event subscription
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/event_subscriptions/subscriptions/{subscription_id}
operation_ids:
    - subscriptions-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Event Subscription

`GET /accounts/{account_id}/event_subscriptions/subscriptions/{subscription_id}`

Operation ID: `subscriptions-get`

Get details about an existing event subscription

## Definition

```yaml
{"operationId": "subscriptions-get", "summary": "Get Event Subscription", "description": "Get details about an existing event subscription", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "subscription_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "responses": {"200": {"description": "Details about an event subscription", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/mq_event-subscription"}}, "type": "object"}]}}}}, "404": {"description": "Event subscription does not exist", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.subscriptions", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
