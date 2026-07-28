---
title: Get Cloudflare Tunnel connector
page_id: operation-get-accounts-account-id-cfd-tunnel-tunnel-id-connectors-connector-id-000f44b7
path: operations/cloudflare-tunnel
description: Fetches connector and connection details for a Cloudflare Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}/connectors/{connector_id}
operation_ids:
    - cloudflare-tunnel-get-cloudflare-tunnel-connector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Cloudflare Tunnel connector

`GET /accounts/{account_id}/cfd_tunnel/{tunnel_id}/connectors/{connector_id}`

Operation ID: `cloudflare-tunnel-get-cloudflare-tunnel-connector`

Fetches connector and connection details for a Cloudflare Tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-get-cloudflare-tunnel-connector", "summary": "Get Cloudflare Tunnel connector", "description": "Fetches connector and connection details for a Cloudflare Tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}, {"name": "connector_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_client_id"}}], "responses": {"200": {"description": "Get Cloudflare Tunnel connector response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_tunnel_client_response"}}}}, "4XX": {"description": "Get Cloudflare Tunnel connector response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_tunnel_client_response"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: cloudflared Write", "Cloudflare One Connector: cloudflared Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared.connectors", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
