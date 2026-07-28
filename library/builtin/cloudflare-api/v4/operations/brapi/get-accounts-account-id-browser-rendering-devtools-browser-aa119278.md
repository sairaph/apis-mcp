---
title: Acquire and connect to browser session.
page_id: operation-get-accounts-account-id-browser-rendering-devtools-browser-7309b772
path: operations/brapi
description: Acquires and establishes a WebSocket connection to a browser session.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/browser
operation_ids:
    - brapi-get_DevtoolsBrowserAcquire
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Acquire and connect to browser session.

`GET /accounts/{account_id}/browser-rendering/devtools/browser`

Operation ID: `brapi-get_DevtoolsBrowserAcquire`

Acquires and establishes a WebSocket connection to a browser session.

## Definition

```yaml
{"operationId": "brapi-get_DevtoolsBrowserAcquire", "summary": "Acquire and connect to browser session.", "description": "Acquires and establishes a WebSocket connection to a browser session.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "keep_alive", "in": "query", "description": "Keep-alive time in ms (only valid when acquiring new session).", "schema": {"description": "Keep-alive time in ms (only valid when acquiring new session).", "type": "number", "default": 60000, "maximum": 1200000, "minimum": 10000}}, {"name": "lab", "in": "query", "description": "Use experimental browser.", "schema": {"description": "Use experimental browser.", "type": "boolean", "default": false}}, {"name": "recording", "in": "query", "schema": {"type": "boolean", "default": false}}], "responses": {"101": {"description": "WebSocket connection established successfully."}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "429": {"description": "Request failed due to rate limiting.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-api-token-group": ["Browser Rendering Write", "Browser Rendering Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.browser", "x-fern-sdk-method-name": "launch", "x-forge-hidden": true}
```
