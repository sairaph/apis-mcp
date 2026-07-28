---
title: List filters
page_id: operation-get-zones-zone-id-filters-4c802421
path: operations/filters
description: Fetches filters in a zone. You can filter the results using several optional parameters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/filters
operation_ids:
    - filters-list-filters
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List filters

`GET /zones/{zone_id}/filters`

Operation ID: `filters-list-filters`

Fetches filters in a zone. You can filter the results using several optional parameters.

## Definition

```yaml
{"operationId": "filters-list-filters", "summary": "List filters", "description": "Fetches filters in a zone. You can filter the results using several optional parameters.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/firewall_identifier"}}, {"name": "paused", "in": "query", "schema": {"oneOf": [{"$ref": "#/components/schemas/firewall_filters_components-schemas-paused"}]}}, {"name": "expression", "in": "query", "schema": {"description": "A case-insensitive string to find in the expression.", "type": "string", "example": "php"}}, {"name": "description", "in": "query", "schema": {"description": "A case-insensitive string to find in the description.", "type": "string", "example": "browsers"}}, {"name": "ref", "in": "query", "schema": {"description": "The filter ref (a short reference tag) to search for. Must be an exact match.", "type": "string", "example": "FIL-100"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Number of filters per page.", "type": "number", "default": 25, "maximum": 100, "minimum": 5}}, {"name": "id", "in": "query", "schema": {"description": "The unique identifier of the filter.", "type": "string", "example": "372e67954025e0ba6aaa6d586b9e0b61", "maxLength": 32, "minLength": 32, "readOnly": true}}], "responses": {"200": {"description": "List filters response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/firewall_filter-response-collection"}}}}, "4XX": {"description": "List filters response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/firewall_filter-response-collection"}, {"$ref": "#/components/schemas/firewall_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Filters"], "x-api-token-group": ["Firewall Services Write", "Firewall Services Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "filters", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
