---
title: Get source connectivity status
page_id: operation-get-accounts-account-id-images-v2-sourcingkit-sources-source-id-connecti-e358b54d
path: operations/cloudflare-images-sourcing-kit
description: Check the current connectivity status of an existing migration source.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/images/v2/sourcingkit/sources/{source_id}/connectivity
operation_ids:
    - cloudflare-images-sourcingkit-get-source-connectivity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get source connectivity status

`GET /accounts/{account_id}/images/v2/sourcingkit/sources/{source_id}/connectivity`

Operation ID: `cloudflare-images-sourcingkit-get-source-connectivity`

Check the current connectivity status of an existing migration source.

## Definition

```yaml
{"operationId": "cloudflare-images-sourcingkit-get-source-connectivity", "summary": "Get source connectivity status", "description": "Check the current connectivity status of an existing migration source.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_account_identifier"}}, {"name": "source_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/images_sourcingkit_identifier"}}], "responses": {"200": {"description": "Source connectivity status response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/images_sourcingkit_connectivity_check_response"}}}}, "4XX": {"description": "Source connectivity status response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/images_sourcingkit_connectivity_check_response"}, {"$ref": "#/components/schemas/images_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Cloudflare Images Sourcing Kit"], "x-api-token-group": ["Images Read", "Images Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "images.sourcing-kit.sources", "x-fern-sdk-method-name": "connectivity", "x-forge-hidden": true}
```
