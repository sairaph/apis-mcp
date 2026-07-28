---
title: Fetch livestream session details for a meeting
page_id: operation-get-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-livestre-fa9f80a9
path: operations/live-streams
description: Returns livestream session details for the given meeting ID. Retreive the meeting ID using the `Create a meeting` API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/livestream
operation_ids:
    - livestream-session-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch livestream session details for a meeting

`GET /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/livestream`

Operation ID: `livestream-session-details`

Returns livestream session details for the given meeting ID. Retreive the meeting ID using the `Create a meeting` API.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meeting_id"}]
```

## Definition

```yaml
{"operationId": "livestream-session-details", "summary": "Fetch livestream session details for a meeting", "description": "Returns livestream session details for the given meeting ID. Retreive the meeting ID using the `Create a meeting` API.", "parameters": [{"name": "page_no", "in": "query", "description": "The page number from which you want your page search results to be displayed.", "schema": {"type": "integer"}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "schema": {"type": "integer"}}, {"name": "meeting_id", "in": "path", "description": "ID of the meeting", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"example-0": {"summary": "Get LiveStream From StreamId Success", "value": {"data": {"livestream": {"created_at": "2023-07-15T07:08:26.634Z", "disabled": false, "id": "b60ef3da-fbbb-441b-9168-8235b0035248", "ingest_server": "rtmps://live.cloudflare.com:443/live/", "meeting_id": "aaa87945-df1c-49d9-8159-262a7ae6cfde", "name": null, "playback_url": "https://customer-s8oj0c1n5ek8ah1e.cloudflarestream.com/7de6a3fec0f9c05bf1df140950d3a237/manifest/video.m3u8", "status": "LIVE", "stream_key": "f26566285faca6fbe2e79a73a66rsrrsrrsr3cde23a2bb7dbc6c2c1761b98f4e4", "updated_at": "2023-07-15T07:08:40.726Z"}, "sessions": [{"created_at": "2023-07-15T07:08:40.742Z", "err_message": null, "id": "cf55a850-c3fd-46b5-b5d2-25b3b9aa1686", "ingest_seconds": 60, "invoked_time": null, "livestream_id": "b60ef3da-fbbb-441b-9168-8235b0035248", "started_time": "2023-07-15T07:08:39.000Z", "stopped_time": null, "updated_at": "2023-07-15T07:08:40.742Z", "viewer_seconds": 60}]}, "paging": {"end_offset": 1, "start_offset": 1, "total_count": 1}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"livestreams": {"type": "array", "items": {"properties": {"created_at": {"description": "The timestamp at which the livestream was created. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "disabled": {"description": "Specifies if the livestream was disabled.", "type": "boolean"}, "id": {"description": "The livestream ID.", "type": "string"}, "ingest_server": {"description": "The server URL to which the RTMP encoder sends the video and audio data.", "type": "string"}, "meeting_id": {"description": "The ID of the meeting that was livestreamed.", "type": "string"}, "name": {"description": "Name of the livestream.", "type": "string", "nullable": true}, "playback_url": {"description": "The web address that viewers can use to watch the livestream.", "type": "string"}, "status": {"type": "string", "enum": ["LIVE", "INVOKED", "ERRORED", "IDLE"]}, "stream_key": {"description": "Unique key for accessing each livestream.", "type": "string"}, "updated_at": {"description": "The timestamp at which the livestream was updated. The time is returned in ISO format.", "type": "string", "format": "date-time"}}, "type": "object"}}, "paging": {"type": "object", "properties": {"end_offset": {"type": "integer", "example": 1}, "start_offset": {"type": "integer", "example": 1}, "total_count": {"type": "integer", "example": 1}}}, "sessions": {"type": "object", "properties": {"created_at": {"description": "The timestamp at which the livestream was created. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "err_message": {"type": "string"}, "id": {"description": "The ID of the livestream session.", "type": "string"}, "ingest_seconds": {"description": "The time duration for which the input was given or the meeting was streamed.", "type": "string"}, "invoked_time": {"description": "The time at which the livestream was invoked.", "type": "string", "format": "date-time"}, "livestream_id": {"description": "The ID of the livestream.", "type": "string"}, "started_time": {"description": "The time at which the livestream was started.", "type": "string", "format": "date-time"}, "stopped_time": {"description": "The time at which the livestream was stopped.", "type": "string", "format": "date-time"}, "updated_at": {"description": "The timestamp at which the livestream was updated. The time is returned in ISO format.", "type": "string", "format": "date-time"}}}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
