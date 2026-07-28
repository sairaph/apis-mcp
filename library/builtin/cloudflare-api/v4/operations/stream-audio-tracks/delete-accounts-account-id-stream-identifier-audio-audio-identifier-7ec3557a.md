---
title: Delete additional audio tracks on a video
page_id: operation-delete-accounts-account-id-stream-identifier-audio-audio-identifier-4d1322f2
path: operations/stream-audio-tracks
description: Deletes additional audio tracks on a video. Deleting a default audio track is not allowed. You must assign another audio track as default prior to deletion.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/audio/{audio_identifier}
operation_ids:
    - delete-audio-tracks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete additional audio tracks on a video

`DELETE /accounts/{account_id}/stream/{identifier}/audio/{audio_identifier}`

Operation ID: `delete-audio-tracks`

Deletes additional audio tracks on a video. Deleting a default audio track is not allowed. You must assign another audio track as default prior to deletion.

## Definition

```yaml
{"operationId": "delete-audio-tracks", "summary": "Delete additional audio tracks on a video", "description": "Deletes additional audio tracks on a video. Deleting a default audio track is not allowed. You must assign another audio track as default prior to deletion.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "audio_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_audio_identifier"}}], "responses": {"200": {"description": "Deletes additional audio tracks on a video.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_deleted_response"}}}}, "4XX": {"description": "Deletes additional audio tracks on a video response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_deleted_response"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Audio Tracks"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.audio-tracks", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
