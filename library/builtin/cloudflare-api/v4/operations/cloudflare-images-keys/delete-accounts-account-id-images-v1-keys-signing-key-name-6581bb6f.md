---
title: Delete Signing Key
page_id: operation-delete-accounts-account-id-images-v1-keys-signing-key-name-c0f8529c
path: operations/cloudflare-images-keys
description: |-
    Delete a CF Images signing key with specified name. Returns all keys available.
    When the last key is removed, a new default signing key will be generated.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/images/v1/keys/{signing_key_name}
operation_ids:
    - cloudflare-images-keys-delete-signing-key
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Signing Key

`DELETE /accounts/{account_id}/images/v1/keys/{signing_key_name}`

Operation ID: `cloudflare-images-keys-delete-signing-key`

Delete a CF Images signing key with specified name. Returns all keys available.
When the last key is removed, a new default signing key will be generated.

## Definition

```yaml
{"operationId": "cloudflare-images-keys-delete-signing-key", "summary": "Delete Signing Key", "description": "Delete a CF Images signing key with specified name. Returns all keys available.\nWhen the last key is removed, a new default signing key will be generated.\n", "parameters": [{"name": "signing_key_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_signing_key_identifier"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}], "responses": {"200": {"description": "Delete Signing Key response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_image_key_response_collection"}}}}, "4XX": {"description": "Delete Signing Key response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_image_key_response_collection"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Keys"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.keys", "x-fern-sdk-method-name": "delete", "x-forge-hidden": false, "x-forge-require-confirmation": "This operation will delete this signing key. This will invalidate any URLs that are signed with this key."}
```
