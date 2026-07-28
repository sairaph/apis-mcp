---
title: Fetch details of a session
page_id: operation-get-accounts-account-id-realtime-kit-app-id-sessions-session-id-ed208e9d
path: operations/sessions
description: Returns data of the given session ID including recording details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}
operation_ids:
    - GetSessionDetails
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch details of a session

`GET /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}`

Operation ID: `GetSessionDetails`

Returns data of the given session ID including recording details.

## Definition

```yaml
{"operationId": "GetSessionDetails", "summary": "Fetch details of a session", "description": "Returns data of the given session ID including recording details.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "include_breakout_rooms", "in": "query", "description": "List all breakout rooms", "schema": {"type": "boolean", "default": false}}, {"name": "session_id", "in": "path", "description": "ID of the session", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetSessionDetails"}}, "security": [{"api_token": []}], "tags": ["Sessions"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
