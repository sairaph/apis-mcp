---
title: Get a Warp Connector Tunnel
page_id: operation-get-accounts-account-id-warp-connector-tunnel-id-459e8e36
path: operations/cloudflare-tunnel
description: Fetches a single Warp Connector Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/warp_connector/{tunnel_id}
operation_ids:
    - cloudflare-tunnel-get-a-warp-connector-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Warp Connector Tunnel

`GET /accounts/{account_id}/warp_connector/{tunnel_id}`

Operation ID: `cloudflare-tunnel-get-a-warp-connector-tunnel`

Fetches a single Warp Connector Tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-get-a-warp-connector-tunnel", "summary": "Get a Warp Connector Tunnel", "description": "Fetches a single Warp Connector Tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "responses": {"200": {"description": "Get a Warp Connector Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_warp-connector-response-single"}}}}, "4XX": {"description": "Get a Warp Connector Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_warp-connector-response-single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: WARP Write", "Cloudflare One Connector: WARP Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.warp-connector", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
