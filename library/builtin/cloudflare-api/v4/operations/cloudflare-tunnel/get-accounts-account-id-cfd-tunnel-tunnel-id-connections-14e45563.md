---
title: List Cloudflare Tunnel connections
page_id: operation-get-accounts-account-id-cfd-tunnel-tunnel-id-connections-4df1157d
path: operations/cloudflare-tunnel
description: Fetches connection details for a Cloudflare Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}/connections
operation_ids:
    - cloudflare-tunnel-list-cloudflare-tunnel-connections
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Cloudflare Tunnel connections

`GET /accounts/{account_id}/cfd_tunnel/{tunnel_id}/connections`

Operation ID: `cloudflare-tunnel-list-cloudflare-tunnel-connections`

Fetches connection details for a Cloudflare Tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-list-cloudflare-tunnel-connections", "summary": "List Cloudflare Tunnel connections", "description": "Fetches connection details for a Cloudflare Tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "responses": {"200": {"description": "List Cloudflare Tunnel connections response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_tunnel_connections_response"}}}}, "4XX": {"description": "List Cloudflare Tunnel connections response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_tunnel_connections_response"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: cloudflared Write", "Cloudflare One Connector: cloudflared Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared.connections", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
