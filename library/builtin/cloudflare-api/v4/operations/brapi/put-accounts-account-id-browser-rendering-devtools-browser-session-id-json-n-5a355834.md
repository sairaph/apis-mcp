---
title: Open a new browser tab.
page_id: operation-put-accounts-account-id-browser-rendering-devtools-browser-session-id-js-b28fc897
path: operations/brapi
description: Opens a new tab in the browser. Optionally specify a URL to navigate to.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/new
operation_ids:
    - brapi-put_DevtoolsJsonNew
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Open a new browser tab.

`PUT /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/new`

Operation ID: `brapi-put_DevtoolsJsonNew`

Opens a new tab in the browser. Optionally specify a URL to navigate to.

## Definition

```yaml
{"operationId": "brapi-put_DevtoolsJsonNew", "summary": "Open a new browser tab.", "description": "Opens a new tab in the browser. Optionally specify a URL to navigate to.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "session_id", "in": "path", "description": "Browser session ID.", "required": true, "schema": {"description": "Browser session ID.", "type": "string", "format": "uuid"}}, {"name": "url", "in": "query", "schema": {"type": "string", "format": "uri"}}, {"name": "liveViewUrlExpiresInMs", "in": "query", "description": "How long the live view URL remains valid, in milliseconds (max 60 minutes)", "schema": {"description": "How long the live view URL remains valid, in milliseconds (max 60 minutes)", "type": "number", "default": 300000, "maximum": 3600000, "minimum": 60000}}], "responses": {"200": {"description": "Information about the newly created tab.", "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"description": "Target description.", "type": "string"}, "devtoolsFrontendUrl": {"description": "DevTools frontend URL.", "type": "string"}, "id": {"description": "Target ID.", "type": "string"}, "title": {"description": "Title of the target.", "type": "string"}, "type": {"description": "Target type (page, background_page, worker, etc.).", "type": "string"}, "url": {"description": "URL of the target.", "type": "string"}, "webSocketDebuggerUrl": {"description": "WebSocket URL for debugging this target.", "type": "string"}}, "required": ["id", "type", "url"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.browser.targets", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
