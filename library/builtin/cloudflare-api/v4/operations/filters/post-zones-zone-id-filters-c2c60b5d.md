---
title: Create filters
page_id: operation-post-zones-zone-id-filters-1d887758
path: operations/filters
description: Creates one or more filters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/filters
operation_ids:
    - filters-create-filters
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create filters

`POST /zones/{zone_id}/filters`

Operation ID: `filters-create-filters`

Creates one or more filters.

## Definition

```yaml
{"operationId": "filters-create-filters", "summary": "Create filters", "description": "Creates one or more filters.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"$ref": "#/components/schemas/firewall_filter"}}}}}, "responses": {"200": {"description": "Create filters response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-response-collection"}]}}}}, "4XX": {"description": "Create filters response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-response-collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Filters"], "x-api-token-group": ["Firewall Services Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "filters", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
