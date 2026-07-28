---
title: Fetch livestream session details using livestream session ID
page_id: operation-get-accounts-account-id-realtime-kit-app-id-livestreams-sessions-livestr-7243c1ca
path: operations/live-streams
description: Returns livestream session details for the given livestream session ID. Retrieve the `livestream_session_id`using the `Fetch livestream session details using a session ID` API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/livestreams/sessions/{livestream-session-id}
operation_ids:
    - get-v2-livestreams-livestream-session-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch livestream session details using livestream session ID

`GET /accounts/{account_id}/realtime/kit/{app_id}/livestreams/sessions/{livestream-session-id}`

Operation ID: `get-v2-livestreams-livestream-session-id`

Returns livestream session details for the given livestream session ID. Retrieve the `livestream_session_id`using the `Fetch livestream session details using a session ID` API.

## Definition

```yaml
{"operationId": "get-v2-livestreams-livestream-session-id", "summary": "Fetch livestream session details using livestream session ID", "description": "Returns livestream session details for the given livestream session ID. Retrieve the `livestream_session_id`using the `Fetch livestream session details using a session ID` API.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"name": "livestream-session-id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"example-1": {"value": {"data": {"created_at": "2023-07-15T07:08:40.742Z", "err_message": null, "id": "cf55a850-c3fd-46b5-b5d2-25b3b9aa1686", "ingest_seconds": 60, "invoked_time": null, "livestream_id": "b60ef3da-fbbb-441b-9168-8235b0035248", "started_time": "2023-07-15T07:08:39.000Z", "stopped_time": null, "updated_at": "2023-07-15T07:08:40.742Z", "viewer_seconds": 60}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"created_at": {"description": "Timestamp the object was created at. The time is returned in ISO format.", "type": "string", "format": "date-time"}, "err_message": {"description": "The server URL to which the RTMP encoder sends the video and audio data.", "type": "string"}, "id": {"description": "The livestream ID.", "type": "string"}, "ingest_seconds": {"description": "Name of the livestream.", "type": "integer"}, "livestream_id": {"type": "string"}, "started_time": {"description": "Unique key for accessing each livestream.", "type": "string"}, "stopped_time": {"description": "The web address that viewers can use to watch the livestream.", "type": "string"}, "updated_at": {"description": "Timestamp the object was updated at. The time is returned in ISO format.", "type": "string"}, "viewer_seconds": {"description": "Specifies if the livestream was disabled.", "type": "integer"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
