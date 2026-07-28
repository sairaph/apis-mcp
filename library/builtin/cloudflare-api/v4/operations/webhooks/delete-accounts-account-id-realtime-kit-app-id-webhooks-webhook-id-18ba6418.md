---
title: Delete a webhook
page_id: operation-delete-accounts-account-id-realtime-kit-app-id-webhooks-webhook-id-e2722110
path: operations/webhooks
description: Removes a webhook for the given webhook ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/webhooks/{webhook_id}
operation_ids:
    - deleteWebhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a webhook

`DELETE /accounts/{account_id}/realtime/kit/{app_id}/webhooks/{webhook_id}`

Operation ID: `deleteWebhook`

Removes a webhook for the given webhook ID.

## Definition

```yaml
{"operationId": "deleteWebhook", "summary": "Delete a webhook", "description": "Removes a webhook for the given webhook ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "webhook_id", "in": "path", "description": "ID of the webhook", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Operation successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/realtimekit_WebhookSuccessResponse"}}}}, "400": {"description": "Error - malformed request", "content": {"application/json": {"examples": {"default": {"value": "{\n  \"success\": false,\n  \"error\": { \"code\": 400, \"message\": \"BAD_REQUEST \\\"id\\\" must be a valid GUID\" }\n}\n"}}, "schema": {"$ref": "#/components/schemas/realtimekit_ErrorResponse"}}}}, "401": {"description": "Invalid credentials", "content": {"application/json": {"examples": {"default": {"value": "{\n  \"message\": \"Unauthorized\"\n}\n"}}}}}}, "security": [{"api_token": []}], "tags": ["Webhooks"], "x-api-token-group": ["Realtime Admin", "Realtime"], "x-stability": "beta"}
```
