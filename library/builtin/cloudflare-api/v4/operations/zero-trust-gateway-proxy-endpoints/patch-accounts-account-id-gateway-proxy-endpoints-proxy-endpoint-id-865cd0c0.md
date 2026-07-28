---
title: Update a proxy endpoint
page_id: operation-patch-accounts-account-id-gateway-proxy-endpoints-proxy-endpoint-id-ad303541
path: operations/zero-trust-gateway-proxy-endpoints
description: Update a configured Zero Trust Gateway proxy endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/gateway/proxy_endpoints/{proxy_endpoint_id}
operation_ids:
    - zero-trust-gateway-proxy-endpoints-update-proxy-endpoint
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a proxy endpoint

`PATCH /accounts/{account_id}/gateway/proxy_endpoints/{proxy_endpoint_id}`

Operation ID: `zero-trust-gateway-proxy-endpoints-update-proxy-endpoint`

Update a configured Zero Trust Gateway proxy endpoint.

## Definition

```yaml
{"operationId": "zero-trust-gateway-proxy-endpoints-update-proxy-endpoint", "summary": "Update a proxy endpoint", "description": "Update a configured Zero Trust Gateway proxy endpoint.", "parameters": [{"name": "proxy_endpoint_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"ips": {"$ref": "#/components/schemas/zero-trust-gateway_ips"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-6"}}}}}}, "responses": {"200": {"description": "Returns an updated proxy endpoint response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-5"}}}}, "4XX": {"description": "Returns an updated proxy endpoint response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-5"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway proxy endpoints"]}
```
