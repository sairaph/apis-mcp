---
title: Update a variant
page_id: operation-patch-accounts-account-id-images-v1-variants-variant-id-9ec7ca58
path: operations/cloudflare-images-variants
description: Update a CF Images variant. This will purge the cache for all images associated with the variant.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/images/v1/variants/{variant_id}
operation_ids:
    - cloudflare-images-variants-update-a-variant
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a variant

`PATCH /accounts/{account_id}/images/v1/variants/{variant_id}`

Operation ID: `cloudflare-images-variants-update-a-variant`

Update a CF Images variant. This will purge the cache for all images associated with the variant.

## Definition

```yaml
{"operationId": "cloudflare-images-variants-update-a-variant", "summary": "Update a variant", "description": "Update a CF Images variant. This will purge the cache for all images associated with the variant.", "parameters": [{"name": "variant_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_image_variant_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_variant_patch_request"}}}}, "responses": {"200": {"description": "Update a variant response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_variant_simple_response"}}}}, "4XX": {"description": "Update a variant response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_variant_simple_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Variants"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.variants", "x-fern-sdk-method-name": "edit", "x-forge-hidden": false}
```
