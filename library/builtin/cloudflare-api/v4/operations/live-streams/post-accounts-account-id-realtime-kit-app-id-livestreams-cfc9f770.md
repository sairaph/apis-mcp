---
title: Create an independent livestream
page_id: operation-post-accounts-account-id-realtime-kit-app-id-livestreams-9190570e
path: operations/live-streams
description: Creates a livestream for the given App ID and returns ingest server, stream key, and playback URL. You can pass custom input to the ingest server and stream key, and freely distribute the content using the playback URL on any player that supports HLS/LHLS.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/livestreams
operation_ids:
    - create_livestream
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an independent livestream

`POST /accounts/{account_id}/realtime/kit/{app_id}/livestreams`

Operation ID: `create_livestream`

Creates a livestream for the given App ID and returns ingest server, stream key, and playback URL. You can pass custom input to the ingest server and stream key, and freely distribute the content using the playback URL on any player that supports HLS/LHLS.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}]
```

## Definition

```yaml
{"operationId": "create_livestream", "summary": "Create an independent livestream", "description": "Creates a livestream for the given App ID and returns ingest server, stream key, and playback URL. You can pass custom input to the ingest server and stream key, and freely distribute the content using the playback URL on any player that supports HLS/LHLS.", "requestBody": {"required": true, "content": {"application/json": {"examples": {"example-1": {"value": {"name": "prdmmp-xhycsl"}}}, "schema": {"type": "object", "properties": {"name": {"description": "Name of the livestream", "type": "string", "nullable": true, "pattern": "^[a-zA-Z0-9-_]*$"}}}}}}, "responses": {"201": {"description": "Successful response", "content": {"application/json": {"examples": {"example-1": {"value": {"data": {"disabled": false, "id": "78dd0b50-4147-4bb8-88d3-2ccc2e98bff0", "ingest_server": "rtmps://live.cloudflare.com:443/live/", "meeting_id": null, "name": "Livestreaming-Demo", "playback_url": "https://customer-s8oj0c1n5ek8ah1e.cloudflarestream.com/7de6a3fec0f9c05bf1df140950d3a237/manifest/video.m3u8", "status": "INVOKED", "stream_key": "f26566285faca6fbe2e79a73a66rsrrsrrsr3cde23a2bb7dbc6c2c1761b98f4e4"}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"disabled": {"description": "Specifies if the livestream was disabled.", "type": "boolean"}, "id": {"description": "The livestream ID.", "type": "string"}, "ingest_server": {"description": "The server URL to which the RTMP encoder should send the video and audio data.", "type": "string"}, "meeting_id": {"type": "string", "nullable": true}, "name": {"type": "string"}, "playback_url": {"description": "The web address that viewers can use to watch the livestream.", "type": "string"}, "status": {"type": "string", "enum": ["LIVE", "IDLE", "ERRORED", "INVOKED"]}, "stream_key": {"description": "Unique key for accessing each livestream.", "type": "string"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
