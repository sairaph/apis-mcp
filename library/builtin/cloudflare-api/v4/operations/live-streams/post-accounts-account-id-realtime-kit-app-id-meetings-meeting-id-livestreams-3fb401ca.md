---
title: Start livestreaming a meeting
page_id: operation-post-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-livestr-7c8664b2
path: operations/live-streams
description: Starts livestream of a meeting associated with the given meeting ID. Retreive the meeting ID using the `Create a meeting` API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/livestreams
operation_ids:
    - start-livestreaming
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start livestreaming a meeting

`POST /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/livestreams`

Operation ID: `start-livestreaming`

Starts livestream of a meeting associated with the given meeting ID. Retreive the meeting ID using the `Create a meeting` API.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meeting_id"}]
```

## Definition

```yaml
{"operationId": "start-livestreaming", "summary": "Start livestreaming a meeting", "description": "Starts livestream of a meeting associated with the given meeting ID. Retreive the meeting ID using the `Create a meeting` API.", "parameters": [{"name": "meeting_id", "in": "path", "description": "ID of the meeting", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"example-1": {"value": {"name": "prdmmp-xhycsl"}}}, "schema": {"type": "object", "properties": {"name": {"type": "string", "nullable": true, "pattern": "^[a-zA-Z0-9-_]*$"}, "video_config": {"type": "object", "properties": {"height": {"description": "Height of the livestreaming video in pixels", "type": "integer"}, "width": {"description": "Width of the livestreaming video in pixels", "type": "integer"}}}}}}}}, "responses": {"201": {"description": "Created", "content": {"application/json": {"examples": {"example-1": {"value": {"data": {"id": "7088bba8-f522-49a8-b59b-3cd0e946bbb0", "ingest_server": "rtmps://live.cloudflare.com:443/live/", "playback_url": "https://customer-s8oj0c1n5ek8ah1e.cloudflarestream.com/7de6a3fec0f9c05bf1df140950d3a237/manifest/video.m3u8", "status": "INVOKED", "stream_key": "f26566285faca6fbe2e79a73a66rsrrsrrsr3cde23a2bb7dbc6c2c1761b98f4e4"}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"id": {"description": "The livestream ID.", "type": "string"}, "ingest_server": {"description": "The server URL to which the RTMP encoder sends the video and audio data.", "type": "string"}, "playback_url": {"description": "The web address that viewers can use to watch the livestream.", "type": "string"}, "status": {"type": "string", "enum": ["LIVE", "IDLE", "ERRORED", "INVOKED"]}, "stream_key": {"description": "Unique key for accessing each livestream.", "type": "string"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
