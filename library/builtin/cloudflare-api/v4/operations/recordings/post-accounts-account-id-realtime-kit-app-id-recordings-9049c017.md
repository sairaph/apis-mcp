---
title: Start recording a meeting
page_id: operation-post-accounts-account-id-realtime-kit-app-id-recordings-aea24d8e
path: operations/recordings
description: Starts recording a meeting. The meeting can be started by an App admin directly, or a participant with permissions to start a recording, based on the type of authorization used.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/recordings
operation_ids:
    - start_recording
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start recording a meeting

`POST /accounts/{account_id}/realtime/kit/{app_id}/recordings`

Operation ID: `start_recording`

Starts recording a meeting. The meeting can be started by an App admin directly, or a participant with permissions to start a recording, based on the type of authorization used.

## Definition

```yaml
{"operationId": "start_recording", "summary": "Start recording a meeting", "description": "Starts recording a meeting. The meeting can be started by an App admin directly, or a participant with permissions to start a recording, based on the type of authorization used.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_StartRecording"}, "responses": {"200": {"$ref": "#/components/responses/realtimekit_GetRecording"}}, "security": [{"api_token": []}], "tags": ["Recordings"], "x-api-token-group": ["Realtime Admin", "Realtime"], "x-stability": "beta"}
```
