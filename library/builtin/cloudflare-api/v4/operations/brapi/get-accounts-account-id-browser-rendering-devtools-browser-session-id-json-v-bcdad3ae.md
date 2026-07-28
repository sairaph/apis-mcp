---
title: Get browser version metadata.
page_id: operation-get-accounts-account-id-browser-rendering-devtools-browser-session-id-js-fc162e66
path: operations/brapi
description: Get browser version metadata.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/version
operation_ids:
    - brapi-get_DevtoolsJsonVersion
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get browser version metadata.

`GET /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/version`

Operation ID: `brapi-get_DevtoolsJsonVersion`

Get browser version metadata.

## Definition

```yaml
{"operationId": "brapi-get_DevtoolsJsonVersion", "summary": "Get browser version metadata.", "description": "Get browser version metadata.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "session_id", "in": "path", "description": "Browser session ID.", "required": true, "schema": {"description": "Browser session ID.", "type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Browser version information.", "content": {"application/json": {"schema": {"type": "object", "properties": {"Browser": {"description": "Browser name and version.", "type": "string"}, "Protocol-Version": {"description": "Chrome DevTools Protocol version.", "type": "string"}, "User-Agent": {"description": "User agent string.", "type": "string"}, "V8-Version": {"description": "V8 JavaScript engine version.", "type": "string"}, "WebKit-Version": {"description": "WebKit version.", "type": "string"}, "webSocketDebuggerUrl": {"description": "WebSocket URL for debugging the browser.", "type": "string"}}, "required": ["Browser", "Protocol-Version", "User-Agent", "V8-Version", "WebKit-Version", "webSocketDebuggerUrl"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-api-token-group": ["Browser Rendering Write", "Browser Rendering Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.browser", "x-fern-sdk-method-name": "version", "x-forge-hidden": true}
```
