---
title: Delete a MCP Portal
page_id: operation-delete-accounts-account-id-access-ai-controls-mcp-portals-id-765b21f9
path: operations/mcp-portal
description: Deletes an MCP portal from the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/ai-controls/mcp/portals/{id}
operation_ids:
    - mcp-portals-api-delete-portals
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a MCP Portal

`DELETE /accounts/{account_id}/access/ai-controls/mcp/portals/{id}`

Operation ID: `mcp-portals-api-delete-portals`

Deletes an MCP portal from the account.

## Definition

```yaml
{"operationId": "mcp-portals-api-delete-portals", "summary": "Delete a MCP Portal", "description": "Deletes an MCP portal from the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string", "example": "a86a8f5c339544d7bdc89926de14fb8c"}}, {"name": "id", "in": "path", "required": true, "schema": {"description": "portal id", "type": "string", "example": "my-mcp-portal", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$"}}], "responses": {"200": {"description": "Returns the Object if it was successfully deleted", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"allow_code_mode": {"description": "Allow remote code execution in Dynamic Workers (beta)", "type": "boolean", "example": true, "default": true}, "created_at": {"type": "string", "format": "date-time", "readOnly": true}, "created_by": {"type": "string", "readOnly": true}, "description": {"type": "string", "example": "This is my custom MCP Portal", "maxLength": 512}, "hostname": {"type": "string", "example": "exmaple.com", "pattern": "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9-]*[A-Za-z0-9])$"}, "id": {"description": "portal id", "type": "string", "example": "my-mcp-portal", "maxLength": 32, "minLength": 1, "pattern": "^[a-z0-9_]+(?:-[a-z0-9_]+)*$"}, "modified_at": {"type": "string", "format": "date-time", "readOnly": true}, "modified_by": {"type": "string", "readOnly": true}, "name": {"type": "string", "example": "My MCP Portal", "maxLength": 350}, "secure_web_gateway": {"description": "Route outbound MCP traffic through Zero Trust Secure Web Gateway", "type": "boolean", "example": false, "default": false}}, "required": ["id", "name", "hostname"]}, "success": {"type": "boolean"}}, "required": ["success", "result"]}}}}, "404": {"description": "Not Found", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "number", "example": 7002}, "message": {"type": "string", "example": "Not Found"}}, "required": ["code", "message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["MCP Portal"], "x-api-token-group": ["MCP Portals Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.mcp_portals"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
