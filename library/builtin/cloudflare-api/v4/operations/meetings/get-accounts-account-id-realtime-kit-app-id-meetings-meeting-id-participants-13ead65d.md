---
title: Fetch all participants of a meeting
page_id: operation-get-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-particip-46e46143
path: operations/meetings
description: Returns all participants detail for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/participants
operation_ids:
    - get_meeting_participants
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all participants of a meeting

`GET /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/participants`

Operation ID: `get_meeting_participants`

Returns all participants detail for the given meeting ID.

## Definition

```yaml
{"operationId": "get_meeting_participants", "summary": "Fetch all participants of a meeting", "description": "Returns all participants detail for the given meeting ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_pageNo"}, {"$ref": "#/components/parameters/realtimekit_perPage"}, {"name": "meeting_id", "in": "path", "description": "ID of the meeting. Fetch the meeting ID using the create a meeting API.", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetAllParticipants"}, "500": {"$ref": "#/components/responses/realtimekit_GenericError"}}, "security": [{"api_token": []}], "tags": ["Meetings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
