---
title: Edit additional audio tracks on a video
page_id: operation-patch-accounts-account-id-stream-identifier-audio-audio-identifier-481107bc
path: operations/stream-audio-tracks
description: Edits additional audio tracks on a video. Editing the default status of an audio track to `true` will mark all other audio tracks on the video default status to `false`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/audio/{audio_identifier}
operation_ids:
    - edit-audio-tracks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit additional audio tracks on a video

`PATCH /accounts/{account_id}/stream/{identifier}/audio/{audio_identifier}`

Operation ID: `edit-audio-tracks`

Edits additional audio tracks on a video. Editing the default status of an audio track to `true` will mark all other audio tracks on the video default status to `false`.

## Definition

```yaml
{"operationId": "edit-audio-tracks", "summary": "Edit additional audio tracks on a video", "description": "Edits additional audio tracks on a video. Editing the default status of an audio track to `true` will mark all other audio tracks on the video default status to `false`.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "audio_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_audio_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_editAudioTrack"}}}}, "responses": {"200": {"description": "Edits additional audio tracks on a video.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_addAudioTrackResponse"}}}}, "4XX": {"description": "Edits additional audio tracks on a video response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Audio Tracks"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.audio-tracks", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
