---
title: Get webhook configuration by ID
page_id: operation-get-accounts-account-id-data-security-posture-webhooks-webhook-id-f501dce7
path: operations/webhooks
description: Retrieves a specific webhook configuration by its unique identifier.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/webhooks/{webhook_id}
operation_ids:
    - GetWebhookConfigByID
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get webhook configuration by ID

`GET /accounts/{account_id}/data-security/posture/webhooks/{webhook_id}`

Operation ID: `GetWebhookConfigByID`

Retrieves a specific webhook configuration by its unique identifier.

## Definition

```yaml
{"operationId": "GetWebhookConfigByID", "summary": "Get webhook configuration by ID", "description": "Retrieves a specific webhook configuration by its unique identifier.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_WebhookId"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_webhook-response"}}}}, "400": {"description": "Bad Request: Invalid webhook ID", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "401": {"description": "Unauthorized: Authentication required", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "403": {"description": "Forbidden: Insufficient permissions", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "404": {"description": "Not Found: Webhook not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["webhooks"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
