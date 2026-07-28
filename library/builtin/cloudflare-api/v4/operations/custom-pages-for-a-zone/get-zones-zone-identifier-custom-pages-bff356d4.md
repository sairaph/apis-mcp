---
title: List custom pages
page_id: operation-get-zones-zone-identifier-custom-pages-033bef01
path: operations/custom-pages-for-a-zone
description: Fetches all the custom pages at the zone level.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_identifier}/custom_pages
operation_ids:
    - custom-pages-for-a-zone-list-custom-pages
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List custom pages

`GET /zones/{zone_identifier}/custom_pages`

Operation ID: `custom-pages-for-a-zone-list-custom-pages`

Fetches all the custom pages at the zone level.

## Definition

```yaml
{"operationId": "custom-pages-for-a-zone-list-custom-pages", "summary": "List custom pages", "description": "Fetches all the custom pages at the zone level.", "parameters": [{"name": "zone_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "responses": {"200": {"description": "List custom pages response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_custom_page_result_list"}}}}, "4XX": {"description": "List custom pages response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_custom_page_result_list"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom pages for a zone"], "x-api-token-group": ["Custom Pages Write", "Custom Pages Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-pages.zone-custom-pages", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
