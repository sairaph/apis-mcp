---
title: Delete a webhook configuration
page_id: operation-delete-accounts-account-id-data-security-posture-webhooks-webhook-id-dde520a0
path: operations/webhooks
description: |-
    Soft deletes a webhook configuration by its unique identifier.
    The webhook will be marked as deleted and will no longer be available for use.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/data-security/posture/webhooks/{webhook_id}
operation_ids:
    - DeleteWebhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a webhook configuration

`DELETE /accounts/{account_id}/data-security/posture/webhooks/{webhook_id}`

Operation ID: `DeleteWebhook`

Soft deletes a webhook configuration by its unique identifier.
The webhook will be marked as deleted and will no longer be available for use.

## Definition

```yaml
{"operationId": "DeleteWebhook", "summary": "Delete a webhook configuration", "description": "Soft deletes a webhook configuration by its unique identifier.\nThe webhook will be marked as deleted and will no longer be available for use.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_WebhookId"}], "responses": {"200": {"description": "OK: Webhook successfully deleted", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_api-response-common"}}}}, "400": {"description": "Bad Request: Invalid webhook ID or webhook already deleted"}, "409": {"description": "Conflict: webhook is associated with one or more policies and cannot be deleted", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["webhooks"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
