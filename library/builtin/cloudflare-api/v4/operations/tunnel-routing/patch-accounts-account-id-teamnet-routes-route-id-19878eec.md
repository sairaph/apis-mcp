---
title: Update a tunnel route
page_id: operation-patch-accounts-account-id-teamnet-routes-route-id-90dbca8b
path: operations/tunnel-routing
description: Updates an existing private network route in an account. The fields that are meant to be updated should be provided in the body of the request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/teamnet/routes/{route_id}
operation_ids:
    - tunnel-route-update-a-tunnel-route
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a tunnel route

`PATCH /accounts/{account_id}/teamnet/routes/{route_id}`

Operation ID: `tunnel-route-update-a-tunnel-route`

Updates an existing private network route in an account. The fields that are meant to be updated should be provided in the body of the request.

## Definition

```yaml
{"operationId": "tunnel-route-update-a-tunnel-route", "summary": "Update a tunnel route", "description": "Updates an existing private network route in an account. The fields that are meant to be updated should be provided in the body of the request.", "parameters": [{"name": "route_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_route_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_route_comment"}, "network": {"$ref": "#/components/schemas/tunnel_ip_network"}, "tunnel_id": {"$ref": "#/components/schemas/tunnel_tunnel_id"}, "virtual_network_id": {"$ref": "#/components/schemas/tunnel_virtual_network_id_computed_optional"}}}}}}, "responses": {"200": {"description": "Update a tunnel route response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_route_response_single"}}}}, "4XX": {"description": "Update a tunnel route response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_route_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Routing"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.routes", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
