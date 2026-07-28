---
title: Clip videos given a start and end time
page_id: operation-post-accounts-account-id-stream-clip-65887a9e
path: operations/stream-video-clipping
description: Clips a video based on the specified start and end times provided in seconds.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/clip
operation_ids:
    - stream-video-clipping-clip-videos-given-a-start-and-end-time
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Clip videos given a start and end time

`POST /accounts/{account_id}/stream/clip`

Operation ID: `stream-video-clipping-clip-videos-given-a-start-and-end-time`

Clips a video based on the specified start and end times provided in seconds.

## Definition

```yaml
{"operationId": "stream-video-clipping-clip-videos-given-a-start-and-end-time", "summary": "Clip videos given a start and end time", "description": "Clips a video based on the specified start and end times provided in seconds.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_videoClipStandard"}}}}, "responses": {"200": {"description": "Clip videos given a start and end time response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_clipResponseSingle"}}}}, "4XX": {"description": "Clip videos given a start and end time response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Video Clipping"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.clip", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
