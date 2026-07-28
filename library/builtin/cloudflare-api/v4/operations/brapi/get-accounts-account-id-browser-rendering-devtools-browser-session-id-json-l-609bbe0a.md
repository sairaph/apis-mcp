---
title: Get a target by ID.
page_id: operation-get-accounts-account-id-browser-rendering-devtools-browser-session-id-js-b2cf022f
path: operations/brapi
description: Returns the debuggable target with the given ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/list/{target_id}
operation_ids:
    - brapi-get_DevtoolsJsonTarget
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a target by ID.

`GET /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/list/{target_id}`

Operation ID: `brapi-get_DevtoolsJsonTarget`

Returns the debuggable target with the given ID.

## Definition

```yaml
{"operationId": "brapi-get_DevtoolsJsonTarget", "summary": "Get a target by ID.", "description": "Returns the debuggable target with the given ID.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "session_id", "in": "path", "description": "Browser session ID.", "required": true, "schema": {"description": "Browser session ID.", "type": "string", "format": "uuid"}}, {"name": "target_id", "in": "path", "description": "Target ID.", "required": true, "schema": {"description": "Target ID.", "type": "string", "pattern": "^[a-zA-Z0-9]+$"}}], "responses": {"200": {"description": "The target with the given ID.", "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"description": "Target description.", "type": "string"}, "devtoolsFrontendUrl": {"description": "DevTools frontend URL.", "type": "string"}, "id": {"description": "Target ID.", "type": "string"}, "title": {"description": "Title of the target.", "type": "string"}, "type": {"description": "Target type (page, background_page, worker, etc.).", "type": "string"}, "url": {"description": "URL of the target.", "type": "string"}, "webSocketDebuggerUrl": {"description": "WebSocket URL for debugging this target.", "type": "string"}}, "required": ["id", "type", "url"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "404": {"description": "Target not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-api-token-group": ["Browser Rendering Write", "Browser Rendering Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.browser.targets", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
