---
title: Image details
page_id: operation-get-accounts-account-id-images-v1-image-id-e643d1f1
path: operations/cloudflare-images
description: Fetch details for a CF Images image.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v1/{image_id}
operation_ids:
    - cloudflare-images-image-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Image details

`GET /accounts/{account_id}/images/v1/{image_id}`

Operation ID: `cloudflare-images-image-details`

Fetch details for a CF Images image.

## Definition

```yaml
{"operationId": "cloudflare-images-image-details", "summary": "Image details", "description": "Fetch details for a CF Images image.", "parameters": [{"name": "image_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_image_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "responses": {"200": {"description": "Image details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_response_single"}}}}, "4XX": {"description": "Image details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_response_single"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images", "x-fern-sdk-method-name": "get", "x-forge-hidden": false}
```
