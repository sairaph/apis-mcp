---
title: Sync MCP Server Capabilities
page_id: operation-post-accounts-account-id-access-ai-controls-mcp-servers-id-sync-a1534eb8
path: operations/mcp-portal-servers
description: Syncs an MCP server's capabilities and returns the updated server state, including any connection errors.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/ai-controls/mcp/servers/{id}/sync
operation_ids:
    - mcp-portals-api-sync-server
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Sync MCP Server Capabilities

`POST /accounts/{account_id}/access/ai-controls/mcp/servers/{id}/sync`

Operation ID: `mcp-portals-api-sync-server`

Syncs an MCP server's capabilities and returns the updated server state, including any connection errors.

## Definition

```yaml
{"operationId": "mcp-portals-api-sync-server", "summary": "Sync MCP Server Capabilities", "description": "Syncs an MCP server's capabilities and returns the updated server state, including any connection errors.", "parameters": [{"name": "id", "in": "path", "required": true, "schema": {"description": "portal id", "type": "string", "example": "my-mcp-portal", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "a86a8f5c339544d7bdc89926de14fb8c"}}], "responses": {"200": {"description": "Sync completed. Check result.status and result.error_details for outcome.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"error": {"type": "string"}, "error_details": {"type": "object", "properties": {"cause": {"description": "Underlying error message", "type": "string"}, "is_upstream": {"description": "True = MCP server returned an error. False = couldn't reach the server", "type": "boolean"}, "mcp_code": {"description": "MCP protocol error code", "type": "number"}, "retryable": {"description": "Whether the error is transient and worth retrying", "type": "boolean"}, "status_code": {"description": "HTTP status code from the server", "type": "number"}}, "readOnly": true}, "status": {"type": "string", "enum": ["waiting", "ready", "stale", "error"]}}}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["MCP Portal Servers"], "x-api-token-group": ["MCP Portals Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.mcp_portals"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
