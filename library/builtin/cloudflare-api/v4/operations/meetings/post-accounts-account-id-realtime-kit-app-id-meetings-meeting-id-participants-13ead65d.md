---
title: Add a participant
page_id: operation-post-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-partici-b2ca9fdf
path: operations/meetings
description: Adds a participant to the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/participants
operation_ids:
    - add_participant
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add a participant

`POST /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/participants`

Operation ID: `add_participant`

Adds a participant to the given meeting ID.

## Definition

```yaml
{"operationId": "add_participant", "summary": "Add a participant", "description": "Adds a participant to the given meeting ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "meeting_id", "in": "path", "description": "ID of the meeting. Fetch the meeting ID using the create a meeting API.", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_AddParticipantBody"}, "responses": {"201": {"$ref": "#/components/responses/realtimekit_AddParticipant"}, "500": {"$ref": "#/components/responses/realtimekit_GenericError"}}, "security": [{"api_token": []}], "tags": ["Meetings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
