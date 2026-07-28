---
title: Get tunnel route by IP
page_id: operation-get-accounts-account-id-teamnet-routes-ip-ip-849aa337
path: operations/tunnel-routing
description: Fetches routes that contain the given IP address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/teamnet/routes/ip/{ip}
operation_ids:
    - tunnel-route-get-tunnel-route-by-ip
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get tunnel route by IP

`GET /accounts/{account_id}/teamnet/routes/ip/{ip}`

Operation ID: `tunnel-route-get-tunnel-route-by-ip`

Fetches routes that contain the given IP address.

## Definition

```yaml
{"operationId": "tunnel-route-get-tunnel-route-by-ip", "summary": "Get tunnel route by IP", "description": "Fetches routes that contain the given IP address.", "parameters": [{"name": "ip", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_ip"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "virtual_network_id", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_virtual_network_id"}}, {"name": "default_virtual_network_fallback", "in": "query", "description": "When the virtual_network_id parameter is not provided the request filter will default search routes that are in the default virtual network for the account. If this parameter is set to false, the search will include routes that do not have a virtual network.", "schema": {"type": "boolean", "default": true}}], "responses": {"200": {"description": "Get tunnel route by IP response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_teamnet_response_single"}}}}, "4XX": {"description": "Get tunnel route by IP response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_teamnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Routing"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare One Networks Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.routes.ips", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
