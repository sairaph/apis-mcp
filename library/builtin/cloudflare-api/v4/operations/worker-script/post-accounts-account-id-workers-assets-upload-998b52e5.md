---
title: Upload Assets
page_id: operation-post-accounts-account-id-workers-assets-upload-f0b12f5c
path: operations/worker-script
description: Upload assets ahead of creating a Worker version.  To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/assets/upload
operation_ids:
    - worker-assets-upload
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload Assets

`POST /accounts/{account_id}/workers/assets/upload`

Operation ID: `worker-assets-upload`

Upload assets ahead of creating a Worker version.  To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.

## Definition

```yaml
{"operationId": "worker-assets-upload", "summary": "Upload Assets", "description": "Upload assets ahead of creating a Worker version.  To learn more about the direct uploads of assets, see https://developers.cloudflare.com/workers/static-assets/direct-upload/.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}, {"name": "base64", "in": "query", "required": true, "schema": {"description": "Whether the file contents are base64-encoded. Must be `true`.", "type": "boolean", "enum": [true]}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"encoding": {"*": {"contentType": "*/*"}}, "schema": {"type": "object", "additionalProperties": {"description": "Base-64 encoded contents of the file. The content type of the file should be included to ensure a valid \"Content-Type\" header is included in asset responses.", "type": "string"}}}}}, "responses": {"201": {"description": "Upload Assets response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_completed-upload-assets-response"}}}}, "202": {"description": "Upload Assets response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_upload-assets-response"}}}}, "4XX": {"description": "Upload Assets response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"assets_jwt": []}], "tags": ["Worker Script"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.assets.upload", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
