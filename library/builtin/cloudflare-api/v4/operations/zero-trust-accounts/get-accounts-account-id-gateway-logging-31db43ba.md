---
title: Get logging settings for the Zero Trust account
page_id: operation-get-accounts-account-id-gateway-logging-d3afd023
path: operations/zero-trust-accounts
description: Retrieve the current logging settings for the Zero Trust account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/logging
operation_ids:
    - zero-trust-accounts-get-logging-settings-for-the-zero-trust-account
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get logging settings for the Zero Trust account

`GET /accounts/{account_id}/gateway/logging`

Operation ID: `zero-trust-accounts-get-logging-settings-for-the-zero-trust-account`

Retrieve the current logging settings for the Zero Trust account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-get-logging-settings-for-the-zero-trust-account", "summary": "Get logging settings for the Zero Trust account", "description": "Retrieve the current logging settings for the Zero Trust account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "responses": {"200": {"description": "Logging settings retrieval response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-logging-settings-response"}}}}, "4XX": {"description": "Logging settings retrieval response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-logging-settings-response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.logging", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
