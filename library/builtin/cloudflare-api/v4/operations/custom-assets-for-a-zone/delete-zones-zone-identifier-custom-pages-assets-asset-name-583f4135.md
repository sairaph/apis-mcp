---
title: Delete a custom asset
page_id: operation-delete-zones-zone-identifier-custom-pages-assets-asset-name-3880a9e8
path: operations/custom-assets-for-a-zone
description: Deletes an existing custom asset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_identifier}/custom_pages/assets/{asset_name}
operation_ids:
    - custom-assets-for-a-zone-delete-a-custom-asset
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a custom asset

`DELETE /zones/{zone_identifier}/custom_pages/assets/{asset_name}`

Operation ID: `custom-assets-for-a-zone-delete-a-custom-asset`

Deletes an existing custom asset.

## Definition

```yaml
{"operationId": "custom-assets-for-a-zone-delete-a-custom-asset", "summary": "Delete a custom asset", "description": "Deletes an existing custom asset.", "parameters": [{"name": "asset_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_asset_name"}}, {"name": "zone_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "responses": {"204": {"description": "Delete a custom asset response"}, "4XX": {"description": "Delete a custom asset response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom assets for a zone"], "x-api-token-group": ["Custom Errors Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-pages.assets.delete", "x-fern-sdk-method-name": "v2", "x-forge-hidden": true}
```
