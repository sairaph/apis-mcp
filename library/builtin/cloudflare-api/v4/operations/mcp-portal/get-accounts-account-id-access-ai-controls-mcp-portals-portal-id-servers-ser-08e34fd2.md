---
title: Resolve the OAuth redirect_uri the admin must register at the upstream
page_id: operation-get-accounts-account-id-access-ai-controls-mcp-portals-portal-id-servers-f33bc92a
path: operations/mcp-portal
description: Returns the redirect URI the gateway will actually send to the upstream OAuth provider for this (portal, server) pair. Stable for the lifetime of the configuration unless MCP-22 rollout state changes. Use this value verbatim when registering the OAuth app at the upstream.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/ai-controls/mcp/portals/{portal_id}/servers/{server_id}/effective-redirect-uri
operation_ids:
    - mcp-portals-api-effective-redirect-uri
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Resolve the OAuth redirect_uri the admin must register at the upstream

`GET /accounts/{account_id}/access/ai-controls/mcp/portals/{portal_id}/servers/{server_id}/effective-redirect-uri`

Operation ID: `mcp-portals-api-effective-redirect-uri`

Returns the redirect URI the gateway will actually send to the upstream OAuth provider for this (portal, server) pair. Stable for the lifetime of the configuration unless MCP-22 rollout state changes. Use this value verbatim when registering the OAuth app at the upstream.

## Definition

```yaml
{"operationId": "mcp-portals-api-effective-redirect-uri", "summary": "Resolve the OAuth redirect_uri the admin must register at the upstream", "description": "Returns the redirect URI the gateway will actually send to the upstream OAuth provider for this (portal, server) pair. Stable for the lifetime of the configuration unless MCP-22 rollout state changes. Use this value verbatim when registering the OAuth app at the upstream.", "parameters": [{"name": "portal_id", "in": "path", "required": true, "schema": {"description": "portal id", "type": "string", "example": "my-mcp-portal", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$"}}, {"name": "server_id", "in": "path", "required": true, "schema": {"type": "string", "example": "github-mcp"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "a86a8f5c339544d7bdc89926de14fb8c"}}], "responses": {"200": {"description": "Returns the effective redirect URI and its source.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"redirect_uri": {"type": "string", "format": "uri"}, "source": {"type": "string", "enum": ["per_portal", "shared_mcp22"]}}, "required": ["redirect_uri", "source"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["MCP Portal"], "x-api-token-group": ["MCP Portals Write", "MCP Portals Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.mcp_portals"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
