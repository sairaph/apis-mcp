---
title: Delete webhooks
page_id: operation-delete-accounts-account-id-stream-webhook-cb2c7749
path: operations/stream-webhook
description: Deletes a webhook.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/webhook
operation_ids:
    - stream-webhook-delete-webhooks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete webhooks

`DELETE /accounts/{account_id}/stream/webhook`

Operation ID: `stream-webhook-delete-webhooks`

Deletes a webhook.

## Definition

```yaml
{"operationId": "stream-webhook-delete-webhooks", "summary": "Delete webhooks", "description": "Deletes a webhook.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "responses": {"200": {"description": "Delete webhooks response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_deleted_response"}}}}, "4XX": {"description": "Delete webhooks response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Webhook"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.webhooks", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
