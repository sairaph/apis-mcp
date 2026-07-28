---
title: Get a filter
page_id: operation-get-zones-zone-id-filters-filter-id-d651cb7c
path: operations/filters
description: Fetches the details of a filter.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/filters/{filter_id}
operation_ids:
    - filters-get-a-filter
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a filter

`GET /zones/{zone_id}/filters/{filter_id}`

Operation ID: `filters-get-a-filter`

Fetches the details of a filter.

## Definition

```yaml
{"operationId": "filters-get-a-filter", "summary": "Get a filter", "description": "Fetches the details of a filter.", "parameters": [{"name": "filter_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_filters_components-schemas-id"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "responses": {"200": {"description": "Get a filter response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-response-single"}}}}, "4XX": {"description": "Get a filter response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-response-single"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Filters"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "filters", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
