---
title: Add a webhook
page_id: operation-post-accounts-account-id-realtime-kit-app-id-webhooks-f3efa749
path: operations/webhooks
description: Adds a new webhook to an App.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/webhooks
operation_ids:
    - addWebhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add a webhook

`POST /accounts/{account_id}/realtime/kit/{app_id}/webhooks`

Operation ID: `addWebhook`

Adds a new webhook to an App.

## Definition

```yaml
{"operationId": "addWebhook", "summary": "Add a webhook", "description": "Adds a new webhook to an App.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/realtimekit_WebhookRequest"}}}}, "responses": {"201": {"description": "Webhook registered successfully", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/realtimekit_WebhookSuccessResponse"}}}}, "400": {"description": "Error - malformed request", "content": {"application/json": {"examples": {"default": {"value": "{\n  success: false,\n  error: { code: 400, message: 'BAD_REQUEST: \"name\" is required' }\n}\n"}}, "schema": {"$ref": "#/components/schemas/realtimekit_ErrorResponse"}}}}, "401": {"description": "Invalid credentials", "content": {"application/json": {"examples": {"default": {"value": "{\n  message: \"Unauthorized\"\n}\n"}}}}}}, "security": [{"api_token": []}], "tags": ["Webhooks"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
