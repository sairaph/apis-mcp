---
title: Add audio tracks to a video
page_id: operation-post-accounts-account-id-stream-identifier-audio-copy-123fccaa
path: operations/stream-audio-tracks
description: Adds an additional audio track to a video using the provided audio track URL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/audio/copy
operation_ids:
    - add-audio-track
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add audio tracks to a video

`POST /accounts/{account_id}/stream/{identifier}/audio/copy`

Operation ID: `add-audio-track`

Adds an additional audio track to a video using the provided audio track URL.

## Definition

```yaml
{"operationId": "add-audio-track", "summary": "Add audio tracks to a video", "description": "Adds an additional audio track to a video using the provided audio track URL.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_copyAudioTrack"}}}}, "responses": {"200": {"description": "Add audio tracks to a video.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_addAudioTrackResponse"}}}}, "4XX": {"description": "Add audio tracks to a video response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Audio Tracks"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.audio-tracks", "x-fern-sdk-method-name": "copy", "x-forge-hidden": true}
```
