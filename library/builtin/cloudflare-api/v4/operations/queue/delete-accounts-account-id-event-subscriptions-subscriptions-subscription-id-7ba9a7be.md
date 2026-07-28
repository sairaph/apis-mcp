---
title: Delete Event Subscription
page_id: operation-delete-accounts-account-id-event-subscriptions-subscriptions-subscriptio-e90efd19
path: operations/queue
description: Delete an existing event subscription
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/event_subscriptions/subscriptions/{subscription_id}
operation_ids:
    - subscriptions-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Event Subscription

`DELETE /accounts/{account_id}/event_subscriptions/subscriptions/{subscription_id}`

Operation ID: `subscriptions-delete`

Delete an existing event subscription

## Definition

```yaml
{"operationId": "subscriptions-delete", "summary": "Delete Event Subscription", "description": "Delete an existing event subscription", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "subscription_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "responses": {"200": {"description": "Successfully created event subscription", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/mq_event-subscription"}}, "type": "object"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.subscriptions", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
