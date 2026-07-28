---
title: Get a Cloudflare Tunnel management token
page_id: operation-post-accounts-account-id-cfd-tunnel-tunnel-id-management-8d55ca91
path: operations/cloudflare-tunnel
description: Gets a management token used to access the management resources (i.e. Streaming Logs) of a tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}/management
operation_ids:
    - cloudflare-tunnel-get-a-cloudflare-tunnel-management-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Cloudflare Tunnel management token

`POST /accounts/{account_id}/cfd_tunnel/{tunnel_id}/management`

Operation ID: `cloudflare-tunnel-get-a-cloudflare-tunnel-management-token`

Gets a management token used to access the management resources (i.e. Streaming Logs) of a tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-get-a-cloudflare-tunnel-management-token", "summary": "Get a Cloudflare Tunnel management token", "description": "Gets a management token used to access the management resources (i.e. Streaming Logs) of a tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"resources": {"type": "array", "items": {"$ref": "#/components/schemas/tunnel_management-resources"}}}, "required": ["resources"]}}}}, "responses": {"200": {"description": "Get a Cloudflare Tunnel management token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_tunnel_response_token"}}}}, "4XX": {"description": "Cloudflare API response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_tunnel_response_token"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: cloudflared Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared.management", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
