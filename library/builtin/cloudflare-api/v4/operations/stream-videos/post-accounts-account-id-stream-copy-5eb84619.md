---
title: Upload videos from a URL
page_id: operation-post-accounts-account-id-stream-copy-3255706c
path: operations/stream-videos
description: Uploads a video to Stream from a provided URL.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/copy
operation_ids:
    - stream-videos-upload-videos-from-a-url
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload videos from a URL

`POST /accounts/{account_id}/stream/copy`

Operation ID: `stream-videos-upload-videos-from-a-url`

Uploads a video to Stream from a provided URL.

## Definition

```yaml
{"operationId": "stream-videos-upload-videos-from-a-url", "summary": "Upload videos from a URL", "description": "Uploads a video to Stream from a provided URL.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}, {"name": "Upload-Creator", "in": "header", "schema": {"$ref": "#/components/schemas/stream_creator"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_video_copy_request"}}}}, "responses": {"200": {"description": "Upload videos from a URL response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_video_response_single"}}}}, "4XX": {"description": "Upload videos from a URL response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.copy", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
