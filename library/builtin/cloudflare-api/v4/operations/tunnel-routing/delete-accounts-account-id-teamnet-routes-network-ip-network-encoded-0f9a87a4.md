---
title: Delete a tunnel route (CIDR Endpoint)
page_id: operation-delete-accounts-account-id-teamnet-routes-network-ip-network-encoded-8f60e36d
path: operations/tunnel-routing
description: Deletes a private network route from an account. The CIDR in `ip_network_encoded` must be written in URL-encoded format. If no virtual_network_id is provided it will delete the route from the default vnet. If no tun_type is provided it will fetch the type from the tunnel_id or if that is missing it will assume Cloudflare Tunnel as default. If tunnel_id is provided it will delete the route from that tunnel, otherwise it will delete the route based on the vnet and tun_type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/teamnet/routes/network/{ip_network_encoded}
operation_ids:
    - tunnel-route-delete-a-tunnel-route-with-cidr
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a tunnel route (CIDR Endpoint)

`DELETE /accounts/{account_id}/teamnet/routes/network/{ip_network_encoded}`

Operation ID: `tunnel-route-delete-a-tunnel-route-with-cidr`

Deletes a private network route from an account. The CIDR in `ip_network_encoded` must be written in URL-encoded format. If no virtual_network_id is provided it will delete the route from the default vnet. If no tun_type is provided it will fetch the type from the tunnel_id or if that is missing it will assume Cloudflare Tunnel as default. If tunnel_id is provided it will delete the route from that tunnel, otherwise it will delete the route based on the vnet and tun_type.

## Definition

```yaml
{"operationId": "tunnel-route-delete-a-tunnel-route-with-cidr", "summary": "Delete a tunnel route (CIDR Endpoint)", "description": "Deletes a private network route from an account. The CIDR in `ip_network_encoded` must be written in URL-encoded format. If no virtual_network_id is provided it will delete the route from the default vnet. If no tun_type is provided it will fetch the type from the tunnel_id or if that is missing it will assume Cloudflare Tunnel as default. If tunnel_id is provided it will delete the route from that tunnel, otherwise it will delete the route based on the vnet and tun_type.\n", "parameters": [{"name": "ip_network_encoded", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_ip_network_encoded"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "virtual_network_id", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_virtual_network_id"}}, {"name": "tun_type", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_tunnel_type"}}, {"name": "tunnel_id", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "responses": {"200": {"description": "Delete a tunnel route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_route_response_single"}}}}, "4XX": {"description": "Delete a tunnel route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_route_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Routing"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-ignore": true, "x-stainless-deprecation-message": "This endpoint and its related APIs are deprecated in favor of the equivalent Tunnel Route (without CIDR) APIs."}
```
