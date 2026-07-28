---
title: Replace a webhook
page_id: operation-put-accounts-account-id-realtime-kit-app-id-webhooks-webhook-id-1adaa0e9
path: operations/webhooks
description: Replace all details for the given webhook ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/webhooks/{webhook_id}
operation_ids:
    - replaceWebhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace a webhook

`PUT /accounts/{account_id}/realtime/kit/{app_id}/webhooks/{webhook_id}`

Operation ID: `replaceWebhook`

Replace all details for the given webhook ID.

## Definition

```yaml
{"operationId": "replaceWebhook", "summary": "Replace a webhook", "description": "Replace all details for the given webhook ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "webhook_id", "in": "path", "description": "ID of the webhook", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"default": {"value": "{\n  \"name\": \"test\",\n  \"events\": [ \"meeting.started\" ],\n  \"url\": \"https://new-test-url.com\"\n}\n"}}, "schema": {"$ref": "#/components/schemas/realtimekit_WebhookRequest"}}}}, "responses": {"200": {"description": "Operation successful", "content": {"application/json": {"examples": {"default": {"value": "{\n  \"success\": true,\n  \"data\": {\n    id: \"901e9adf-ebd9-41f0-9098-28171bc04ddd\",\n    \"name\": \"test\",\n    \"events\": [ \"meeting.started\" ],\n    \"url\": \"https://new-test-url.com\",\n    \"created_at\": \"2021-09-09T10:25:12.330Z\",\n    \"updated_at\": \"2021-09-09T10:25:12.376Z\"\n  }\n}\n"}}, "schema": {"$ref": "#/components/schemas/realtimekit_WebhookSuccessResponse"}}}}, "400": {"description": "Error - malformed request", "content": {"application/json": {"examples": {"default": {"value": "{\n  \"success\": false,\n  \"error\": { \"code\": 400, \"message\": \"BAD_REQUEST \\\"id\\\" must be a valid GUID\" }\n}\n"}}, "schema": {"$ref": "#/components/schemas/realtimekit_ErrorResponse"}}}}, "401": {"description": "Invalid credentials", "content": {"application/json": {"examples": {"default": {"value": "{\n  message: \"Unauthorized\"\n}\n"}}}}}}, "security": [{"api_token": []}], "tags": ["Webhooks"], "x-api-token-group": ["Realtime Admin", "Realtime"], "x-stability": "beta"}
```
