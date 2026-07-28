---
title: Replace a meeting
page_id: operation-put-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-2f4a0dfd
path: operations/meetings
description: Replaces all the details for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}
operation_ids:
    - replace_meeting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace a meeting

`PUT /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}`

Operation ID: `replace_meeting`

Replaces all the details for the given meeting ID.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meeting_id"}]
```

## Definition

```yaml
{"operationId": "replace_meeting", "summary": "Replace a meeting", "description": "Replaces all the details for the given meeting ID.", "parameters": [{"name": "meeting_id", "in": "path", "description": "ID of the meeting. Fetch the meeting ID using the create a meeting API.", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_CreateMeetingBody"}, "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetMeeting"}}, "security": [{"api_token": []}], "tags": ["Meetings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
