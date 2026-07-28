---
title: Create a new webhook configuration
page_id: operation-post-accounts-account-id-data-security-posture-webhooks-99dec6c1
path: operations/webhooks
description: Creates a new webhook configuration for sending finding notifications to external endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/webhooks
operation_ids:
    - CreateWebhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new webhook configuration

`POST /accounts/{account_id}/data-security/posture/webhooks`

Operation ID: `CreateWebhook`

Creates a new webhook configuration for sending finding notifications to external endpoints.

## Definition

```yaml
{"operationId": "CreateWebhook", "summary": "Create a new webhook configuration", "description": "Creates a new webhook configuration for sending finding notifications to external endpoints.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_CreateWebhookRequest"}}}}, "responses": {"200": {"description": "OK: Webhook created successfully", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/posture-api_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/posture-api_Webhook"}}, "type": "object"}]}}}}, "400": {"description": "Bad Request: Invalid request parameters", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "401": {"description": "Unauthorized: Authentication required", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "403": {"description": "Forbidden: Insufficient permissions", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "422": {"description": "Unprocessable Entity: Webhook configuration limit reached for this account", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["webhooks"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
