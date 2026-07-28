---
title: Upsert asset hashes
page_id: operation-post-pages-assets-upsert-hashes-5bc04c49
path: operations/pages-assets
description: |-
    Register the provided file hashes as recently uploaded to the Pages
    asset store. Used as part of the Pages Direct Upload workflow so future
    deployments can avoid re-uploading files that are already present.

    Authenticate with the JWT obtained from the upload-token endpoint:
    GET /accounts/{account_id}/pages/projects/{project_name}/upload-token
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /pages/assets/upsert-hashes
operation_ids:
    - pages-assets-upsert-hashes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upsert asset hashes

`POST /pages/assets/upsert-hashes`

Operation ID: `pages-assets-upsert-hashes`

Register the provided file hashes as recently uploaded to the Pages
asset store. Used as part of the Pages Direct Upload workflow so future
deployments can avoid re-uploading files that are already present.

Authenticate with the JWT obtained from the upload-token endpoint:
GET /accounts/{account_id}/pages/projects/{project_name}/upload-token

## Definition

```yaml
{"operationId": "pages-assets-upsert-hashes", "summary": "Upsert asset hashes", "description": "Register the provided file hashes as recently uploaded to the Pages\nasset store. Used as part of the Pages Direct Upload workflow so future\ndeployments can avoid re-uploading files that are already present.\n\nAuthenticate with the JWT obtained from the upload-token endpoint:\nGET /accounts/{account_id}/pages/projects/{project_name}/upload-token\n", "requestBody": {"required": true, "content": {"application/json": {"example": {"hashes": ["a948904f2f0f479b8f936b8a0c5d9882", "b026324c6904b2a9cb4b88d6d61c81d1"]}, "schema": {"$ref": "#/components/schemas/pages_pages_assets_upsert_hashes_request"}}}}, "responses": {"200": {"description": "Upsert hashes response.", "content": {"application/json": {"example": {"success": true}, "schema": {"$ref": "#/components/schemas/pages_api-response-common"}}}}, "4XX": {"description": "Upsert hashes failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"pages_upload_token": []}], "tags": ["Pages Assets"], "x-forge-hidden": true}
```
