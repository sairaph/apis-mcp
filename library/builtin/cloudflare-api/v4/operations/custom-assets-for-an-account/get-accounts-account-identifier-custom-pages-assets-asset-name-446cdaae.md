---
title: Get a custom asset
page_id: operation-get-accounts-account-identifier-custom-pages-assets-asset-name-b37cd16c
path: operations/custom-assets-for-an-account
description: Fetches the details of a custom asset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_identifier}/custom_pages/assets/{asset_name}
operation_ids:
    - custom-assets-for-an-account-get-a-custom-asset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a custom asset

`GET /accounts/{account_identifier}/custom_pages/assets/{asset_name}`

Operation ID: `custom-assets-for-an-account-get-a-custom-asset`

Fetches the details of a custom asset.

## Definition

```yaml
{"operationId": "custom-assets-for-an-account-get-a-custom-asset", "summary": "Get a custom asset", "description": "Fetches the details of a custom asset.", "parameters": [{"name": "asset_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_asset_name"}}, {"name": "account_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "responses": {"200": {"description": "Get a custom asset response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_custom_asset_result"}}}}, "4XX": {"description": "Get a custom asset response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_custom_asset_result"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom assets for an account"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-pages.assets.get.custom", "x-fern-sdk-method-name": "asset", "x-forge-hidden": true}
```
