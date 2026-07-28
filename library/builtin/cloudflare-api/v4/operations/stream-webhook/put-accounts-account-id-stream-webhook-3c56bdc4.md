---
title: Create webhooks
page_id: operation-put-accounts-account-id-stream-webhook-a724b781
path: operations/stream-webhook
description: Creates a webhook notification.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/stream/webhook
operation_ids:
    - stream-webhook-create-webhooks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create webhooks

`PUT /accounts/{account_id}/stream/webhook`

Operation ID: `stream-webhook-create-webhooks`

Creates a webhook notification.

## Definition

```yaml
{"operationId": "stream-webhook-create-webhooks", "summary": "Create webhooks", "description": "Creates a webhook notification.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_webhook_request"}}}}, "responses": {"200": {"description": "Create webhooks response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_webhook_response_single"}}}}, "4XX": {"description": "Create webhooks response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Webhook"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.webhooks", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
