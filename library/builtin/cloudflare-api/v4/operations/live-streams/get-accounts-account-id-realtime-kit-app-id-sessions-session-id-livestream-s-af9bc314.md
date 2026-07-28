---
title: Fetch livestream session details using a session ID
page_id: operation-get-accounts-account-id-realtime-kit-app-id-sessions-session-id-livestre-f3cfcddd
path: operations/live-streams
description: Returns livestream session details for the given session ID. Retreive the session ID using the `Fetch all sessions of an App` API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/livestream-sessions
operation_ids:
    - get-v2-livestreamsession-session-meetingId-active-livestream
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch livestream session details using a session ID

`GET /accounts/{account_id}/realtime/kit/{app_id}/sessions/{session_id}/livestream-sessions`

Operation ID: `get-v2-livestreamsession-session-meetingId-active-livestream`

Returns livestream session details for the given session ID. Retreive the session ID using the `Fetch all sessions of an App` API.

## Definition

```yaml
{"operationId": "get-v2-livestreamsession-session-meetingId-active-livestream", "summary": "Fetch livestream session details using a session ID", "description": "Returns livestream session details for the given session ID. Retreive the session ID using the `Fetch all sessions of an App` API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "schema": {"type": "number"}}, {"name": "page_no", "in": "query", "description": "The page number from which you want your page search results to be displayed.", "schema": {"type": "number"}}, {"name": "session_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"example-1": {"value": {"data": {"data": [{"created_at": "2023-07-15T07:08:40.742Z", "err_message": null, "id": "cf55a850-c3fd-46b5-b5d2-25b3b9aa1686", "ingest_seconds": 118, "invoked_time": null, "livestream_id": "b60ef3da-fbbb-441b-9168-8235b0035123", "started_time": "2023-07-15T07:08:39.000Z", "stopped_time": "2023-07-15T07:10:45.000Z", "updated_at": "2023-07-15T07:11:47.274Z", "viewer_seconds": 116.11599994199997}], "paging": {"end_offset": 1, "start_offset": 1, "total_count": 1}}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"created_at": {"description": "Timestamp the object was created at. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "err_message": {"type": "string", "enum": ["LIVE", "IDLE", "ERRORED", "INVOKED"]}, "id": {"description": "The livestream session ID.", "type": "string"}, "ingest_seconds": {"description": "The time duration for which the input was given or the meeting was streamed.", "type": "number"}, "invoked_time": {"description": "Name of the livestream.", "type": "string", "nullable": true}, "livestream_id": {"description": "The ID of the livestream.", "type": "string"}, "paging": {"type": "object", "properties": {"end_offset": {"type": "number"}, "start_offset": {"type": "number"}, "total_count": {"type": "number"}}}, "stopped_time": {"description": "Specifies if the livestream was disabled.", "type": "string", "format": "date-time"}, "updated_at": {"description": "Timestamp the object was updated at. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "viewer_seconds": {"description": "The total view time for which the viewers watched the stream.", "type": "number"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
