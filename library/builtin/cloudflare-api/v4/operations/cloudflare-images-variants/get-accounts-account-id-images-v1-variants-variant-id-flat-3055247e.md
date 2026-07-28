---
title: Variant details (flat)
page_id: operation-get-accounts-account-id-images-v1-variants-variant-id-flat-b2b4d7c9
path: operations/cloudflare-images-variants
description: Fetch details for a single variant with properties at the top level of the result.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v1/variants/{variant_id}/flat
operation_ids:
    - cloudflare-images-variants-variant-details-flat
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Variant details (flat)

`GET /accounts/{account_id}/images/v1/variants/{variant_id}/flat`

Operation ID: `cloudflare-images-variants-variant-details-flat`

Fetch details for a single variant with properties at the top level of the result.

## Definition

```yaml
{"operationId": "cloudflare-images-variants-variant-details-flat", "summary": "Variant details (flat)", "description": "Fetch details for a single variant with properties at the top level of the result.", "parameters": [{"name": "variant_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_image_variant_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "responses": {"200": {"description": "Variant details flat response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_variant_flat_response"}}}}, "4XX": {"description": "Variant details flat response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_variant_flat_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Variants"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.v1.variants-flat", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
