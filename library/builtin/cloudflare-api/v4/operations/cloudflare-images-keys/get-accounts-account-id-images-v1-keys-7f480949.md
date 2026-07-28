---
title: List Signing Keys
page_id: operation-get-accounts-account-id-images-v1-keys-87f527a7
path: operations/cloudflare-images-keys
description: List your CF Images signing keys.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v1/keys
operation_ids:
    - cloudflare-images-keys-list-signing-keys
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Signing Keys

`GET /accounts/{account_id}/images/v1/keys`

Operation ID: `cloudflare-images-keys-list-signing-keys`

List your CF Images signing keys.

## Definition

```yaml
{"operationId": "cloudflare-images-keys-list-signing-keys", "summary": "List Signing Keys", "description": "List your CF Images signing keys.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "responses": {"200": {"description": "List Signing Keys response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_key_response_collection"}}}}, "4XX": {"description": "List Signing Keys response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_key_response_collection"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Keys"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.keys", "x-fern-sdk-method-name": "list", "x-forge-hidden": false}
```
