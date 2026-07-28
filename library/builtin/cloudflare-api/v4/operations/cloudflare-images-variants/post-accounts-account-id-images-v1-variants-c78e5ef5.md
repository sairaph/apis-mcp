---
title: Create a variant
page_id: operation-post-accounts-account-id-images-v1-variants-4bbd00aa
path: operations/cloudflare-images-variants
description: Create a CF Images variant that allows you to resize images for different use cases.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/images/v1/variants
operation_ids:
    - cloudflare-images-variants-create-a-variant
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a variant

`POST /accounts/{account_id}/images/v1/variants`

Operation ID: `cloudflare-images-variants-create-a-variant`

Create a CF Images variant that allows you to resize images for different use cases.

## Definition

```yaml
{"operationId": "cloudflare-images-variants-create-a-variant", "summary": "Create a variant", "description": "Create a CF Images variant that allows you to resize images for different use cases.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_variant_definition"}}}}, "responses": {"200": {"description": "Create a variant response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_variant_simple_response"}}}}, "4XX": {"description": "Create a variant response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_variant_simple_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Variants"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.variants", "x-fern-sdk-method-name": "create", "x-forge-hidden": false}
```
