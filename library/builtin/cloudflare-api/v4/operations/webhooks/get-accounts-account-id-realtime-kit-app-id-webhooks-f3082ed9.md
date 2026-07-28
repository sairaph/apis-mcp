---
title: Fetch all webhooks details
page_id: operation-get-accounts-account-id-realtime-kit-app-id-webhooks-68a7d5d8
path: operations/webhooks
description: Returns details of all webhooks for an App.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/webhooks
operation_ids:
    - getAllWebhooks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all webhooks details

`GET /accounts/{account_id}/realtime/kit/{app_id}/webhooks`

Operation ID: `getAllWebhooks`

Returns details of all webhooks for an App.

## Definition

```yaml
{"operationId": "getAllWebhooks", "summary": "Fetch all webhooks details", "description": "Returns details of all webhooks for an App.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}], "responses": {"200": {"description": "Operation successful", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/realtimekit_WebhooksListSuccessResponse"}}}}, "401": {"description": "Invalid credentials", "content": {"application/json": {"examples": {"default": {"value": "{\n  message: \"Unauthorized\"\n}\n"}}}}}}, "security": [{"api_token": []}], "tags": ["Webhooks"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
