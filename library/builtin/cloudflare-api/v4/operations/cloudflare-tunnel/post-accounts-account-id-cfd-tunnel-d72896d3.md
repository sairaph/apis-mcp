---
title: Create a Cloudflare Tunnel
page_id: operation-post-accounts-account-id-cfd-tunnel-2d00f3f9
path: operations/cloudflare-tunnel
description: Creates a new Cloudflare Tunnel in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel
operation_ids:
    - cloudflare-tunnel-create-a-cloudflare-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Cloudflare Tunnel

`POST /accounts/{account_id}/cfd_tunnel`

Operation ID: `cloudflare-tunnel-create-a-cloudflare-tunnel`

Creates a new Cloudflare Tunnel in an account.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-create-a-cloudflare-tunnel", "summary": "Create a Cloudflare Tunnel", "description": "Creates a new Cloudflare Tunnel in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"config_src": {"$ref": "#/components/schemas/tunnel_config_src"}, "name": {"$ref": "#/components/schemas/tunnel_tunnel_name"}, "tunnel_secret": {"$ref": "#/components/schemas/tunnel_tunnel_secret"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Create a Cloudflare Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_cfd-tunnel-response-single"}}}}, "4XX": {"description": "Create a Cloudflare Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_cfd-tunnel-response-single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: cloudflared Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
