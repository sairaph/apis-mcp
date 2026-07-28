---
title: Update a tunnel route (CIDR Endpoint)
page_id: operation-patch-accounts-account-id-teamnet-routes-network-ip-network-encoded-35222cdb
path: operations/tunnel-routing
description: Updates an existing private network route in an account. The CIDR in `ip_network_encoded` must be written in URL-encoded format.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/teamnet/routes/network/{ip_network_encoded}
operation_ids:
    - tunnel-route-update-a-tunnel-route-with-cidr
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a tunnel route (CIDR Endpoint)

`PATCH /accounts/{account_id}/teamnet/routes/network/{ip_network_encoded}`

Operation ID: `tunnel-route-update-a-tunnel-route-with-cidr`

Updates an existing private network route in an account. The CIDR in `ip_network_encoded` must be written in URL-encoded format.

## Definition

```yaml
{"operationId": "tunnel-route-update-a-tunnel-route-with-cidr", "summary": "Update a tunnel route (CIDR Endpoint)", "description": "Updates an existing private network route in an account. The CIDR in `ip_network_encoded` must be written in URL-encoded format.", "parameters": [{"name": "ip_network_encoded", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_ip_network_encoded"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "responses": {"200": {"description": "Update a tunnel route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_route_response_single"}}}}, "4XX": {"description": "Update a tunnel route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_route_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Routing"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-ignore": true, "x-stainless-deprecation-message": "This endpoint and its related APIs are deprecated in favor of the equivalent Tunnel Route (without CIDR) APIs."}
```
