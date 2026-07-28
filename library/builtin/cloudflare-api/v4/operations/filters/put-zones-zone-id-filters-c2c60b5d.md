---
title: Update filters
page_id: operation-put-zones-zone-id-filters-b32bd02f
path: operations/filters
description: Updates one or more existing filters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/filters
operation_ids:
    - filters-update-filters
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update filters

`PUT /zones/{zone_id}/filters`

Operation ID: `filters-update-filters`

Updates one or more existing filters.

## Definition

```yaml
{"operationId": "filters-update-filters", "summary": "Update filters", "description": "Updates one or more existing filters.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/firewall_filter-rule-update-request"}}}}}, "responses": {"200": {"description": "Update filters response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-response-collection"}}}}, "4XX": {"description": "Update filters response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-response-collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Filters"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "filters", "x-fern-sdk-method-name": "bulk-update", "x-forge-hidden": true}
```
