---
title: List variants
page_id: operation-get-accounts-account-id-images-v1-variants-9a0e21dd
path: operations/cloudflare-images-variants
description: List existing CF Images variants.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v1/variants
operation_ids:
    - cloudflare-images-variants-list-variants
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List variants

`GET /accounts/{account_id}/images/v1/variants`

Operation ID: `cloudflare-images-variants-list-variants`

List existing CF Images variants.

## Definition

```yaml
{"operationId": "cloudflare-images-variants-list-variants", "summary": "List variants", "description": "List existing CF Images variants.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "responses": {"200": {"description": "List variants response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_variant_list_response"}}}}, "4XX": {"description": "List variants response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_variant_list_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Variants"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.variants", "x-fern-sdk-method-name": "list", "x-forge-hidden": false}
```
