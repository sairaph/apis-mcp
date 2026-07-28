---
title: Get sourcing kit source
page_id: operation-get-accounts-account-id-images-v2-sourcingkit-sources-source-id-2b3cd1c1
path: operations/cloudflare-images-sourcing-kit
description: Fetch details for a single migration source.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/sources/{source_id}
operation_ids:
    - cloudflare-images-sourcingkit-get-source
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get sourcing kit source

`GET /accounts/{account_id}/images/v2/sourcingkit/sources/{source_id}`

Operation ID: `cloudflare-images-sourcingkit-get-source`

Fetch details for a single migration source.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-get-source", "summary": "Get sourcing kit source", "description": "Fetch details for a single migration source.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "source_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}], "responses": {"200": {"description": "Get sourcing kit source response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_source_single_response"}}}}, "4XX": {"description": "Get sourcing kit source response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_source_single_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.sources", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
