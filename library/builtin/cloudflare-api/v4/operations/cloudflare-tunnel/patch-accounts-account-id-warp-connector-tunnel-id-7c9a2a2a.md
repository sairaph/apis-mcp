---
title: Update a Warp Connector Tunnel
page_id: operation-patch-accounts-account-id-warp-connector-tunnel-id-80619959
path: operations/cloudflare-tunnel
description: Updates an existing Warp Connector Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/warp_connector/{tunnel_id}
operation_ids:
    - cloudflare-tunnel-update-a-warp-connector-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Warp Connector Tunnel

`PATCH /accounts/{account_id}/warp_connector/{tunnel_id}`

Operation ID: `cloudflare-tunnel-update-a-warp-connector-tunnel`

Updates an existing Warp Connector Tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-update-a-warp-connector-tunnel", "summary": "Update a Warp Connector Tunnel", "description": "Updates an existing Warp Connector Tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"$ref": "#/components/schemas/tunnel_tunnel_name"}, "tunnel_secret": {"$ref": "#/components/schemas/tunnel_tunnel_secret"}}}}}}, "responses": {"200": {"description": "Update a Warp Connector Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_warp-connector-response-single"}}}}, "4XX": {"description": "Update a Warp Connector Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_warp-connector-response-single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: WARP Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.warp-connector", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
