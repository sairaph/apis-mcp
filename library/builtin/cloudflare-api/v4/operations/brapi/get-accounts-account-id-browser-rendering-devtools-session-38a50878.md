---
title: List sessions.
page_id: operation-get-accounts-account-id-browser-rendering-devtools-session-331c080f
path: operations/brapi
description: List active browser sessions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/session
operation_ids:
    - brapi-get_DevtoolsSessionList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List sessions.

`GET /accounts/{account_id}/browser-rendering/devtools/session`

Operation ID: `brapi-get_DevtoolsSessionList`

List active browser sessions.

## Definition

```yaml
{"operationId": "brapi-get_DevtoolsSessionList", "summary": "List sessions.", "description": "List active browser sessions.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "limit", "in": "query", "schema": {"type": "number", "default": 200, "maximum": 200, "minimum": 1}}, {"name": "offset", "in": "query", "schema": {"type": "number", "default": 0, "minimum": 0}}], "responses": {"200": {"description": "Returns the account's sessions.", "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"closeReason": {"description": "Reason for session closure.", "type": "string"}, "closeReasonText": {"description": "Human-readable close reason.", "type": "string"}, "connectionEndTime": {"description": "Connection end time.", "type": "number"}, "connectionId": {"description": "Connection ID.", "type": "string"}, "connectionStartTime": {"description": "Connection start time.", "type": "number"}, "devtoolsFrontendUrl": {"description": "DevTools frontend URL.", "type": "string"}, "endTime": {"description": "Session end time.", "type": "number"}, "lastUpdated": {"description": "Last updated timestamp.", "type": "number"}, "sessionId": {"description": "Session ID.", "type": "string", "format": "uuid"}, "startTime": {"description": "Session start time.", "type": "number"}, "webSocketDebuggerUrl": {"description": "WebSocket URL for debugging this target.", "type": "string"}}, "required": ["sessionId"], "type": "object"}}}}}, "500": {"description": "Internal server error."}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-api-token-group": ["Browser Rendering Write", "Browser Rendering Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.session", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
