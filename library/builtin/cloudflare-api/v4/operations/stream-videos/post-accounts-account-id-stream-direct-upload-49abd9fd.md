---
title: Upload videos via direct upload URLs
page_id: operation-post-accounts-account-id-stream-direct-upload-457631f8
path: operations/stream-videos
description: Creates a direct upload that allows video uploads without an API key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/stream/direct_upload
operation_ids:
    - stream-videos-upload-videos-via-direct-upload-ur-ls
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload videos via direct upload URLs

`POST /accounts/{account_id}/stream/direct_upload`

Operation ID: `stream-videos-upload-videos-via-direct-upload-ur-ls`

Creates a direct upload that allows video uploads without an API key.

## Definition

```yaml
{"operationId": "stream-videos-upload-videos-via-direct-upload-ur-ls", "summary": "Upload videos via direct upload URLs", "description": "Creates a direct upload that allows video uploads without an API key.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_account_identifier"}}, {"name": "Upload-Creator", "in": "header", "schema": {"$ref": "#/components/schemas/stream_creator"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_direct_upload_request"}}}}, "responses": {"200": {"description": "Upload videos via direct upload URLs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_direct_upload_response"}}}}, "4XX": {"description": "Upload videos via direct upload URLs response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream Videos"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.direct-upload", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
