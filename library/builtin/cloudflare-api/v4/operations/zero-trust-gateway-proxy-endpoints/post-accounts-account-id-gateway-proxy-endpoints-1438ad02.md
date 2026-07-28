---
title: Create a proxy endpoint
page_id: operation-post-accounts-account-id-gateway-proxy-endpoints-a90a4327
path: operations/zero-trust-gateway-proxy-endpoints
description: Create a new Zero Trust Gateway proxy endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/proxy_endpoints
operation_ids:
    - zero-trust-gateway-proxy-endpoints-create-proxy-endpoint
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a proxy endpoint

`POST /accounts/{account_id}/gateway/proxy_endpoints`

Operation ID: `zero-trust-gateway-proxy-endpoints-create-proxy-endpoint`

Create a new Zero Trust Gateway proxy endpoint.

## Definition

```yaml
{"operationId": "zero-trust-gateway-proxy-endpoints-create-proxy-endpoint", "summary": "Create a proxy endpoint", "description": "Create a new Zero Trust Gateway proxy endpoint.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"kind": {"description": "The proxy endpoint kind.", "type": "string", "default": "ip", "enum": ["ip", "identity"]}}, "discriminator": {"mapping": {"identity": "#/components/schemas/zero-trust-gateway_proxy-endpoint-identity-create", "ip": "#/components/schemas/zero-trust-gateway_proxy-endpoint-ip-create"}, "propertyName": "kind"}, "oneOf": [{"$ref": "#/components/schemas/zero-trust-gateway_proxy-endpoint-ip-create"}, {"$ref": "#/components/schemas/zero-trust-gateway_proxy-endpoint-identity-create"}]}}}}, "responses": {"200": {"description": "Returns a created proxy endpoint response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-5"}}}}, "4XX": {"description": "Returns a created proxy endpoint response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-5"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway proxy endpoints"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.proxy-endpoints", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
