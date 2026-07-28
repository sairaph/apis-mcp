---
title: Get tunnel route
page_id: operation-get-accounts-account-id-teamnet-routes-route-id-74c78e1c
path: operations/tunnel-routing
description: Get a private network route in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/teamnet/routes/{route_id}
operation_ids:
    - tunnel-route-get-tunnel-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get tunnel route

`GET /accounts/{account_id}/teamnet/routes/{route_id}`

Operation ID: `tunnel-route-get-tunnel-route`

Get a private network route in an account.

## Definition

```yaml
{"operationId": "tunnel-route-get-tunnel-route", "summary": "Get tunnel route", "description": "Get a private network route in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_route_id"}}], "responses": {"200": {"description": "Get a tunnel route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_route_response_single"}}}}, "4XX": {"description": "Get a tunnel route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_route_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Routing"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.routes", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
