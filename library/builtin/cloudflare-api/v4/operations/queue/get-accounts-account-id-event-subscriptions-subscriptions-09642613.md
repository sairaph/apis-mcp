---
title: List Event Subscriptions
page_id: operation-get-accounts-account-id-event-subscriptions-subscriptions-7bbf45d9
path: operations/queue
description: Get a paginated list of event subscriptions with optional sorting and filtering
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/event_subscriptions/subscriptions
operation_ids:
    - subscriptions-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Event Subscriptions

`GET /accounts/{account_id}/event_subscriptions/subscriptions`

Operation ID: `subscriptions-list`

Get a paginated list of event subscriptions with optional sorting and filtering

## Definition

```yaml
{"operationId": "subscriptions-list", "summary": "List Event Subscriptions", "description": "Get a paginated list of event subscriptions with optional sorting and filtering", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/mq_identifier"}}, {"name": "page", "in": "query", "description": "Page number for pagination", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of items per page", "schema": {"type": "integer", "default": 20, "maximum": 100, "minimum": 1}}, {"name": "order", "in": "query", "description": "Field to sort by", "schema": {"type": "string", "default": "name", "enum": ["created_at", "name", "enabled", "source"]}}, {"name": "direction", "in": "query", "description": "Sort direction", "schema": {"type": "string", "default": "asc", "enum": ["asc", "desc"]}}], "responses": {"200": {"description": "List of event subscriptions", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/mq_api-v4-success"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/mq_event-subscription"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Number of items in current page", "type": "integer"}, "page": {"description": "Current page number", "type": "integer"}, "per_page": {"description": "Items per page", "type": "integer"}, "total_count": {"description": "Total number of items", "type": "integer"}, "total_pages": {"description": "Total number of pages", "type": "integer"}}, "required": ["count", "total_count", "page", "per_page", "total_pages"]}}, "type": "object"}]}}}}, "4XX": {"description": "Failure response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/mq_api-v4-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Queue"], "x-api-token-group": ["Queues Write", "Queues Read", "Workers Scripts Write", "Workers Scripts Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "queues.subscriptions", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
