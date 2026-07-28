---
title: Update Zero Trust account logging settings
page_id: operation-put-accounts-account-id-gateway-logging-d75ff171
path: operations/zero-trust-accounts
description: Update logging settings for the current Zero Trust account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/gateway/logging
operation_ids:
    - zero-trust-accounts-update-logging-settings-for-the-zero-trust-account
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zero Trust account logging settings

`PUT /accounts/{account_id}/gateway/logging`

Operation ID: `zero-trust-accounts-update-logging-settings-for-the-zero-trust-account`

Update logging settings for the current Zero Trust account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-update-logging-settings-for-the-zero-trust-account", "summary": "Update Zero Trust account logging settings", "description": "Update logging settings for the current Zero Trust account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-logging-settings"}}}}, "responses": {"200": {"description": "Logging settings update response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-logging-settings-response"}}}}, "4XX": {"description": "Logging settings update response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-logging-settings-response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.logging", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
