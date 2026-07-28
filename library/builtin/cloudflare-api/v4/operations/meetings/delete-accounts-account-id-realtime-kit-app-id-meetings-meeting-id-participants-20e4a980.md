---
title: Delete a participant
page_id: operation-delete-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-parti-d942af97
path: operations/meetings
description: Deletes a participant for the given meeting and participant ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/participants/{participant_id}
operation_ids:
    - delete_meeting_participant
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a participant

`DELETE /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/participants/{participant_id}`

Operation ID: `delete_meeting_participant`

Deletes a participant for the given meeting and participant ID.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meeting_id"}]
```

## Definition

```yaml
{"operationId": "delete_meeting_participant", "summary": "Delete a participant", "description": "Deletes a participant for the given meeting and participant ID.", "parameters": [{"name": "meeting_id", "in": "path", "description": "ID of the meeting. You can fetch the meeting ID using the create a meeting API.", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "participant_id", "in": "path", "description": "ID of the participant. You can fetch the participant ID using the add a participant API.", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_DeleteParticipant"}, "500": {"$ref": "#/components/responses/realtimekit_GenericError"}}, "security": [{"api_token": []}], "tags": ["Meetings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
