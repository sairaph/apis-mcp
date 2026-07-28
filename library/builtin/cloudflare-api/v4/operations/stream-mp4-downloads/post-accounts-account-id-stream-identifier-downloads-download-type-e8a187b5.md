---
title: Create download
page_id: operation-post-accounts-account-id-stream-identifier-downloads-download-type-bbc10499
path: operations/stream-mp4-downloads
description: Creates a download for a video of specified type. For backwards-compatibility, POST requests to /downloads will enable the default download.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/downloads/{download_type}
operation_ids:
    - stream-downloads-create-type-specific-downloads
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create download

`POST /accounts/{account_id}/stream/{identifier}/downloads/{download_type}`

Operation ID: `stream-downloads-create-type-specific-downloads`

Creates a download for a video of specified type. For backwards-compatibility, POST requests to /downloads will enable the default download.

## Definition

```yaml
{"operationId": "stream-downloads-create-type-specific-downloads", "summary": "Create download", "description": "Creates a download for a video of specified type. For backwards-compatibility, POST requests to /downloads will enable the default download.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}, {"name": "download_type", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_download_type"}}], "responses": {"200": {"description": "Create download of specified type response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_downloads_response"}}}}, "4XX": {"description": "Create downloads of specified type response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream MP4 Downloads"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.typed-downloads", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
