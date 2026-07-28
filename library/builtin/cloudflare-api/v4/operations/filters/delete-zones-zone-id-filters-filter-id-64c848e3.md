---
title: Delete a filter
page_id: operation-delete-zones-zone-id-filters-filter-id-c27aa9de
path: operations/filters
description: Deletes an existing filter.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/filters/{filter_id}
operation_ids:
    - filters-delete-a-filter
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a filter

`DELETE /zones/{zone_id}/filters/{filter_id}`

Operation ID: `filters-delete-a-filter`

Deletes an existing filter.

## Definition

```yaml
{"operationId": "filters-delete-a-filter", "summary": "Delete a filter", "description": "Deletes an existing filter.", "parameters": [{"name": "filter_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_filters_components-schemas-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "responses": {"200": {"description": "Delete a filter response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-delete-response-single"}}}}, "4XX": {"description": "Delete a filter response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-delete-response-single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Filters"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "filters", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
