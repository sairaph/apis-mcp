---
title: Create a Warp Connector Tunnel
page_id: operation-post-accounts-account-id-warp-connector-d67b1805
path: operations/cloudflare-tunnel
description: Creates a new Warp Connector Tunnel in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/warp_connector
operation_ids:
    - cloudflare-tunnel-create-a-warp-connector-tunnel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Warp Connector Tunnel

`POST /accounts/{account_id}/warp_connector`

Operation ID: `cloudflare-tunnel-create-a-warp-connector-tunnel`

Creates a new Warp Connector Tunnel in an account.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-create-a-warp-connector-tunnel", "summary": "Create a Warp Connector Tunnel", "description": "Creates a new Warp Connector Tunnel in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"ha": {"description": "Indicates that the tunnel will be created to be highly available. If omitted, defaults to false.", "type": "boolean", "default": false}, "name": {"$ref": "#/components/schemas/tunnel_tunnel_name"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Create a Warp Connector Tunnel response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_warp-connector-response-single"}}}}, "4XX": {"description": "Create a Warp Connector Tunnel response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_warp-connector-response-single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: WARP Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.warp-connector", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
