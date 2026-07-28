---
title: Create authenticated direct upload URL V2
page_id: operation-post-accounts-account-id-images-v2-direct-upload-51f6c030
path: operations/cloudflare-images
description: 'Direct uploads allow users to upload images without API keys. A common use case are web apps, client-side applications, or mobile devices where users upload content directly to Cloudflare Images. This method creates a draft record for a future image. It returns an upload URL and an image identifier. To verify if the image itself has been uploaded, send an image details request (accounts/:account_identifier/images/v1/:identifier), and check that the `draft: true` property is not present.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/images/v2/direct_upload
operation_ids:
    - cloudflare-images-create-authenticated-direct-upload-url-v-2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create authenticated direct upload URL V2

`POST /accounts/{account_id}/images/v2/direct_upload`

Operation ID: `cloudflare-images-create-authenticated-direct-upload-url-v-2`

Direct uploads allow users to upload images without API keys. A common use case are web apps, client-side applications, or mobile devices where users upload content directly to Cloudflare Images. This method creates a draft record for a future image. It returns an upload URL and an image identifier. To verify if the image itself has been uploaded, send an image details request (accounts/:account_identifier/images/v1/:identifier), and check that the `draft: true` property is not present.

## Definition

```yaml
{"operationId": "cloudflare-images-create-authenticated-direct-upload-url-v-2", "summary": "Create authenticated direct upload URL V2", "description": "Direct uploads allow users to upload images without API keys. A common use case are web apps, client-side applications, or mobile devices where users upload content directly to Cloudflare Images. This method creates a draft record for a future image. It returns an upload URL and an image identifier. To verify if the image itself has been uploaded, send an image details request (accounts/:account_identifier/images/v1/:identifier), and check that the `draft: true` property is not present.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"schema": {"$ref": "#/components/schemas/images_image_direct_upload_request_v2"}}}}, "responses": {"200": {"description": "Create authenticated direct upload URL V2 response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_direct_upload_response_v2"}}}}, "4XX": {"description": "Create authenticated direct upload URL V2 response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_direct_upload_response_v2"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images", "x-fern-sdk-method-name": "direct-upload", "x-forge-hidden": false}
```
