---
title: Create a new Signing Key
page_id: operation-put-accounts-account-id-images-v1-keys-signing-key-name-87ae2c80
path: operations/cloudflare-images-keys
description: Create a new CF Images signing key with specified name. Returns all keys available.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/images/v1/keys/{signing_key_name}
operation_ids:
    - cloudflare-images-keys-add-signing-key
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Signing Key

`PUT /accounts/{account_id}/images/v1/keys/{signing_key_name}`

Operation ID: `cloudflare-images-keys-add-signing-key`

Create a new CF Images signing key with specified name. Returns all keys available.

## Definition

```yaml
{"operationId": "cloudflare-images-keys-add-signing-key", "summary": "Create a new Signing Key", "description": "Create a new CF Images signing key with specified name. Returns all keys available.", "parameters": [{"name": "signing_key_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_signing_key_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "responses": {"200": {"description": "Add Signing Key response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_key_response_collection"}}}}, "4XX": {"description": "Add Signing Key response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_key_response_collection"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Keys"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.keys", "x-fern-sdk-method-name": "create", "x-forge-hidden": false}
```
