---
title: Delete filters
page_id: operation-delete-zones-zone-id-filters-5ec30895
path: operations/filters
description: Deletes one or more existing filters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/filters
operation_ids:
    - filters-delete-filters
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete filters

`DELETE /zones/{zone_id}/filters`

Operation ID: `filters-delete-filters`

Deletes one or more existing filters.

## Definition

```yaml
{"operationId": "filters-delete-filters", "summary": "Delete filters", "description": "Deletes one or more existing filters.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "id", "in": "query", "required": true, "schema": {"type": "array", "items": {"$ref": "#/components/schemas/firewall_filters_components-schemas-id"}}}], "responses": {"200": {"description": "Delete filters response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-delete-response-collection"}}}}, "4XX": {"description": "Delete filters response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-delete-response-collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Filters"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "filters", "x-fern-sdk-method-name": "bulk-delete", "x-forge-hidden": true}
```
