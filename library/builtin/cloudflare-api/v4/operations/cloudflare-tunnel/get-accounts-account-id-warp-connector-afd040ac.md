---
title: List Warp Connector Tunnels
page_id: operation-get-accounts-account-id-warp-connector-ce1c605b
path: operations/cloudflare-tunnel
description: Lists and filters Warp Connector Tunnels in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/warp_connector
operation_ids:
    - cloudflare-tunnel-list-warp-connector-tunnels
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Warp Connector Tunnels

`GET /accounts/{account_id}/warp_connector`

Operation ID: `cloudflare-tunnel-list-warp-connector-tunnels`

Lists and filters Warp Connector Tunnels in an account.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-list-warp-connector-tunnels", "summary": "List Warp Connector Tunnels", "description": "Lists and filters Warp Connector Tunnels in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "name", "in": "query", "schema": {"description": "A user-friendly name for the tunnel.", "type": "string", "example": "blog"}}, {"name": "is_deleted", "in": "query", "schema": {"description": "If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If empty, all tunnels will be included.", "type": "boolean", "example": true}}, {"name": "existed_at", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_existed_at"}}, {"name": "uuid", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}, {"name": "was_active_at", "in": "query", "schema": {"type": "string", "format": "date-time", "example": "2009-11-10T23:00:00Z"}}, {"name": "was_inactive_at", "in": "query", "schema": {"type": "string", "format": "date-time", "example": "2009-11-10T23:00:00Z"}}, {"name": "include_prefix", "in": "query", "schema": {"type": "string", "example": "vpc1-"}}, {"name": "exclude_prefix", "in": "query", "schema": {"type": "string", "example": "vpc1-"}}, {"name": "status", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_status"}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_per_page"}}, {"name": "page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_page_number"}}], "responses": {"200": {"description": "List Warp Connector Tunnels response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_warp-connector-response-collection"}}}}, "4XX": {"description": "List Warp Connector Tunnels response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_warp-connector-response-collection"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: WARP Write", "Cloudflare One Connector: WARP Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.warp-connector", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
