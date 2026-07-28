---
title: Fetch livestream details using livestream ID
page_id: operation-get-accounts-account-id-realtime-kit-app-id-livestreams-livestream-id-dfb31750
path: operations/live-streams
description: Returns details of a livestream with sessions for the given livestream ID. Retreive the livestream ID using the `Start livestreaming a meeting` API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/livestreams/{livestream_id}
operation_ids:
    - get-v2-livestream-session-livestream-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch livestream details using livestream ID

`GET /accounts/{account_id}/realtime/kit/{app_id}/livestreams/{livestream_id}`

Operation ID: `get-v2-livestream-session-livestream-id`

Returns details of a livestream with sessions for the given livestream ID. Retreive the livestream ID using the `Start livestreaming a meeting` API.

## Definition

```yaml
{"operationId": "get-v2-livestream-session-livestream-id", "summary": "Fetch livestream details using livestream ID", "description": "Returns details of a livestream with sessions for the given livestream ID. Retreive the livestream ID using the `Start livestreaming a meeting` API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "page_no", "in": "query", "description": "The page number from which you want your page search results to be displayed.", "schema": {"type": "integer"}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "schema": {"type": "integer"}}, {"name": "livestream_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"example-1": {"value": {"data": {"livestream": {"created_at": "2023-07-15T07:08:26.634Z", "disabled": false, "id": "b60ef3da-fbbb-441b-9168-8235b0035248", "ingest_server": "rtmps://live.cloudflare.com:443/live/", "meeting_id": "bbb87945-df1c-49d9-8159-262a7ae6cfde", "name": null, "playback_url": "https://customer-s8oj0c1n5ek8ah1e.cloudflarestream.com/7de6a3fec0f9c05bf1df140950d3a237/manifest/video.m3u8", "status": "LIVE", "stream_key": "f26566285faca6fbe2e79a73a66rsrrsrrsr3cde23a2bb7dbc6c2c1761b98f4e4", "updated_at": "2023-07-15T07:08:40.726Z"}, "sessions": [{"created_at": "2023-07-15T07:08:40.742Z", "err_message": null, "id": "cf55a850-c3fd-46b5-b5d2-25b3b9aa1686", "ingest_seconds": 60, "invoked_time": null, "livestream_id": "b60ef3da-fbbb-441b-9168-8235b0035248", "started_time": "2023-07-15T07:08:39.000Z", "stopped_time": null, "updated_at": "2023-07-15T07:08:40.742Z", "viewer_seconds": 60}]}, "paging": {"end_offset": 1, "start_offset": 1, "total_count": 1}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"livestream": {"type": "object", "properties": {"created_at": {"description": "Timestamp the object was created at. The time is returned in ISO format.", "type": "string"}, "disabled": {"description": "Specifies if the livestream was disabled.", "type": "string"}, "id": {"description": "ID of the livestream.", "type": "string"}, "ingest_server": {"description": "The server URL to which the RTMP encoder sends the video and audio data.", "type": "string"}, "meeting_id": {"description": "The ID of the meeting.", "type": "string"}, "name": {"description": "Name of the livestream.", "type": "string"}, "playback_url": {"description": "The web address that viewers can use to watch the livestream.", "type": "string"}, "status": {"type": "string", "enum": ["LIVE", "IDLE", "ERRORED", "INVOKED"]}, "stream_key": {"description": "Unique key for accessing each livestream.", "type": "string"}, "updated_at": {"description": "Timestamp the object was updated at. The time is returned in ISO format.", "type": "string"}}}, "paging": {"type": "object", "properties": {"end_offset": {"type": "integer", "example": 1}, "start_offset": {"type": "integer", "example": 1}, "total_count": {"type": "integer", "example": 1}}}, "session": {"type": "object", "properties": {"created_at": {"description": "Timestamp the object was created at. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "err_message": {"type": "string"}, "id": {"description": "ID of the session.", "type": "string"}, "ingest_seconds": {"description": "The time duration for which the input was given or the meeting was streamed.", "type": "number"}, "invoked_time": {"description": "Timestamp the object was invoked. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "livestream_id": {"type": "string"}, "started_time": {"description": "Timestamp the object was started. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "stopped_time": {"description": "Timestamp the object was stopped. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "updated_at": {"description": "Timestamp the object was updated at. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "viewer_seconds": {"description": "The total view time for which the viewers watched the stream.", "type": "number"}}}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
