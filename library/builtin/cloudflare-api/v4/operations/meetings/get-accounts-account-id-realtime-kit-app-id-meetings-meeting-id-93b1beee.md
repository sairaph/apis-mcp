---
title: Fetch a meeting for an App
page_id: operation-get-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-750ff94b
path: operations/meetings
description: Returns a meeting details in an App for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}
operation_ids:
    - get_meeting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch a meeting for an App

`GET /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}`

Operation ID: `get_meeting`

Returns a meeting details in an App for the given meeting ID.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meeting_id"}]
```

## Definition

```yaml
{"operationId": "get_meeting", "summary": "Fetch a meeting for an App", "description": "Returns a meeting details in an App for the given meeting ID.", "parameters": [{"name": "name", "in": "query", "schema": {"type": "string"}}, {"name": "meeting_id", "in": "path", "description": "ID of the meeting. Fetch the meeting ID using the create a meeting API.", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetMeeting"}, "500": {"$ref": "#/components/responses/realtimekit_GenericError"}}, "security": [{"api_token": []}], "tags": ["Meetings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
