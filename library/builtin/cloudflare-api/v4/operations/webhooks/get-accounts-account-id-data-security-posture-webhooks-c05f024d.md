---
title: List webhook configurations
page_id: operation-get-accounts-account-id-data-security-posture-webhooks-28f372c4
path: operations/webhooks
description: |-
    Retrieves all webhook configurations for the authenticated account.
    Returns an array of webhook configurations that can be used to send finding notifications.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/data-security/posture/webhooks
operation_ids:
    - ListWebhooks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List webhook configurations

`GET /accounts/{account_id}/data-security/posture/webhooks`

Operation ID: `ListWebhooks`

Retrieves all webhook configurations for the authenticated account.
Returns an array of webhook configurations that can be used to send finding notifications.

## Definition

```yaml
{"operationId": "ListWebhooks", "summary": "List webhook configurations", "description": "Retrieves all webhook configurations for the authenticated account.\nReturns an array of webhook configurations that can be used to send finding notifications.", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}], "responses": {"200": {"description": "OK: Successful HTTP request", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_webhook-list-response"}}}}, "401": {"description": "Unauthorized: Authentication required", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "403": {"description": "Forbidden: Insufficient permissions", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["webhooks"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
