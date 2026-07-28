---
title: Create a tunnel route (CIDR Endpoint)
page_id: operation-post-accounts-account-id-teamnet-routes-network-ip-network-encoded-5ad2eb8d
path: operations/tunnel-routing
description: Routes a private network through a Cloudflare Tunnel. The CIDR in `ip_network_encoded` must be written in URL-encoded format.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/teamnet/routes/network/{ip_network_encoded}
operation_ids:
    - tunnel-route-create-a-tunnel-route-with-cidr
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a tunnel route (CIDR Endpoint)

`POST /accounts/{account_id}/teamnet/routes/network/{ip_network_encoded}`

Operation ID: `tunnel-route-create-a-tunnel-route-with-cidr`

Routes a private network through a Cloudflare Tunnel. The CIDR in `ip_network_encoded` must be written in URL-encoded format.

## Definition

```yaml
{"operationId": "tunnel-route-create-a-tunnel-route-with-cidr", "summary": "Create a tunnel route (CIDR Endpoint)", "description": "Routes a private network through a Cloudflare Tunnel. The CIDR in `ip_network_encoded` must be written in URL-encoded format.", "parameters": [{"name": "ip_network_encoded", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_ip_network_encoded"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_route_comment"}, "tunnel_id": {"$ref": "#/components/schemas/tunnel_tunnel_id"}, "virtual_network_id": {"$ref": "#/components/schemas/tunnel_virtual_network_id"}}, "required": ["tunnel_id"]}}}}, "responses": {"200": {"description": "Create a tunnel route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_route_response_single"}}}}, "4XX": {"description": "Create a tunnel route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_route_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Routing"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-ignore": true, "x-stainless-deprecation-message": "This endpoint and its related APIs are deprecated in favor of the equivalent Tunnel Route (without CIDR) APIs."}
```
