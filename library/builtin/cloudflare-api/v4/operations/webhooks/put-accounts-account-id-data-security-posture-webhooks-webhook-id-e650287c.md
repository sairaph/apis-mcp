---
title: Update an existing webhook configuration
page_id: operation-put-accounts-account-id-data-security-posture-webhooks-webhook-id-dfa5324c
path: operations/webhooks
description: Updates an existing webhook configuration with new settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/data-security/posture/webhooks/{webhook_id}
operation_ids:
    - UpdateWebhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an existing webhook configuration

`PUT /accounts/{account_id}/data-security/posture/webhooks/{webhook_id}`

Operation ID: `UpdateWebhook`

Updates an existing webhook configuration with new settings.

## Definition

```yaml
{"operationId": "UpdateWebhook", "summary": "Update an existing webhook configuration", "description": "Updates an existing webhook configuration with new settings.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_WebhookId"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_UpdateWebhookRequest"}}}}, "responses": {"200": {"description": "OK: Webhook updated successfully", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/posture-api_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/posture-api_Webhook"}}, "type": "object"}]}}}}, "400": {"description": "Bad Request: Invalid request parameters", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "401": {"description": "Unauthorized: Authentication required", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "403": {"description": "Forbidden: Insufficient permissions", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "404": {"description": "Not Found: Webhook not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["webhooks"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
