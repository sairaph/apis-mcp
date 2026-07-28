---
title: List Zero Trust Gateway rules inherited from the parent account
page_id: operation-get-accounts-account-id-gateway-rules-tenant-65aa0b67
path: operations/zero-trust-gateway-rules
description: List Zero Trust Gateway rules for the parent account of an account in the MSP configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/rules/tenant
operation_ids:
    - zero-trust-gateway-rules-list-zero-trust-gateway-rules-tenant
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Zero Trust Gateway rules inherited from the parent account

`GET /accounts/{account_id}/gateway/rules/tenant`

Operation ID: `zero-trust-gateway-rules-list-zero-trust-gateway-rules-tenant`

List Zero Trust Gateway rules for the parent account of an account in the MSP configuration.

## Definition

```yaml
{"operationId": "zero-trust-gateway-rules-list-zero-trust-gateway-rules-tenant", "summary": "List Zero Trust Gateway rules inherited from the parent account", "description": "List Zero Trust Gateway rules for the parent account of an account in the MSP configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "List Zero Trust Gateway rules response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-6"}}}}, "4XX": {"description": "List Zero Trust Gateway rules response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"type": "object"}, {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-6"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway rules"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.rules", "x-fern-sdk-method-name": "list-tenant", "x-forge-hidden": true}
```
