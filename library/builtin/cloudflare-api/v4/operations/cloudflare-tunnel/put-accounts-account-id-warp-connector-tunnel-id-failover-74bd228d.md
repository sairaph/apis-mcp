---
title: Trigger a manual failover for a WARP Connector Tunnel
page_id: operation-put-accounts-account-id-warp-connector-tunnel-id-failover-2d378c8d
path: operations/cloudflare-tunnel
description: Triggers a manual failover for a specific WARP Connector Tunnel, setting the specified client as the active connector. The tunnel must be configured for high availability (HA) and the client must be linked to the tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/warp_connector/{tunnel_id}/failover
operation_ids:
    - cloudflare-tunnel-manual-failover-warp-connector-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Trigger a manual failover for a WARP Connector Tunnel

`PUT /accounts/{account_id}/warp_connector/{tunnel_id}/failover`

Operation ID: `cloudflare-tunnel-manual-failover-warp-connector-tunnel`

Triggers a manual failover for a specific WARP Connector Tunnel, setting the specified client as the active connector. The tunnel must be configured for high availability (HA) and the client must be linked to the tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-manual-failover-warp-connector-tunnel", "summary": "Trigger a manual failover for a WARP Connector Tunnel", "description": "Triggers a manual failover for a specific WARP Connector Tunnel, setting the specified client as the active connector. The tunnel must be configured for high availability (HA) and the client must be linked to the tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"client_id": {"$ref": "#/components/schemas/tunnel_client_id_input"}}, "required": ["client_id"]}}}}, "responses": {"200": {"description": "Manual failover response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_empty_response"}}}}, "4XX": {"description": "Manual failover response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_empty_response"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: WARP Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.warp-connector.failover", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
