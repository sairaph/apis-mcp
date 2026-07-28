---
title: Clean up Cloudflare Tunnel connections
page_id: operation-delete-accounts-account-id-cfd-tunnel-tunnel-id-connections-5eb2d8f6
path: operations/cloudflare-tunnel
description: Removes a connection (aka Cloudflare Tunnel Connector) from a Cloudflare Tunnel independently of its current state. If no connector id (client_id) is provided all connectors will be removed. We recommend running this command after rotating tokens.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}/connections
operation_ids:
    - cloudflare-tunnel-clean-up-cloudflare-tunnel-connections
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Clean up Cloudflare Tunnel connections

`DELETE /accounts/{account_id}/cfd_tunnel/{tunnel_id}/connections`

Operation ID: `cloudflare-tunnel-clean-up-cloudflare-tunnel-connections`

Removes a connection (aka Cloudflare Tunnel Connector) from a Cloudflare Tunnel independently of its current state. If no connector id (client_id) is provided all connectors will be removed. We recommend running this command after rotating tokens.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-clean-up-cloudflare-tunnel-connections", "summary": "Clean up Cloudflare Tunnel connections", "description": "Removes a connection (aka Cloudflare Tunnel Connector) from a Cloudflare Tunnel independently of its current state. If no connector id (client_id) is provided all connectors will be removed. We recommend running this command after rotating tokens.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}, {"name": "client_id", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_client_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object"}}}}, "responses": {"200": {"description": "Clean up Cloudflare Tunnel connections response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_empty_response"}}}}, "4XX": {"description": "Clean up Cloudflare Tunnel connections response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_empty_response"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: cloudflared Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared.connections", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
