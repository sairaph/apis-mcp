---
title: Update image
page_id: operation-patch-accounts-account-id-images-v1-image-id-fa034132
path: operations/cloudflare-images
description: Update a CF Images image's metadata, creator, or access control. On access control change, all copies of the image are purged from cache.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/images/v1/{image_id}
operation_ids:
    - cloudflare-images-update-image
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update image

`PATCH /accounts/{account_id}/images/v1/{image_id}`

Operation ID: `cloudflare-images-update-image`

Update a CF Images image's metadata, creator, or access control. On access control change, all copies of the image are purged from cache.

## Definition

```yaml
{"operationId": "cloudflare-images-update-image", "summary": "Update image", "description": "Update a CF Images image's metadata, creator, or access control. On access control change, all copies of the image are purged from cache.", "parameters": [{"name": "image_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_image_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_patch_request"}}}}, "responses": {"200": {"description": "Update image response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_response_single"}}}}, "4XX": {"description": "Update image response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_response_single"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images", "x-fern-sdk-method-name": "edit", "x-forge-hidden": false}
```
