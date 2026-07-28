---
title: Download image
page_id: operation-get-accounts-account-id-images-v1-image-id-blob-b9c1598f
path: operations/cloudflare-images
description: Download an image from CF Images. For most images this will be the originally uploaded file. For larger images it can be a near-lossless version of the original.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v1/{image_id}/blob
operation_ids:
    - cloudflare-images-base-image
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download image

`GET /accounts/{account_id}/images/v1/{image_id}/blob`

Operation ID: `cloudflare-images-base-image`

Download an image from CF Images. For most images this will be the originally uploaded file. For larger images it can be a near-lossless version of the original.

## Definition

```yaml
{"operationId": "cloudflare-images-base-image", "summary": "Download image", "description": "Download an image from CF Images. For most images this will be the originally uploaded file. For larger images it can be a near-lossless version of the original.", "parameters": [{"name": "image_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_image_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "responses": {"200": {"description": "Base image response. Returns uploaded image data.", "content": {"image/*": {"schema": {"type": "string", "format": "binary"}}}}, "4XX": {"description": "Base image response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_response_blob"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images", "x-fern-sdk-method-name": "get-blob", "x-forge-hidden": false}
```
