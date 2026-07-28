---
title: Fetch active livestreams for a meeting
page_id: operation-get-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-active-l-a0f0b378
path: operations/live-streams
description: Returns details of all active livestreams for the given meeting ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-livestream
operation_ids:
    - get-v2-meetings-meetingId-active-livestream
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch active livestreams for a meeting

`GET /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-livestream`

Operation ID: `get-v2-meetings-meetingId-active-livestream`

Returns details of all active livestreams for the given meeting ID.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meeting_id"}]
```

## Definition

```yaml
{"operationId": "get-v2-meetings-meetingId-active-livestream", "summary": "Fetch active livestreams for a meeting", "description": "Returns details of all active livestreams for the given meeting ID.", "parameters": [{"name": "meeting_id", "in": "path", "description": "ID of the meeting", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"example-1": {"value": {"data": {"created_at": "2023-07-15T11:58:08.404Z", "disabled": false, "id": "61e03cb4-d4da-42ad-9169-f53360f4f4b4", "meeting_id": "bbb03399-24e3-4b0c-90ec-172e97e636be", "name": null, "playback_url": "https://customer-s8oj0c1n5ek8ah1e.cloudflarestream.com/7de6a3fec0f9c05bf1df140950d3a237/manifest/video.m3u8", "status": "LIVE", "updated_at": "2023-07-15T11:58:24.278Z"}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"created_at": {"description": "Timestamp the object was created at. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "disabled": {"description": "Specifies if the livestream was disabled.", "type": "string"}, "id": {"description": "The livestream ID.", "type": "string"}, "ingest_server": {"description": "The server URL to which the RTMP encoder sends the video and audio data.", "type": "string"}, "meeting_id": {"type": "string"}, "name": {"description": "Name of the livestream.", "type": "string", "nullable": true}, "playback_url": {"description": "The web address that viewers can use to watch the livestream.", "type": "string"}, "status": {"type": "string", "enum": ["LIVE", "IDLE", "ERRORED", "INVOKED"]}, "stream_key": {"description": "Unique key for accessing each livestream.", "type": "string"}, "updated_at": {"description": "Timestamp the object was updated at. The time is returned in ISO format.", "type": "string", "format": "date-time"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
