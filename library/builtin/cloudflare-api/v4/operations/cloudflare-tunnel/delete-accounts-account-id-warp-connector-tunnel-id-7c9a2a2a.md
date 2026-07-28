---
title: Delete a Warp Connector Tunnel
page_id: operation-delete-accounts-account-id-warp-connector-tunnel-id-97ef5acf
path: operations/cloudflare-tunnel
description: Deletes a Warp Connector Tunnel from an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/warp_connector/{tunnel_id}
operation_ids:
    - cloudflare-tunnel-delete-a-warp-connector-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Warp Connector Tunnel

`DELETE /accounts/{account_id}/warp_connector/{tunnel_id}`

Operation ID: `cloudflare-tunnel-delete-a-warp-connector-tunnel`

Deletes a Warp Connector Tunnel from an account.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-delete-a-warp-connector-tunnel", "summary": "Delete a Warp Connector Tunnel", "description": "Deletes a Warp Connector Tunnel from an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object"}}}}, "responses": {"200": {"description": "Delete a Warp Connector Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_warp-connector-response-single"}}}}, "4XX": {"description": "Delete a Warp Connector Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_warp-connector-response-single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: WARP Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.warp-connector", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
