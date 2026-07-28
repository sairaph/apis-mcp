---
title: List All Tunnels
page_id: operation-get-accounts-account-id-tunnels-f434569e
path: operations/cloudflare-tunnel
description: Lists and filters all types of Tunnels in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tunnels
operation_ids:
    - cloudflare-tunnel-list-all-tunnels
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List All Tunnels

`GET /accounts/{account_id}/tunnels`

Operation ID: `cloudflare-tunnel-list-all-tunnels`

Lists and filters all types of Tunnels in an account.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-list-all-tunnels", "summary": "List All Tunnels", "description": "Lists and filters all types of Tunnels in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "name", "in": "query", "schema": {"description": "A user-friendly name for the tunnel.", "type": "string", "example": "blog"}}, {"name": "is_deleted", "in": "query", "schema": {"description": "If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If empty, all tunnels will be included.", "type": "boolean", "example": true}}, {"name": "existed_at", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_existed_at"}}, {"name": "uuid", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id"}}, {"name": "was_active_at", "in": "query", "schema": {"type": "string", "format": "date-time", "example": "2009-11-10T23:00:00Z"}}, {"name": "was_inactive_at", "in": "query", "schema": {"type": "string", "format": "date-time", "example": "2009-11-10T23:00:00Z"}}, {"name": "include_prefix", "in": "query", "schema": {"type": "string", "example": "vpc1-"}}, {"name": "exclude_prefix", "in": "query", "schema": {"type": "string", "example": "vpc1-"}}, {"name": "tun_types", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_tunnel_types"}}, {"name": "status", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_status"}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_per_page"}}, {"name": "page", "in": "query", "schema": {"$ref": "#/components/schemas/tunnel_page_number"}}], "responses": {"200": {"description": "List Tunnels response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_tunnel-response-collection"}}}}, "4XX": {"description": "List Tunnels response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_tunnel-response-collection"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: cloudflared Write", "Cloudflare One Connector: cloudflared Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
