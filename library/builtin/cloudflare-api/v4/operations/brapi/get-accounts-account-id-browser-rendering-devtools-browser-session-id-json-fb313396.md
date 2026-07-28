---
title: List targets.
page_id: operation-get-accounts-account-id-browser-rendering-devtools-browser-session-id-js-953e794b
path: operations/brapi
description: Returns a list of all debuggable targets including tabs, pages, service workers, and other browser contexts.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json
operation_ids:
    - brapi-get_DevtoolsJson
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List targets.

`GET /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json`

Operation ID: `brapi-get_DevtoolsJson`

Returns a list of all debuggable targets including tabs, pages, service workers, and other browser contexts.

## Definition

```yaml
{"operationId": "brapi-get_DevtoolsJson", "summary": "List targets.", "description": "Returns a list of all debuggable targets including tabs, pages, service workers, and other browser contexts.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "session_id", "in": "path", "description": "Browser session ID.", "required": true, "schema": {"description": "Browser session ID.", "type": "string", "format": "uuid"}}, {"name": "liveViewUrlExpiresInMs", "in": "query", "description": "How long the live view URLs remain valid, in milliseconds (max 60 minutes)", "schema": {"description": "How long the live view URLs remain valid, in milliseconds (max 60 minutes)", "type": "number", "default": 300000, "maximum": 3600000, "minimum": 60000}}], "responses": {"200": {"description": "List of targets.", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"description": {"description": "Target description.", "type": "string"}, "devtoolsFrontendUrl": {"description": "DevTools frontend URL.", "type": "string"}, "id": {"description": "Target ID.", "type": "string"}, "title": {"description": "Title of the target.", "type": "string"}, "type": {"description": "Target type (page, background_page, worker, etc.).", "type": "string"}, "url": {"description": "URL of the target.", "type": "string"}, "webSocketDebuggerUrl": {"description": "WebSocket URL for debugging this target.", "type": "string"}}, "required": ["id", "type", "url"], "type": "object"}}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-api-token-group": ["Browser Rendering Write", "Browser Rendering Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.json", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
