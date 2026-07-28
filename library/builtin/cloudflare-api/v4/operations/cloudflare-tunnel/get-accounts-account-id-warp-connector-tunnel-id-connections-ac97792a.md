---
title: List WARP Connector Tunnel connections
page_id: operation-get-accounts-account-id-warp-connector-tunnel-id-connections-1f9802be
path: operations/cloudflare-tunnel
description: Fetches connection details for a WARP Connector Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/warp_connector/{tunnel_id}/connections
operation_ids:
    - cloudflare-tunnel-list-warp-connector-tunnel-connections
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List WARP Connector Tunnel connections

`GET /accounts/{account_id}/warp_connector/{tunnel_id}/connections`

Operation ID: `cloudflare-tunnel-list-warp-connector-tunnel-connections`

Fetches connection details for a WARP Connector Tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-list-warp-connector-tunnel-connections", "summary": "List WARP Connector Tunnel connections", "description": "Fetches connection details for a WARP Connector Tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "responses": {"200": {"description": "List WARP Connector Tunnel connections response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_tunnel_warp_connector_connections_response"}}}}, "4XX": {"description": "List WARP Connector connections response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_tunnel_warp_connector_connections_response"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: WARP Write", "Cloudflare One Connector: WARP Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.warp-connector.connections", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
