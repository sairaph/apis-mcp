---
title: Edit video details
page_id: operation-post-accounts-account-id-stream-identifier-8e3caaad
path: operations/stream-videos
description: Edit details for a single video.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}
operation_ids:
    - stream-videos-update-video-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit video details

`POST /accounts/{account_id}/stream/{identifier}`

Operation ID: `stream-videos-update-video-details`

Edit details for a single video.

## Definition

```yaml
{"operationId": "stream-videos-update-video-details", "summary": "Edit video details", "description": "Edit details for a single video.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_video_update"}}}}, "responses": {"200": {"description": "Edit video details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_video_response_single"}}}}, "4XX": {"description": "Edit video details response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
