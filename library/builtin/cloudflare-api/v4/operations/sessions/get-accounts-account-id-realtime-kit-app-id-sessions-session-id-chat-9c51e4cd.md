---
title: Fetch all chat messages of a session
page_id: operation-get-accounts-account-id-realtime-kit-app-id-sessions-session-id-chat-a85b86f5
path: operations/sessions
description: Returns a URL to download all chat messages of the session ID in CSV format.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/chat
operation_ids:
    - GetSessionChat
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all chat messages of a session

`GET /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/chat`

Operation ID: `GetSessionChat`

Returns a URL to download all chat messages of the session ID in CSV format.

## Definition

```yaml
{"operationId": "GetSessionChat", "summary": "Fetch all chat messages of a session", "description": "Returns a URL to download all chat messages of the session ID in CSV format.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "session_id", "in": "path", "description": "ID of the session", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetSessionChat"}}, "security": [{"api_token": []}], "tags": ["Sessions"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
