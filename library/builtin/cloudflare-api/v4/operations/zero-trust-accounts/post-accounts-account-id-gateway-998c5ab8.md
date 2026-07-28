---
title: Create Zero Trust account
page_id: operation-post-accounts-account-id-gateway-ba919678
path: operations/zero-trust-accounts
description: Create a Zero Trust account for an existing Cloudflare account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway
operation_ids:
    - zero-trust-accounts-create-zero-trust-account
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Zero Trust account

`POST /accounts/{account_id}/gateway`

Operation ID: `zero-trust-accounts-create-zero-trust-account`

Create a Zero Trust account for an existing Cloudflare account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-create-zero-trust-account", "summary": "Create Zero Trust account", "description": "Create a Zero Trust account for an existing Cloudflare account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "responses": {"200": {"description": "Create Zero Trust account response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway_account"}}}}, "4XX": {"description": "Create Zero Trust account response failure.", "content": {"application/json": {"schema": {"allOf": [{"type": "object"}, {"$ref": "#/components/schemas/zero-trust-gateway_gateway_account"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
