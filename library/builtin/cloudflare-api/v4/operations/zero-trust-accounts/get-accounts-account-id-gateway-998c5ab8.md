---
title: Get Zero Trust account information
page_id: operation-get-accounts-account-id-gateway-09bbf52e
path: operations/zero-trust-accounts
description: Retrieve information about the current Zero Trust account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway
operation_ids:
    - zero-trust-accounts-get-zero-trust-account-information
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust account information

`GET /accounts/{account_id}/gateway`

Operation ID: `zero-trust-accounts-get-zero-trust-account-information`

Retrieve information about the current Zero Trust account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-get-zero-trust-account-information", "summary": "Get Zero Trust account information", "description": "Retrieve information about the current Zero Trust account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "responses": {"200": {"description": "Zero Trust account information response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway_account"}}}}, "4XX": {"description": "Zero Trust account information response failure.", "content": {"application/json": {"schema": {"allOf": [{"type": "object"}, {"$ref": "#/components/schemas/zero-trust-gateway_gateway_account"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
