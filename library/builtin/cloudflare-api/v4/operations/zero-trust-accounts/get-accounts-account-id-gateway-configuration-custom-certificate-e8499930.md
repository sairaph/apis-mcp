---
title: Get Zero Trust certificate configuration
page_id: operation-get-accounts-account-id-gateway-configuration-custom-certificate-d84a624a
path: operations/zero-trust-accounts
description: Retrieve the current Zero Trust certificate configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/configuration/custom_certificate
operation_ids:
    - zero-trust-accounts-get-zero-trust-certificate-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust certificate configuration

`GET /accounts/{account_id}/gateway/configuration/custom_certificate`

Operation ID: `zero-trust-accounts-get-zero-trust-certificate-configuration`

Retrieve the current Zero Trust certificate configuration.

## Definition

```yaml
{"operationId": "zero-trust-accounts-get-zero-trust-certificate-configuration", "summary": "Get Zero Trust certificate configuration", "description": "Retrieve the current Zero Trust certificate configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "responses": {"200": {"description": "Zero Trust account configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_custom-certificate-settings"}}}}, "4XX": {"description": "Zero Trust account configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_custom-certificate-settings"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.configurations.custom-certificate", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
