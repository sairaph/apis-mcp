---
title: Delete downloads
page_id: operation-delete-accounts-account-id-stream-identifier-downloads-2ec14a8a
path: operations/stream-mp4-downloads
description: Delete the downloads for a video. Use `/downloads/{download_type}` instead for type-specific downloads. Available types are `default` and `audio`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/downloads
operation_ids:
    - stream-m-p-4-downloads-delete-downloads
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete downloads

`DELETE /accounts/{account_id}/stream/{identifier}/downloads`

Operation ID: `stream-m-p-4-downloads-delete-downloads`

Delete the downloads for a video. Use `/downloads/{download_type}` instead for type-specific downloads. Available types are `default` and `audio`.

## Definition

```yaml
{"operationId": "stream-m-p-4-downloads-delete-downloads", "summary": "Delete downloads", "description": "Delete the downloads for a video. Use `/downloads/{download_type}` instead for type-specific downloads. Available types are `default` and `audio`.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "Delete downloads response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_deleted_response"}}}}, "4XX": {"description": "Delete downloads response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream MP4 Downloads"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.downloads", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
