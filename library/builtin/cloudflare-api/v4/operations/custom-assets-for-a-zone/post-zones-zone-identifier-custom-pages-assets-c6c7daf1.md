---
title: Create a custom asset
page_id: operation-post-zones-zone-identifier-custom-pages-assets-1780f36a
path: operations/custom-assets-for-a-zone
description: Creates a new custom asset at the zone level.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_identifier}/custom_pages/assets
operation_ids:
    - custom-assets-for-a-zone-create-a-custom-asset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a custom asset

`POST /zones/{zone_identifier}/custom_pages/assets`

Operation ID: `custom-assets-for-a-zone-create-a-custom-asset`

Creates a new custom asset at the zone level.

## Definition

```yaml
{"operationId": "custom-assets-for-a-zone-create-a-custom-asset", "summary": "Create a custom asset", "description": "Creates a new custom asset at the zone level.", "parameters": [{"name": "zone_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"description": {"$ref": "#/components/schemas/custom-pages_asset_description"}, "name": {"$ref": "#/components/schemas/custom-pages_asset_name"}, "url": {"$ref": "#/components/schemas/custom-pages_asset_url"}}, "required": ["name", "description", "url"]}}}}, "responses": {"200": {"description": "Create custom asset response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_custom_asset_result"}}}}, "4XX": {"description": "Create custom asset response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_custom_asset_result"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom assets for a zone"], "x-api-token-group": ["Custom Errors Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-pages.assets.create", "x-fern-sdk-method-name": "v2", "x-forge-hidden": true}
```
