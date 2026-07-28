---
title: Start recording participant audio tracks
page_id: operation-post-accounts-account-id-realtime-kit-app-id-recordings-track-b281cd21
path: operations/recordings
description: Starts track recording for a meeting. Track recording currently records separate participant audio tracks as WebM files in the RealtimeKit bucket. Video track recording is in development. For more information, refer to [Track recording](/realtime/realtimekit/recording-guide/track-recording/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/realtime/kit/{app_id}/recordings/track
operation_ids:
    - startTrackRecordingForAMeeting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Start recording participant audio tracks

`POST /accounts/{account_id}/realtime/kit/{app_id}/recordings/track`

Operation ID: `startTrackRecordingForAMeeting`

Starts track recording for a meeting. Track recording currently records separate participant audio tracks as WebM files in the RealtimeKit bucket. Video track recording is in development. For more information, refer to [Track recording](/realtime/realtimekit/recording-guide/track-recording/).

## Definition

```yaml
{"operationId": "startTrackRecordingForAMeeting", "summary": "Start recording participant audio tracks", "description": "Starts track recording for a meeting. Track recording currently records separate participant audio tracks as WebM files in the RealtimeKit bucket. Video track recording is in development. For more information, refer to [Track recording](/realtime/realtimekit/recording-guide/track-recording/).", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_account_identifier"}}, {"name": "app_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/realtimekit_app_id"}}], "requestBody": {"$ref": "#/components/requestBodies/realtimekit_StartTrackRecordingBody"}, "responses": {"200": {"description": "OK", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/realtimekit_GenericSuccessResponse"}, {"properties": {"data": {"type": "object", "properties": {"recording": {"$ref": "#/components/schemas/realtimekit_Recording"}}, "required": ["recording"]}}, "type": "object"}]}}}}}, "security": [{"api_token": []}], "tags": ["Recordings"], "x-api-token-group": ["Realtime Admin", "Realtime"]}
```
