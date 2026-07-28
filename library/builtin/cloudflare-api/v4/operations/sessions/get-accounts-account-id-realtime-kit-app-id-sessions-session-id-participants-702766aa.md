---
title: Fetch details of a participant
page_id: operation-get-accounts-account-id-realtime-kit-app-id-sessions-session-id-particip-016d30e9
path: operations/sessions
description: Returns details of the given participant ID along with call statistics for the given session ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/participants/{participant_id}
operation_ids:
    - GetParticipantDetails
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch details of a participant

`GET /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/participants/{participant_id}`

Operation ID: `GetParticipantDetails`

Returns details of the given participant ID along with call statistics for the given session ID.

## Definition

```yaml
{"operationId": "GetParticipantDetails", "summary": "Fetch details of a participant", "description": "Returns details of the given participant ID along with call statistics for the given session ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "include_peer_events", "in": "query", "description": "if true, response includes all the peer events of participant.", "schema": {"type": "boolean", "default": false}}, {"name": "participant_id", "in": "path", "description": "ID of the participant", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "session_id", "in": "path", "description": "ID of the session", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetParticipantDetails"}}, "security": [{"api_token": []}], "tags": ["Sessions"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
