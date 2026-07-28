---
title: Update a custom page
page_id: operation-put-zones-zone-identifier-custom-pages-identifier-b4aabf55
path: operations/custom-pages-for-a-zone
description: Updates the configuration of an existing custom page.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_identifier}/custom_pages/{identifier}
operation_ids:
    - custom-pages-for-a-zone-update-a-custom-page
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a custom page

`PUT /zones/{zone_identifier}/custom_pages/{identifier}`

Operation ID: `custom-pages-for-a-zone-update-a-custom-page`

Updates the configuration of an existing custom page.

## Definition

```yaml
{"operationId": "custom-pages-for-a-zone-update-a-custom-page", "summary": "Update a custom page", "description": "Updates the configuration of an existing custom page.", "parameters": [{"name": "identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_error_page_type"}}, {"name": "zone_identifier", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-pages_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"state": {"$ref": "#/components/schemas/custom-pages_state"}, "url": {"$ref": "#/components/schemas/custom-pages_url"}}, "required": ["url", "state"]}}}}, "responses": {"200": {"description": "Update a custom page response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-pages_custom_page_result"}}}}, "4XX": {"description": "Update a custom page response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-pages_custom_page_result"}, {"$ref": "#/components/schemas/custom-pages_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom pages for a zone"], "x-api-token-group": ["Custom Pages Write", "Zone Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "custom-pages", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
