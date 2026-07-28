---
title: Get a browser session ID.
page_id: operation-post-accounts-account-id-browser-rendering-devtools-browser-62318c79
path: operations/brapi
description: Acquires a browser and returns its session ID and websocket URL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/browser
operation_ids:
    - brapi-post_DevtoolsAcquire
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a browser session ID.

`POST /accounts/{account_id}/browser-rendering/devtools/browser`

Operation ID: `brapi-post_DevtoolsAcquire`

Acquires a browser and returns its session ID and websocket URL.

## Definition

```yaml
{"operationId": "brapi-post_DevtoolsAcquire", "summary": "Get a browser session ID.", "description": "Acquires a browser and returns its session ID and websocket URL.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "keep_alive", "in": "query", "description": "Keep-alive time in milliseconds.", "schema": {"description": "Keep-alive time in milliseconds.", "type": "number", "default": 60000, "maximum": 1200000, "minimum": 10000}}, {"name": "lab", "in": "query", "description": "Use experimental browser.", "schema": {"description": "Use experimental browser.", "type": "boolean", "default": false}}, {"name": "targets", "in": "query", "description": "Include browser targets in response.", "schema": {"description": "Include browser targets in response.", "type": "boolean", "default": false}}, {"name": "liveViewUrlExpiresInMs", "in": "query", "description": "How long the live view URL remains valid, in milliseconds (max 60 minutes). Only used when targets is true.", "schema": {"description": "How long the live view URL remains valid, in milliseconds (max 60 minutes). Only used when targets is true.", "type": "number", "default": 300000, "maximum": 3600000, "minimum": 60000}}, {"name": "recording", "in": "query", "schema": {"type": "boolean", "default": false}}], "responses": {"200": {"description": "Returns a session ID ready to be connected to.", "content": {"application/json": {"schema": {"type": "object", "properties": {"sessionId": {"description": "Browser session ID.", "type": "string"}, "webSocketDebuggerUrl": {"description": "WebSocket URL for the session.", "type": "string"}}, "required": ["sessionId"]}}}}, "429": {"description": "Request failed due to rate limiting. The Retry-After header indicates when the client should retry the request.", "headers": {"Retry-After": {"description": "Number of seconds to wait before retrying the request.", "required": true, "schema": {"description": "Number of seconds to wait before retrying the request.", "type": "number", "example": 60}}}, "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "status": {"description": "Response status.", "type": "boolean"}}, "required": ["status"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-api-token-group": ["Browser Rendering Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.browser", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
