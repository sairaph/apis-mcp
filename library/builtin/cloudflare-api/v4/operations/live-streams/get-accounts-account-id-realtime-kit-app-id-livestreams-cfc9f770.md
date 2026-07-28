---
title: Fetch all livestreams
page_id: operation-get-accounts-account-id-realtime-kit-app-id-livestreams-21129be2
path: operations/live-streams
description: Returns details of livestreams associated with the given App ID. It includes livestreams created by your App and RealtimeKit meetings that are livestreamed by your App. If you only want details of livestreams created by your App and not RealtimeKit meetings, you can use the `exclude_meetings` query parameter.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/livestreams
operation_ids:
    - fetch_all_livestreams
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch all livestreams

`GET /accounts/{account_id}/realtime/kit/{app_id}/livestreams`

Operation ID: `fetch_all_livestreams`

Returns details of livestreams associated with the given App ID. It includes livestreams created by your App and RealtimeKit meetings that are livestreamed by your App. If you only want details of livestreams created by your App and not RealtimeKit meetings, you can use the `exclude_meetings` query parameter.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}]
```

## Definition

```yaml
{"operationId": "fetch_all_livestreams", "summary": "Fetch all livestreams", "description": "Returns details of livestreams associated with the given App ID. It includes livestreams created by your App and RealtimeKit meetings that are livestreamed by your App. If you only want details of livestreams created by your App and not RealtimeKit meetings, you can use the `exclude_meetings` query parameter.", "parameters": [{"name": "exclude_meetings", "in": "query", "description": "Exclude the RealtimeKit meetings that are livestreamed.", "schema": {"type": "boolean", "default": false}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "schema": {"type": "integer"}}, {"name": "page_no", "in": "query", "description": "The page number from which you want your page search results to be displayed.", "schema": {"type": "integer"}}, {"name": "status", "in": "query", "description": "Specifies the status of the operation.", "schema": {"type": "string", "enum": ["LIVE", "IDLE", "ERRORED", "INVOKED"]}}, {"name": "start_time", "in": "query", "description": "Specify the start time range in ISO format to access the live stream.", "schema": {"type": "string", "format": "date-time"}}, {"name": "end_time", "in": "query", "description": "Specify the end time range in ISO format to access the live stream.", "schema": {"type": "string", "format": "date-time"}}, {"name": "sort_order", "in": "query", "description": "Specifies the sorting order for the results.", "schema": {"type": "string", "enum": ["ASC", "DSC"]}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"example-0": {"summary": "Get All Live streams Of An Org Staging Success", "value": {"data": [{"created_at": "2023-07-15T11:48:34.753Z", "disabled": false, "id": "3fd739f4-3c41-456e-bfba-6ebd51e16d2c", "ingest_server": "rtmps://live.cloudflare.com:443/live/", "meeting_id": null, "name": null, "playback_url": "https://customer-s8oj0c1n5ek8ah1e.cloudflarestream.com/7de6a3fec0f9c05bf1df140950d3a237/manifest/video.m3u8", "status": "IDLE", "stream_key": "f26566285faca6fbe2e79a73a66rsrrsrrsr3cde23a2bb7dbc6c2c1761b98f4e4", "updated_at": "2023-07-15T11:48:34.753Z"}], "paging": {"end_offset": 1, "start_offset": 1, "total_count": 1}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"created_at": {"description": "Timestamp the object was created at. The time is returned in ISO format.", "type": "string", "format": "date-time", "example": "2023-07-15T11:48:34.753Z"}, "disabled": {"description": "Specifies if the livestream was disabled.", "type": "string"}, "id": {"description": "The ID of the livestream.", "type": "string", "format": "uuid", "example": "3fd739f4-3c41-456e-bfba-6ebd51e16d2d"}, "ingest_server": {"description": "The server URL to which the RTMP encoder sends the video and audio data.", "type": "string", "example": "rtmps://live.cloudflare.com:443/live/"}, "meeting_id": {"description": "ID of the meeting.", "type": "string"}, "name": {"description": "Name of the livestream.", "type": "string", "example": "test"}, "paging": {"type": "object", "properties": {"end_offset": {"type": "integer", "example": 1}, "start_offset": {"type": "integer", "example": 1}, "total_count": {"type": "integer", "example": 1}}}, "playback_url": {"description": "The web address that viewers can use to watch the livestream.", "type": "string", "example": "https://customer-s8oj0c1n5ek8ah1e.cloudflarestream.com/7de6a3fec0f9c05bf1df140950d3a237/manifest/video.m3u8"}, "status": {"type": "string", "enum": ["LIVE", "IDLE", "ERRORED", "INVOKED"]}, "stream_key": {"description": "Unique key for accessing each livestream.", "type": "string", "example": "f26566285faca6fbe2e79a73a66rsrrsrrsr3cde23a2bb7dbc6c2c1761b98f4e4"}, "updated_at": {"description": "Timestamp the object was updated at. The time is returned in ISO format.", "type": "string", "format": "date-time", "example": "2023-07-15T11:48:34.753Z"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
