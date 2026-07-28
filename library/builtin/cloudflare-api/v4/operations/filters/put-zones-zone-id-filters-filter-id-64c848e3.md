---
title: Update a filter
page_id: operation-put-zones-zone-id-filters-filter-id-ced91a50
path: operations/filters
description: Updates an existing filter.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/filters/{filter_id}
operation_ids:
    - filters-update-a-filter
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a filter

`PUT /zones/{zone_id}/filters/{filter_id}`

Operation ID: `filters-update-a-filter`

Updates an existing filter.

## Definition

```yaml
{"operationId": "filters-update-a-filter", "summary": "Update a filter", "description": "Updates an existing filter.", "parameters": [{"name": "filter_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_filters_components-schemas-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter"}}}}, "responses": {"200": {"description": "Update a filter response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-response-single"}}}}, "4XX": {"description": "Update a filter response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-response-single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Filters"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "filters", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
