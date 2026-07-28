---
title: Get Zero Trust account configuration
page_id: operation-get-accounts-account-id-gateway-configuration-6f917e70
path: operations/zero-trust-accounts
description: Retrieve the current Zero Trust account configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/configuration
operation_ids:
    - zero-trust-accounts-get-zero-trust-account-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust account configuration

`GET /accounts/{account_id}/gateway/configuration`

Operation ID: `zero-trust-accounts-get-zero-trust-account-configuration`

Retrieve the current Zero Trust account configuration.

## Definition

```yaml
{"operationId": "zero-trust-accounts-get-zero-trust-account-configuration", "summary": "Get Zero Trust account configuration", "description": "Retrieve the current Zero Trust account configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "responses": {"200": {"description": "Zero Trust account configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway_account_config"}}}}, "4XX": {"description": "Zero Trust account configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_gateway_account_config"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.configurations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
