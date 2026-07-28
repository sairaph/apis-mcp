---
title: Delete download
page_id: operation-delete-accounts-account-id-stream-identifier-downloads-download-type-f5464656
path: operations/stream-mp4-downloads
description: Delete specific type of download. For backwards-compatibility, DELETE requests to /downloads will delete the default download.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/stream/{identifier}/downloads/{download_type}
operation_ids:
    - stream-downloads-delete-type-specific-downloads
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete download

`DELETE /accounts/{account_id}/stream/{identifier}/downloads/{download_type}`

Operation ID: `stream-downloads-delete-type-specific-downloads`

Delete specific type of download. For backwards-compatibility, DELETE requests to /downloads will delete the default download.

## Definition

```yaml
{"operationId": "stream-downloads-delete-type-specific-downloads", "summary": "Delete download", "description": "Delete specific type of download. For backwards-compatibility, DELETE requests to /downloads will delete the default download.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_identifier-2"}}, {"name": "download_type", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/stream_download_type"}}], "responses": {"200": {"description": "Delete downloads response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_deleted_response"}}}}, "4XX": {"description": "Delete downloads response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/stream_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Stream MP4 Downloads"], "x-api-token-group": ["Stream Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "stream.typed-downloads", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
