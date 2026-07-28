---
title: Create authenticated direct upload URL V1
page_id: operation-post-accounts-account-id-images-v1-direct-upload-103e8b88
path: operations/cloudflare-images
description: |-
    Direct uploads allow users to upload images without API keys. A common use
    case are web apps, client-side applications, or mobile devices where users
    upload content directly to Cloudflare Images. This method creates a one-time
    upload URL. Use the V2 endpoint for additional features such as custom IDs and
    metadata.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/images/v1/direct_upload
operation_ids:
    - cloudflare-images-create-authenticated-direct-upload-url-v-1
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create authenticated direct upload URL V1

`POST /accounts/{account_id}/images/v1/direct_upload`

Operation ID: `cloudflare-images-create-authenticated-direct-upload-url-v-1`

Direct uploads allow users to upload images without API keys. A common use
case are web apps, client-side applications, or mobile devices where users
upload content directly to Cloudflare Images. This method creates a one-time
upload URL. Use the V2 endpoint for additional features such as custom IDs and
metadata.

## Definition

```yaml
{"operationId": "cloudflare-images-create-authenticated-direct-upload-url-v-1", "summary": "Create authenticated direct upload URL V1", "description": "Direct uploads allow users to upload images without API keys. A common use\ncase are web apps, client-side applications, or mobile devices where users\nupload content directly to Cloudflare Images. This method creates a one-time\nupload URL. Use the V2 endpoint for additional features such as custom IDs and\nmetadata.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_direct_upload_request_v1"}}}}, "responses": {"200": {"description": "Create authenticated direct upload URL V1 response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_direct_upload_response_v1"}}}}, "4XX": {"description": "Create authenticated direct upload URL V1 response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_direct_upload_response_v1"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images"], "x-api-token-group": ["Images Write"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "images.v1.direct-uploads", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
