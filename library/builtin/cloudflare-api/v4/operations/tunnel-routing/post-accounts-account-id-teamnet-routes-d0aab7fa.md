---
title: Create a tunnel route
page_id: operation-post-accounts-account-id-teamnet-routes-c4b9cd97
path: operations/tunnel-routing
description: Routes a private network through a Cloudflare Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/teamnet/routes
operation_ids:
    - tunnel-route-create-a-tunnel-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a tunnel route

`POST /accounts/{account_id}/teamnet/routes`

Operation ID: `tunnel-route-create-a-tunnel-route`

Routes a private network through a Cloudflare Tunnel.

## Definition

```yaml
{"operationId": "tunnel-route-create-a-tunnel-route", "summary": "Create a tunnel route", "description": "Routes a private network through a Cloudflare Tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_route_comment"}, "network": {"$ref": "#/components/schemas/tunnel_ip_network"}, "tunnel_id": {"$ref": "#/components/schemas/tunnel_tunnel_id"}, "virtual_network_id": {"$ref": "#/components/schemas/tunnel_virtual_network_id_computed_optional"}}, "required": ["network", "tunnel_id"]}}}}, "responses": {"200": {"description": "Create a tunnel route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_route_response_single"}}}}, "4XX": {"description": "Create a tunnel route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_route_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Routing"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.routes", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
