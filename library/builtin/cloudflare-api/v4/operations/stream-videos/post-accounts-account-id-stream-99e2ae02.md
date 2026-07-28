---
title: Initiate video uploads using TUS
page_id: operation-post-accounts-account-id-stream-d7afeee9
path: operations/stream-videos
description: Initiates a video upload using the TUS protocol. On success, the server responds with a status code 201 (created) and includes a `location` header to indicate where the content should be uploaded. Refer to https://tus.io for protocol details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream
operation_ids:
    - stream-videos-initiate-video-uploads-using-tus
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Initiate video uploads using TUS

`POST /accounts/{account_id}/stream`

Operation ID: `stream-videos-initiate-video-uploads-using-tus`

Initiates a video upload using the TUS protocol. On success, the server responds with a status code 201 (created) and includes a `location` header to indicate where the content should be uploaded. Refer to https://tus.io for protocol details.

## Definition

```yaml
{"operationId": "stream-videos-initiate-video-uploads-using-tus", "summary": "Initiate video uploads using TUS", "description": "Initiates a video upload using the TUS protocol. On success, the server responds with a status code 201 (created) and includes a `location` header to indicate where the content should be uploaded. Refer to https://tus.io for protocol details.", "parameters": [{"name": "Tus-Resumable", "in": "header", "required": true, "schema": {"$ref": "#/components/schemas/stream_tus_resumable"}}, {"name": "Upload-Creator", "in": "header", "schema": {"$ref": "#/components/schemas/stream_creator"}}, {"name": "Upload-Length", "in": "header", "required": true, "schema": {"$ref": "#/components/schemas/stream_upload_length"}}, {"name": "Upload-Metadata", "in": "header", "schema": {"$ref": "#/components/schemas/stream_upload_metadata"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}, {"name": "direct_user", "in": "query", "schema": {"$ref": "#/components/schemas/stream_direct_user"}}], "responses": {"201": {"description": "Initiate video uploads using TUS response. Returns a 201 Created status with Location header for TUS uploads.", "headers": {"Location": {"description": "The URL where the upload should be sent.", "schema": {"type": "string"}}, "Stream-Media-ID": {"description": "The unique identifier for the video being uploaded.", "schema": {"type": "string"}}, "Tus-Resumable": {"description": "The TUS protocol version.", "schema": {"type": "string", "example": "1.0.0"}}}}, "4XX": {"description": "Initiate video uploads using TUS response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
