---
title: Get a proxy endpoint
page_id: operation-get-accounts-account-id-gateway-proxy-endpoints-proxy-endpoint-id-62048990
path: operations/zero-trust-gateway-proxy-endpoints
description: Get a single Zero Trust Gateway proxy endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/proxy_endpoints/{proxy_endpoint_id}
operation_ids:
    - zero-trust-gateway-proxy-endpoints-proxy-endpoint-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a proxy endpoint

`GET /accounts/{account_id}/gateway/proxy_endpoints/{proxy_endpoint_id}`

Operation ID: `zero-trust-gateway-proxy-endpoints-proxy-endpoint-details`

Get a single Zero Trust Gateway proxy endpoint.

## Definition

```yaml
{"operationId": "zero-trust-gateway-proxy-endpoints-proxy-endpoint-details", "summary": "Get a proxy endpoint", "description": "Get a single Zero Trust Gateway proxy endpoint.", "parameters": [{"name": "proxy_endpoint_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Returns a proxy endpoint response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-5"}}}}, "4XX": {"description": "Returns a proxy endpoint response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"type": "object"}, {"$ref": "#/components/schemas/zero-trust-gateway_single_response-5"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway proxy endpoints"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.proxy-endpoints", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
