---
title: Close a browser target.
page_id: operation-get-accounts-account-id-browser-rendering-devtools-browser-session-id-js-7f2e6205
path: operations/brapi
description: Closes a specific browser target (tab, page, etc.) by its ID. Returns 'Target is closing' on success or an error if the target is not found.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/close/{target_id}
operation_ids:
    - brapi-get_DevtoolsJsonClose
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Close a browser target.

`GET /accounts/{account_id}/browser-rendering/devtools/browser/{session_id}/json/close/{target_id}`

Operation ID: `brapi-get_DevtoolsJsonClose`

Closes a specific browser target (tab, page, etc.) by its ID. Returns 'Target is closing' on success or an error if the target is not found.

## Definition

```yaml
{"operationId": "brapi-get_DevtoolsJsonClose", "summary": "Close a browser target.", "description": "Closes a specific browser target (tab, page, etc.) by its ID. Returns 'Target is closing' on success or an error if the target is not found.", "parameters": [{"name": "session_id", "in": "path", "description": "Browser session ID.", "required": true, "schema": {"description": "Browser session ID.", "type": "string", "format": "uuid"}}, {"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "target_id", "in": "path", "description": "Target ID to close.", "required": true, "schema": {"description": "Target ID to close.", "type": "string", "pattern": "^[a-zA-Z0-9]+$"}}], "responses": {"200": {"description": "Target is closing.", "content": {"application/json": {"schema": {"type": "object", "properties": {"message": {"description": "Target is closing.", "type": "string"}}, "required": ["message"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "404": {"description": "Target not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"description": "Error code.", "type": "number"}, "message": {"description": "Error message.", "type": "string"}}, "required": ["message", "code"], "type": "object"}}, "success": {"description": "Response status.", "type": "boolean"}}, "required": ["success"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-api-token-group": ["Browser Rendering Write", "Browser Rendering Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.browser.targets", "x-fern-sdk-method-name": "close", "x-forge-hidden": true}
```
