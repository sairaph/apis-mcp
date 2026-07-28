---
title: Fetch active recording
page_id: operation-get-accounts-account-id-realtime-kit-app-id-recordings-active-recording-4f576b7b
path: operations/recordings
description: Returns the active recording details for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/recordings/active-recording/{meeting_id}
operation_ids:
    - get_active_recording
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch active recording

`GET /accounts/{account_id}/realtime/kit/{app_id}/recordings/active-recording/{meeting_id}`

Operation ID: `get_active_recording`

Returns the active recording details for the given meeting ID.

## Definition

```yaml
{"operationId": "get_active_recording", "summary": "Fetch active recording", "description": "Returns the active recording details for the given meeting ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "meeting_id", "in": "path", "description": "ID of the meeting", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetActiveRecording"}, "404": {"$ref": "#/components/responses/realtimekit_GenericError"}}, "security": [{"api_token": []}], "tags": ["Recordings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
