---
title: Fetch participants list of a session
page_id: operation-get-accounts-account-id-realtime-kit-app-id-sessions-session-id-particip-7a18d10a
path: operations/sessions
description: Returns a list of participants for the given session ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/participants
operation_ids:
    - GetSessionParticipants
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch participants list of a session

`GET /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/participants`

Operation ID: `GetSessionParticipants`

Returns a list of participants for the given session ID.

## Definition

```yaml
{"operationId": "GetSessionParticipants", "summary": "Fetch participants list of a session", "description": "Returns a list of participants for the given session ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "search", "in": "query", "description": "The search query string. You can search using participant ID, custom participant ID, or display name.", "schema": {"type": "string"}}, {"$ref": "#/components/parameters/realtimekit_pageNo"}, {"$ref": "#/components/parameters/realtimekit_perPage"}, {"$ref": "#/components/parameters/realtimekit_sortOrder"}, {"$ref": "#/components/parameters/realtimekit_participantsSortBy"}, {"name": "include_peer_events", "in": "query", "description": "if true, response includes all the peer events of participants.", "schema": {"type": "boolean", "default": false}}, {"name": "view", "in": "query", "description": "In breakout room sessions, the view parameter can be set to `raw` for session specific duration for participants or `consolidated` to accumulate breakout room durations.", "schema": {"type": "string", "default": "raw", "enum": ["raw", "consolidated"]}}, {"name": "session_id", "in": "path", "description": "ID of the session", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetSessionParticipants"}}, "security": [{"api_token": []}], "tags": ["Sessions"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
