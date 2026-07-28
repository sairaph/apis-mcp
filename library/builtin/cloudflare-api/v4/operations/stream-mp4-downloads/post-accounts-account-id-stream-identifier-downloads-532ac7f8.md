---
title: Create downloads
page_id: operation-post-accounts-account-id-stream-identifier-downloads-e57949ef
path: operations/stream-mp4-downloads
description: Creates a download for a video when a video is ready to view. Use `/downloads/{download_type}` instead for type-specific downloads. Available types are `default` and `audio`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/downloads
operation_ids:
    - stream-m-p-4-downloads-create-downloads
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create downloads

`POST /accounts/{account_id}/stream/{identifier}/downloads`

Operation ID: `stream-m-p-4-downloads-create-downloads`

Creates a download for a video when a video is ready to view. Use `/downloads/{download_type}` instead for type-specific downloads. Available types are `default` and `audio`.

## Definition

```yaml
{"operationId": "stream-m-p-4-downloads-create-downloads", "summary": "Create downloads", "description": "Creates a download for a video when a video is ready to view. Use `/downloads/{download_type}` instead for type-specific downloads. Available types are `default` and `audio`.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "Create downloads response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_downloads_response"}}}}, "4XX": {"description": "Create downloads response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream MP4 Downloads"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.downloads", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
