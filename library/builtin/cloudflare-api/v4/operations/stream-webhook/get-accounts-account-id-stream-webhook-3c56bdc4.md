---
title: View webhooks
page_id: operation-get-accounts-account-id-stream-webhook-d672c91c
path: operations/stream-webhook
description: Retrieves a list of webhooks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/webhook
operation_ids:
    - stream-webhook-view-webhooks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# View webhooks

`GET /accounts/{account_id}/stream/webhook`

Operation ID: `stream-webhook-view-webhooks`

Retrieves a list of webhooks.

## Definition

```yaml
{"operationId": "stream-webhook-view-webhooks", "summary": "View webhooks", "description": "Retrieves a list of webhooks.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "responses": {"200": {"description": "View webhooks response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_webhook_response_single"}}}}, "4XX": {"description": "View webhooks response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Webhook"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.webhooks", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
