---
title: Get session details.
page_id: operation-get-accounts-account-id-browser-rendering-devtools-session-session-id-543a478d
path: operations/brapi
description: Get details for a specific browser session.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/browser-rendering/devtools/session/{session_id}
operation_ids:
    - brapi-get_DevtoolsSessionDetails
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get session details.

`GET /accounts/{account_id}/browser-rendering/devtools/session/{session_id}`

Operation ID: `brapi-get_DevtoolsSessionDetails`

Get details for a specific browser session.

## Definition

```yaml
{"operationId": "brapi-get_DevtoolsSessionDetails", "summary": "Get session details.", "description": "Get details for a specific browser session.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "session_id", "in": "path", "description": "Session ID.", "required": true, "schema": {"description": "Session ID.", "type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Returns the session details.", "content": {"application/json": {"schema": {"type": "object", "properties": {"closeReason": {"description": "Reason for session closure.", "type": "string"}, "closeReasonText": {"description": "Human-readable close reason.", "type": "string"}, "connectionEndTime": {"description": "Connection end time.", "type": "number"}, "connectionId": {"description": "Connection ID.", "type": "string"}, "connectionStartTime": {"description": "Connection start time.", "type": "number"}, "devtoolsFrontendUrl": {"description": "DevTools frontend URL.", "type": "string"}, "endTime": {"description": "Session end time.", "type": "number"}, "lastUpdated": {"description": "Last updated timestamp.", "type": "number"}, "sessionId": {"description": "Session ID.", "type": "string", "format": "uuid"}, "startTime": {"description": "Session start time.", "type": "number"}, "webSocketDebuggerUrl": {"description": "WebSocket URL for debugging this target.", "type": "string"}}, "nullable": true, "required": ["sessionId"]}}}}, "404": {"description": "Session not found."}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["brapi"], "x-api-token-group": ["Browser Rendering Write", "Browser Rendering Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.brapi.read"], "type": "string"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "browser-run.devtools.session", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
