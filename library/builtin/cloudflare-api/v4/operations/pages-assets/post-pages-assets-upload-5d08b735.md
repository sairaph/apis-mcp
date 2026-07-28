---
title: Upload asset
page_id: operation-post-pages-assets-upload-7db93517
path: operations/pages-assets
description: |-
    Upload one or more files to the Pages asset store. Each file is
    identified by its content hash and is uploaded using the same JSON shape
    as the Cloudflare KV bulk write API. Used as part of the Pages Direct
    Upload workflow.

    Authenticate with the JWT obtained from the upload-token endpoint:
    GET /accounts/{account_id}/pages/projects/{project_name}/upload-token
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /pages/assets/upload
operation_ids:
    - pages-assets-upload
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload asset

`POST /pages/assets/upload`

Operation ID: `pages-assets-upload`

Upload one or more files to the Pages asset store. Each file is
identified by its content hash and is uploaded using the same JSON shape
as the Cloudflare KV bulk write API. Used as part of the Pages Direct
Upload workflow.

Authenticate with the JWT obtained from the upload-token endpoint:
GET /accounts/{account_id}/pages/projects/{project_name}/upload-token

## Definition

```yaml
{"operationId": "pages-assets-upload", "summary": "Upload asset", "description": "Upload one or more files to the Pages asset store. Each file is\nidentified by its content hash and is uploaded using the same JSON shape\nas the Cloudflare KV bulk write API. Used as part of the Pages Direct\nUpload workflow.\n\nAuthenticate with the JWT obtained from the upload-token endpoint:\nGET /accounts/{account_id}/pages/projects/{project_name}/upload-token\n", "requestBody": {"required": true, "content": {"application/json": {"example": [{"base64": true, "key": "b026324c6904b2a9cb4b88d6d61c81d1", "metadata": {"contentType": "text/plain"}, "value": "SGVsbG8sIFdvcmxkIQ=="}], "schema": {"$ref": "#/components/schemas/pages_pages_assets_upload_request"}}}}, "responses": {"200": {"description": "Upload asset response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common"}}}}, "4XX": {"description": "Upload asset failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"pages_upload_token": []}], "tags": ["Pages Assets"], "x-forge-hidden": true}
```
