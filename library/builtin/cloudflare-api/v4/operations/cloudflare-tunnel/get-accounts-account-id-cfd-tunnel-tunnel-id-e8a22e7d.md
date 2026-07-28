---
title: Get a Cloudflare Tunnel
page_id: operation-get-accounts-account-id-cfd-tunnel-tunnel-id-8c72da3e
path: operations/cloudflare-tunnel
description: Fetches a single Cloudflare Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}
operation_ids:
    - cloudflare-tunnel-get-a-cloudflare-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Cloudflare Tunnel

`GET /accounts/{account_id}/cfd_tunnel/{tunnel_id}`

Operation ID: `cloudflare-tunnel-get-a-cloudflare-tunnel`

Fetches a single Cloudflare Tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-get-a-cloudflare-tunnel", "summary": "Get a Cloudflare Tunnel", "description": "Fetches a single Cloudflare Tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "responses": {"200": {"description": "Get a Cloudflare Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_cfd-tunnel-response-single"}}}}, "4XX": {"description": "Get a Cloudflare Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_cfd-tunnel-response-single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: cloudflared Write", "Cloudflare One Connector: cloudflared Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
