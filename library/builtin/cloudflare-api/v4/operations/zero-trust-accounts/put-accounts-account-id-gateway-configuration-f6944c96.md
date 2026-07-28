---
title: Update Zero Trust account configuration
page_id: operation-put-accounts-account-id-gateway-configuration-c7b3d8ba
path: operations/zero-trust-accounts
description: Update the current Zero Trust account configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/gateway/configuration
operation_ids:
    - zero-trust-accounts-update-zero-trust-account-configuration.
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zero Trust account configuration

`PUT /accounts/{account_id}/gateway/configuration`

Operation ID: `zero-trust-accounts-update-zero-trust-account-configuration.`

Update the current Zero Trust account configuration.

## Definition

```yaml
{"operationId": "zero-trust-accounts-update-zero-trust-account-configuration.", "summary": "Update Zero Trust account configuration", "description": "Update the current Zero Trust account configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-settings"}}}}, "responses": {"200": {"description": "Zero Trust account configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway_account_config"}}}}, "4XX": {"description": "Zero Trust account configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_gateway_account_config"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.configurations", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
