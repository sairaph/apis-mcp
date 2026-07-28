---
title: Stop livestreaming a meeting
page_id: operation-post-accounts-account-id-realtime-kit-app-id-meetings-meeting-id-active-de55066c
path: operations/live-streams
description: Stops the active livestream of a meeting associated with the given meeting ID. Retreive the meeting ID using the `Create a meeting` API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-livestream/stop
operation_ids:
    - stop_livestreaming
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Stop livestreaming a meeting

`POST /accounts/{account_id}/realtime/kit/{app_id}/meetings/{meeting_id}/active-livestream/stop`

Operation ID: `stop_livestreaming`

Stops the active livestream of a meeting associated with the given meeting ID. Retreive the meeting ID using the `Create a meeting` API.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}, {"$ref": "#/components/parameters/realtimekit_meeting_id"}]
```

## Definition

```yaml
{"operationId": "stop_livestreaming", "summary": "Stop livestreaming a meeting", "description": "Stops the active livestream of a meeting associated with the given meeting ID. Retreive the meeting ID using the `Create a meeting` API.", "parameters": [{"name": "meeting_id", "in": "path", "description": "ID of the meeting", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"content": {}}, "responses": {"200": {"description": "OK", "content": {"application/json": {"examples": {"example-1": {"value": {"data": {"message": "Stopped live stream successfully"}, "success": true}}}, "schema": {"type": "object", "properties": {"data": {"type": "object", "properties": {"message": {"type": "string"}}}, "success": {"type": "boolean"}}}}}}}, "security": [{"api_token": []}], "tags": ["Live streams"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
