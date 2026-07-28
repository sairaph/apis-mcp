---
title: Upload an image
page_id: operation-post-accounts-account-id-images-v1-2a5f7142
path: operations/cloudflare-images
description: |-
    Upload an image to CF Images. Images up to 10 Megabytes can be uploaded using a
    single HTTP POST (multipart/form-data) request by sending an image file or
    passing a URL accessible to the API.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/images/v1
operation_ids:
    - cloudflare-images-upload-an-image-via-url
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Upload an image

`POST /accounts/{account_id}/images/v1`

Operation ID: `cloudflare-images-upload-an-image-via-url`

Upload an image to CF Images. Images up to 10 Megabytes can be uploaded using a
single HTTP POST (multipart/form-data) request by sending an image file or
passing a URL accessible to the API.

## Definition

```yaml
{"operationId": "cloudflare-images-upload-an-image-via-url", "summary": "Upload an image", "description": "Upload an image to CF Images. Images up to 10 Megabytes can be uploaded using a\nsingle HTTP POST (multipart/form-data) request by sending an image file or\npassing a URL accessible to the API.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"schema": {"$ref": "#/components/schemas/images_image_basic_upload"}}}}, "responses": {"200": {"description": "Upload an image response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_response_single"}}}}, "4XX": {"description": "Upload an image response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_response_single"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images", "x-fern-sdk-method-name": "create", "x-forge-hidden": false}
```
