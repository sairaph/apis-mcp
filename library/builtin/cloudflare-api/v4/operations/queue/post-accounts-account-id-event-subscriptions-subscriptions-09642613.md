---
title: Create Event Subscription
page_id: operation-post-accounts-account-id-event-subscriptions-subscriptions-c23defcc
path: operations/queue
description: Create a new event subscription for a queue
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/event_subscriptions/subscriptions
operation_ids:
    - subscriptions-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Event Subscription

`POST /accounts/{account_id}/event_subscriptions/subscriptions`

Operation ID: `subscriptions-create`

Create a new event subscription for a queue

## Definition

```yaml
{"operationId": "subscriptions-create", "summary": "Create Event Subscription", "description": "Create a new event subscription for a queue", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"destination": {"$ref": "#/components/schemas/mq_event-destination"}, "enabled": {"description": "Whether the subscription is active", "type": "boolean", "x-auditable": true}, "events": {"description": "List of event types this subscription handles", "type": "array", "items": {"type": "string"}, "minItems": 1, "x-auditable": true}, "name": {"description": "Name of the subscription", "type": "string", "x-auditable": true}, "source": {"$ref": "#/components/schemas/mq_event-source"}}}}}}, "responses": {"200": {"description": "Successfully created event subscription", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"$ref": "#/components/schemas/mq_event-subscription"}}, "type": "object"}]}}}}, "400": {"description": "Invalid request body or validation errors", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}, "404": {"description": "Queue does not exist or resource not found on source", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}, "405": {"description": "Multiple subscriptions on same resource not supported", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Workers Scripts Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.subscriptions", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
