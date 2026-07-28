---
title: Update a Cloudflare Tunnel
page_id: operation-patch-accounts-account-id-cfd-tunnel-tunnel-id-d4fc6ef5
path: operations/cloudflare-tunnel
description: Updates an existing Cloudflare Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}
operation_ids:
    - cloudflare-tunnel-update-a-cloudflare-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Cloudflare Tunnel

`PATCH /accounts/{account_id}/cfd_tunnel/{tunnel_id}`

Operation ID: `cloudflare-tunnel-update-a-cloudflare-tunnel`

Updates an existing Cloudflare Tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-update-a-cloudflare-tunnel", "summary": "Update a Cloudflare Tunnel", "description": "Updates an existing Cloudflare Tunnel.", "parameters": [{"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"$ref": "#/components/schemas/tunnel_tunnel_name"}, "tunnel_secret": {"$ref": "#/components/schemas/tunnel_tunnel_secret"}}}}}}, "responses": {"200": {"description": "Update a Cloudflare Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_cfd-tunnel-response-single"}}}}, "4XX": {"description": "Update a Cloudflare Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_cfd-tunnel-response-single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: cloudflared Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
