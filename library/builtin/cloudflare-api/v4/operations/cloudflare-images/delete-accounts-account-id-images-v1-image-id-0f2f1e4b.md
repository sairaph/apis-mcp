---
title: Delete image
page_id: operation-delete-accounts-account-id-images-v1-image-id-0c325eaa
path: operations/cloudflare-images
description: Delete an image on Cloudflare Images. On success, all copies of the image are deleted and purged from cache.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/images/v1/{image_id}
operation_ids:
    - cloudflare-images-delete-image
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete image

`DELETE /accounts/{account_id}/images/v1/{image_id}`

Operation ID: `cloudflare-images-delete-image`

Delete an image on Cloudflare Images. On success, all copies of the image are deleted and purged from cache.

## Definition

```yaml
{"operationId": "cloudflare-images-delete-image", "summary": "Delete image", "description": "Delete an image on Cloudflare Images. On success, all copies of the image are deleted and purged from cache.", "parameters": [{"name": "image_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_image_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete image response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_deleted_response"}}}}, "4XX": {"description": "Delete image response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_deleted_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images", "x-fern-sdk-method-name": "delete", "x-forge-hidden": false, "x-forge-require-confirmation": "This operation will delete this image."}
```
