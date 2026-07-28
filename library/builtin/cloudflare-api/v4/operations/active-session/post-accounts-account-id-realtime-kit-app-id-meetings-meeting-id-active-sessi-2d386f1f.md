---
title: Kick all participants
page_id: operation-post-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-active-7c35119c
path: operations/active-session
description: Kicks all participants from an active session for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session/kick-all
operation_ids:
    - KickAllParticipants
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Kick all participants

`POST /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session/kick-all`

Operation ID: `KickAllParticipants`

Kicks all participants from an active session for the given meeting ID.

## Definition

```yaml
{"operationId": "KickAllParticipants", "summary": "Kick all participants", "description": "Kicks all participants from an active session for the given meeting ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meetingId"}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_KickAllParticipants"}}, "security": [{"api_token": []}], "tags": ["Active session"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
