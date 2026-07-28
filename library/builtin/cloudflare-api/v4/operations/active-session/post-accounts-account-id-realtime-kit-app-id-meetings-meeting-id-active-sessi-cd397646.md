---
title: Create a poll
page_id: operation-post-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-active-83e2ecb6
path: operations/active-session
description: Creates a new poll in an active session for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session/poll
operation_ids:
    - CreatePoll
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a poll

`POST /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-session/poll`

Operation ID: `CreatePoll`

Creates a new poll in an active session for the given meeting ID.

## Definition

```yaml
{"operationId": "CreatePoll", "summary": "Create a poll", "description": "Creates a new poll in an active session for the given meeting ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meetingId"}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_CreatePollBody"}, "responses": {"201": {"$ref": "#/components/responses/realtimekit_CreatePoll"}, "400": {"description": "Bad Request"}}, "security": [{"api_token": []}], "tags": ["Active session"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
