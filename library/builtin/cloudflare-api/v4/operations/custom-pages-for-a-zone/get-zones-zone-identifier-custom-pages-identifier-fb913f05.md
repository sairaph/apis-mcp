---
title: Get a custom page
page_id: operation-get-zones-zone-identifier-custom-pages-identifier-74332c7f
path: operations/custom-pages-for-a-zone
description: Fetches the details of a custom page.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_identifier}/custom_pages/{identifier}
operation_ids:
    - custom-pages-for-a-zone-get-a-custom-page
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a custom page

`GET /zones/{zone_identifier}/custom_pages/{identifier}`

Operation ID: `custom-pages-for-a-zone-get-a-custom-page`

Fetches the details of a custom page.

## Definition

```yaml
{"operationId": "custom-pages-for-a-zone-get-a-custom-page", "summary": "Get a custom page", "description": "Fetches the details of a custom page.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_error_page_type"}}, {"name": "zone_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "responses": {"200": {"description": "Get a custom page response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_custom_page"}}}}, "4XX": {"description": "Get a custom page response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_custom_page_result"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom pages for a zone"], "x-api-token-group": ["Custom Pages Write", "Custom Pages Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-pages.zone-custom-pages", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
