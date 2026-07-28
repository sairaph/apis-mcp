---
title: Update a custom asset
page_id: operation-put-zones-zone-identifier-custom-pages-assets-asset-name-9363bdbd
path: operations/custom-assets-for-a-zone
description: Updates the configuration of an existing custom asset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_identifier}/custom_pages/assets/{asset_name}
operation_ids:
    - custom-assets-for-a-zone-update-a-custom-asset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a custom asset

`PUT /zones/{zone_identifier}/custom_pages/assets/{asset_name}`

Operation ID: `custom-assets-for-a-zone-update-a-custom-asset`

Updates the configuration of an existing custom asset.

## Definition

```yaml
{"operationId": "custom-assets-for-a-zone-update-a-custom-asset", "summary": "Update a custom asset", "description": "Updates the configuration of an existing custom asset.", "parameters": [{"name": "asset_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_asset_name"}}, {"name": "zone_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"description": {"$ref": "#/components/schemas/custom-pages_asset_description"}, "url": {"$ref": "#/components/schemas/custom-pages_asset_url"}}, "required": ["description", "url"]}}}}, "responses": {"200": {"description": "Update a custom asset response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_custom_asset_result"}}}}, "4XX": {"description": "Update a custom asset response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_custom_asset_result"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom assets for a zone"], "x-api-token-group": ["Custom Errors Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-pages.assets.update", "x-fern-sdk-method-name": "v2", "x-forge-hidden": true}
```
