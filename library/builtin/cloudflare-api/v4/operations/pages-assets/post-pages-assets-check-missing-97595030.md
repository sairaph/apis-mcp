---
title: Check missing assets
page_id: operation-post-pages-assets-check-missing-acae3b74
path: operations/pages-assets
description: |-
    Check which of the provided file hashes are missing from the Pages
    asset store. Returns a list of missing hashes that need to be uploaded.
    Used as part of the Pages Direct Upload workflow.

    Authenticate with the JWT obtained from the upload-token endpoint:
    GET /accounts/{account_id}/pages/projects/{project_name}/upload-token
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /pages/assets/check-missing
operation_ids:
    - pages-assets-check-missing
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Check missing assets

`POST /pages/assets/check-missing`

Operation ID: `pages-assets-check-missing`

Check which of the provided file hashes are missing from the Pages
asset store. Returns a list of missing hashes that need to be uploaded.
Used as part of the Pages Direct Upload workflow.

Authenticate with the JWT obtained from the upload-token endpoint:
GET /accounts/{account_id}/pages/projects/{project_name}/upload-token

## Definition

```yaml
{"operationId": "pages-assets-check-missing", "summary": "Check missing assets", "description": "Check which of the provided file hashes are missing from the Pages\nasset store. Returns a list of missing hashes that need to be uploaded.\nUsed as part of the Pages Direct Upload workflow.\n\nAuthenticate with the JWT obtained from the upload-token endpoint:\nGET /accounts/{account_id}/pages/projects/{project_name}/upload-token\n", "requestBody": {"required": true, "content": {"application/json": {"example": {"hashes": ["a948904f2f0f479b8f936b8a0c5d9882", "b026324c6904b2a9cb4b88d6d61c81d1"]}, "schema": {"$ref": "#/components/schemas/pages_pages_assets_check_missing_request"}}}}, "responses": {"200": {"description": "Check missing response.", "content": {"application/json": {"example": {"result": ["b026324c6904b2a9cb4b88d6d61c81d1"], "success": true}, "schema": {"$ref": "#/components/schemas/pages_pages_assets_check_missing_response"}}}}, "4XX": {"description": "Check missing failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"pages_upload_token": []}], "tags": ["Pages Assets"], "x-forge-hidden": true}
```
