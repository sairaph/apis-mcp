---
title: List additional audio tracks on a video
page_id: operation-get-accounts-account-id-stream-identifier-audio-f2e700b3
path: operations/stream-audio-tracks
description: Lists additional audio tracks on a video. Note this API will not return information for audio attached to the video upload.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/audio
operation_ids:
    - list-audio-tracks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List additional audio tracks on a video

`GET /accounts/{account_id}/stream/{identifier}/audio`

Operation ID: `list-audio-tracks`

Lists additional audio tracks on a video. Note this API will not return information for audio attached to the video upload.

## Definition

```yaml
{"operationId": "list-audio-tracks", "summary": "List additional audio tracks on a video", "description": "Lists additional audio tracks on a video. Note this API will not return information for audio attached to the video upload.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}, {"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}], "responses": {"200": {"description": "Lists additional audio tracks on a video.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_listAudioTrackResponse"}}}}, "4XX": {"description": "Lists additional audio tracks on a video response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Audio Tracks"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.audio-tracks", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
