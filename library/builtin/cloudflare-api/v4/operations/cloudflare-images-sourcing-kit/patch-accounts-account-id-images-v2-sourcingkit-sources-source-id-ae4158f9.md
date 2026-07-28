---
title: Update a sourcing kit source
page_id: operation-patch-accounts-account-id-images-v2-sourcingkit-sources-source-id-af3bc082
path: operations/cloudflare-images-sourcing-kit
description: Update the name of an existing migration source.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/sources/{source_id}
operation_ids:
    - cloudflare-images-sourcingkit-update-source
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a sourcing kit source

`PATCH /accounts/{account_id}/images/v2/sourcingkit/sources/{source_id}`

Operation ID: `cloudflare-images-sourcingkit-update-source`

Update the name of an existing migration source.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-update-source", "summary": "Update a sourcing kit source", "description": "Update the name of an existing migration source.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "source_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_source_update_request"}}}}, "responses": {"200": {"description": "Update sourcing kit source response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_source_update_response"}}}}, "4XX": {"description": "Update sourcing kit source response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_source_update_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.sources", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
