---
title: Get WARP Connector Tunnel connector
page_id: operation-get-accounts-account-id-warp-connector-tunnel-id-connectors-connector-id-291f8dde
path: operations/cloudflare-tunnel
description: Fetches connector and connection details for a WARP Connector Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/warp_connector/{tunnel_id}/connectors/{connector_id}
operation_ids:
    - cloudflare-tunnel-get-warp-connector-tunnel-connector
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get WARP Connector Tunnel connector

`GET /accounts/{account_id}/warp_connector/{tunnel_id}/connectors/{connector_id}`

Operation ID: `cloudflare-tunnel-get-warp-connector-tunnel-connector`

Fetches connector and connection details for a WARP Connector Tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-get-warp-connector-tunnel-connector", "summary": "Get WARP Connector Tunnel connector", "description": "Fetches connector and connection details for a WARP Connector Tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}, {"name": "connector_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_client_id"}}], "responses": {"200": {"description": "Get WARP Connector Tunnel connector response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_tunnel_warp_connector_client_response"}}}}, "4XX": {"description": "Get WARP Connector Tunnel connector response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_tunnel_warp_connector_client_response"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: WARP Write", "Cloudflare One Connector: WARP Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.warp-connector.connectors", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
