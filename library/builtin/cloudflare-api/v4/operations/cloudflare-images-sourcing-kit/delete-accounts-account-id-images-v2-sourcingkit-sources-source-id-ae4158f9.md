---
title: Delete a sourcing kit source
page_id: operation-delete-accounts-account-id-images-v2-sourcingkit-sources-source-id-38113b93
path: operations/cloudflare-images-sourcing-kit
description: Delete an existing migration source. Sources with active migrations cannot be deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/sources/{source_id}
operation_ids:
    - cloudflare-images-sourcingkit-delete-source
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a sourcing kit source

`DELETE /accounts/{account_id}/images/v2/sourcingkit/sources/{source_id}`

Operation ID: `cloudflare-images-sourcingkit-delete-source`

Delete an existing migration source. Sources with active migrations cannot be deleted.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-delete-source", "summary": "Delete a sourcing kit source", "description": "Delete an existing migration source. Sources with active migrations cannot be deleted.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "source_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}], "responses": {"200": {"description": "Delete sourcing kit source response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_deleted_response"}}}}, "4XX": {"description": "Delete sourcing kit source response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_deleted_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.sources", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
