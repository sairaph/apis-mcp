---
title: List tunnel routes
page_id: operation-get-accounts-account-id-teamnet-routes-662a9cb7
path: operations/tunnel-routing
description: Lists and filters private network routes in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/teamnet/routes
operation_ids:
    - tunnel-route-list-tunnel-routes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tunnel routes

`GET /accounts/{account_id}/teamnet/routes`

Operation ID: `tunnel-route-list-tunnel-routes`

Lists and filters private network routes in an account.

## Definition

```yaml
{"operationId": "tunnel-route-list-tunnel-routes", "summary": "List tunnel routes", "description": "Lists and filters private network routes in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "comment", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_route_comment"}}, {"name": "is_deleted", "in": "query", "schema": {"description": "If `true`, only include deleted routes. If `false`, exclude deleted routes. If empty, all routes will be included.", "type": "boolean"}}, {"name": "network_subset", "in": "query", "schema": {"description": "If set, only list routes that are contained within this IP range.", "allOf": [{"$ref": "#/components/schemas/tunnel_ip_network"}]}}, {"name": "network_superset", "in": "query", "schema": {"description": "If set, only list routes that contain this IP range.", "allOf": [{"$ref": "#/components/schemas/tunnel_ip_network"}]}}, {"name": "existed_at", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_existed_at"}}, {"name": "tunnel_id", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}, {"name": "route_id", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_route_id"}}, {"name": "tun_types", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_tunnel_types"}}, {"name": "virtual_network_id", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_virtual_network_id"}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_per_page"}}, {"name": "page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_page_number"}}], "responses": {"200": {"description": "List tunnel routes response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_teamnet_response_collection"}}}}, "4XX": {"description": "List tunnel routes response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_teamnet_response_collection"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Routing"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare One Networks Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.routes", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
