---
title: Fetch all supported webhook events
page_id: operation-get-accounts-account-id-realtime-kit-app-id-webhooks-all-c455cbb3
path: operations/webhooks
description: Returns the list of webhook event names supported by RealtimeKit.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/webhooks/all
operation_ids:
    - getAllWebhookEvents
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all supported webhook events

`GET /accounts/{account_id}/realtime/kit/{app_id}/webhooks/all`

Operation ID: `getAllWebhookEvents`

Returns the list of webhook event names supported by RealtimeKit.

## Definition

```yaml
{"operationId": "getAllWebhookEvents", "summary": "Fetch all supported webhook events", "description": "Returns the list of webhook event names supported by RealtimeKit.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}], "responses": {"200": {"description": "Operation successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/realtimekit_WebhookEventsSuccessResponse"}}}}}, "security": [{"api_token": []}], "tags": ["Webhooks"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
