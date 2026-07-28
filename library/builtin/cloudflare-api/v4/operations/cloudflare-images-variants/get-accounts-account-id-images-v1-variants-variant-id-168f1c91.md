---
title: Variant details
page_id: operation-get-accounts-account-id-images-v1-variants-variant-id-221e5292
path: operations/cloudflare-images-variants
description: Fetch details for a CF Images variant.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v1/variants/{variant_id}
operation_ids:
    - cloudflare-images-variants-variant-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Variant details

`GET /accounts/{account_id}/images/v1/variants/{variant_id}`

Operation ID: `cloudflare-images-variants-variant-details`

Fetch details for a CF Images variant.

## Definition

```yaml
{"operationId": "cloudflare-images-variants-variant-details", "summary": "Variant details", "description": "Fetch details for a CF Images variant.", "parameters": [{"name": "variant_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_image_variant_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "responses": {"200": {"description": "Variant details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_variant_simple_response"}}}}, "4XX": {"description": "Variant details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_variant_simple_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Variants"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.variants", "x-fern-sdk-method-name": "get", "x-forge-hidden": false}
```
