---
title: List proxy endpoints
page_id: operation-get-accounts-account-id-gateway-proxy-endpoints-e5769003
path: operations/zero-trust-gateway-proxy-endpoints
description: List all Zero Trust Gateway proxy endpoints for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/proxy_endpoints
operation_ids:
    - zero-trust-gateway-proxy-endpoints-list-proxy-endpoints
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List proxy endpoints

`GET /accounts/{account_id}/gateway/proxy_endpoints`

Operation ID: `zero-trust-gateway-proxy-endpoints-list-proxy-endpoints`

List all Zero Trust Gateway proxy endpoints for an account.

## Definition

```yaml
{"operationId": "zero-trust-gateway-proxy-endpoints-list-proxy-endpoints", "summary": "List proxy endpoints", "description": "List all Zero Trust Gateway proxy endpoints for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Returns a list of proxy endpoints response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-9"}}}}, "4XX": {"description": "Returns a list of proxy endpoints response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection-9"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway proxy endpoints"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.proxy-endpoints", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
