---
title: Get a Cloudflare Tunnel token
page_id: operation-get-accounts-account-id-cfd-tunnel-tunnel-id-token-994fecce
path: operations/cloudflare-tunnel
description: Gets the token used to associate cloudflared with a specific tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}/token
operation_ids:
    - cloudflare-tunnel-get-a-cloudflare-tunnel-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Cloudflare Tunnel token

`GET /accounts/{account_id}/cfd_tunnel/{tunnel_id}/token`

Operation ID: `cloudflare-tunnel-get-a-cloudflare-tunnel-token`

Gets the token used to associate cloudflared with a specific tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-get-a-cloudflare-tunnel-token", "summary": "Get a Cloudflare Tunnel token", "description": "Gets the token used to associate cloudflared with a specific tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "responses": {"200": {"description": "Get a Cloudflare Tunnel token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_tunnel_response_token"}}}}, "4XX": {"description": "Get a Cloudflare Tunnel token response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_tunnel_response_token"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: cloudflared Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared.token", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
