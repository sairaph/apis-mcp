---
title: List custom assets
page_id: operation-get-zones-zone-identifier-custom-pages-assets-38f7efe8
path: operations/custom-assets-for-a-zone
description: Fetches all the custom assets at the zone level.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_identifier}/custom_pages/assets
operation_ids:
    - custom-assets-for-a-zone-list-custom-assets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List custom assets

`GET /zones/{zone_identifier}/custom_pages/assets`

Operation ID: `custom-assets-for-a-zone-list-custom-assets`

Fetches all the custom assets at the zone level.

## Definition

```yaml
{"operationId": "custom-assets-for-a-zone-list-custom-assets", "summary": "List custom assets", "description": "Fetches all the custom assets at the zone level.", "parameters": [{"name": "zone_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}, {"name": "page", "in": "query", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 20, "maximum": 200, "minimum": 1}}], "responses": {"200": {"description": "List custom assets response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_custom_asset_result_list"}}}}, "4XX": {"description": "List custom assets response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_custom_asset_result_list"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom assets for a zone"], "x-api-token-group": ["Custom Errors Write", "Custom Errors Read"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-pages.assets.list", "x-fern-sdk-method-name": "v2", "x-forge-hidden": true}
```
