---
title: Kick participants from an active session
page_id: operation-post-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-active-7e60e9ae
path: operations/active-session
description: Kicks one or more participants from an active session using user ID or custom participant ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session/kick
operation_ids:
    - KickPartcipants
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Kick participants from an active session

`POST /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session/kick`

Operation ID: `KickPartcipants`

Kicks one or more participants from an active session using user ID or custom participant ID.

## Definition

```yaml
{"operationId": "KickPartcipants", "summary": "Kick participants from an active session", "description": "Kicks one or more participants from an active session using user ID or custom participant ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meetingId"}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_KickParticipantsBody"}, "responses": {"200": {"$ref": "#/components/responses/realtimekit_KickParticipants"}, "404": {"$ref": "#/components/responses/realtimekit_ParticipantNotFound"}}, "security": [{"api_token": []}], "tags": ["Active session"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
