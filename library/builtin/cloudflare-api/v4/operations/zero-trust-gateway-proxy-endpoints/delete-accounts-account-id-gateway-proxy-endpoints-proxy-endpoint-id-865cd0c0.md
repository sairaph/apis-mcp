---
title: Delete a proxy endpoint
page_id: operation-delete-accounts-account-id-gateway-proxy-endpoints-proxy-endpoint-id-51627d5a
path: operations/zero-trust-gateway-proxy-endpoints
description: Delete a configured Zero Trust Gateway proxy endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/gateway/proxy_endpoints/{proxy_endpoint_id}
operation_ids:
    - zero-trust-gateway-proxy-endpoints-delete-proxy-endpoint
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a proxy endpoint

`DELETE /accounts/{account_id}/gateway/proxy_endpoints/{proxy_endpoint_id}`

Operation ID: `zero-trust-gateway-proxy-endpoints-delete-proxy-endpoint`

Delete a configured Zero Trust Gateway proxy endpoint.

## Definition

```yaml
{"operationId": "zero-trust-gateway-proxy-endpoints-delete-proxy-endpoint", "summary": "Delete a proxy endpoint", "description": "Delete a configured Zero Trust Gateway proxy endpoint.", "parameters": [{"name": "proxy_endpoint_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-3"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Returns a deleted proxy endpoint response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_empty_response"}}}}, "4XX": {"description": "Returns a deleted proxy endpoint response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"type": "object"}, {"$ref": "#/components/schemas/zero-trust-gateway_empty_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway proxy endpoints"]}
```
