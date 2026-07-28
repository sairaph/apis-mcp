---
title: List downloads
page_id: operation-get-accounts-account-id-stream-identifier-downloads-5a0e2cee
path: operations/stream-mp4-downloads
description: Lists the downloads created for a video.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/downloads
operation_ids:
    - stream-m-p-4-downloads-list-downloads
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List downloads

`GET /accounts/{account_id}/stream/{identifier}/downloads`

Operation ID: `stream-m-p-4-downloads-list-downloads`

Lists the downloads created for a video.

## Definition

```yaml
{"operationId": "stream-m-p-4-downloads-list-downloads", "summary": "List downloads", "description": "Lists the downloads created for a video.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}], "responses": {"200": {"description": "List downloads response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_downloads_response"}}}}, "4XX": {"description": "List downloads response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream MP4 Downloads"], "x-api-token-group": ["Stream Write", "Stream Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.downloads", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
