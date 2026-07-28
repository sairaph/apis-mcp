---
title: Mute all participants
page_id: operation-post-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-active-0ec06834
path: operations/active-session
description: Mutes all participants of an active session for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session/mute-all
operation_ids:
    - MuteAllParticipants
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Mute all participants

`POST /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session/mute-all`

Operation ID: `MuteAllParticipants`

Mutes all participants of an active session for the given meeting ID.

## Definition

```yaml
{"operationId": "MuteAllParticipants", "summary": "Mute all participants", "description": "Mutes all participants of an active session for the given meeting ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meetingId"}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_MuteAllParticipantsBody"}, "responses": {"200": {"$ref": "#/components/responses/realtimekit_MuteAllParticipants"}}, "security": [{"api_token": []}], "tags": ["Active session"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
