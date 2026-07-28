---
title: List locations
page_id: operation-get-radar-entities-locations-28f99364
path: operations/radar-locations
description: Retrieves a list of locations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/entities/locations
operation_ids:
    - radar-get-entities-locations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List locations

`GET /radar/entities/locations`

Operation ID: `radar-get-entities-locations`

Retrieves a list of locations.

## Definition

```yaml
{"operationId": "radar-get-entities-locations", "summary": "List locations", "description": "Retrieves a list of locations.", "parameters": [{"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "offset", "in": "query", "description": "Skips the specified number of objects before fetching the results.", "schema": {"description": "Skips the specified number of objects before fetching the results.", "type": "integer", "minimum": 0}}, {"name": "location", "in": "query", "description": "Filters results by location. Specify a comma-separated list of alpha-2 location codes.", "schema": {"description": "Filters results by location. Specify a comma-separated list of alpha-2 location codes.", "type": "string", "example": "US,CA"}}, {"name": "region", "in": "query", "description": "Filters results by region.", "schema": {"description": "Filters results by region.", "type": "string", "example": "Middle East", "maxLength": 100}}, {"name": "subregion", "in": "query", "description": "Filters results by subregion.", "schema": {"description": "Filters results by subregion.", "type": "string", "example": "Southern Europe", "maxLength": 100}}, {"name": "continent", "in": "query", "description": "Filters results by continent code.", "schema": {"description": "Filters results by continent code.", "type": "string", "example": "EU", "enum": ["AF", "AS", "EU", "NA", "OC", "SA"]}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"locations": {"type": "array", "items": {"properties": {"alpha2": {"type": "string", "example": "AF"}, "continent": {"type": "string", "example": "AS"}, "latitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "longitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "name": {"type": "string", "example": "Afghanistan"}, "region": {"type": "string", "example": "Middle East"}, "subregion": {"type": "string", "example": "Southern Asia"}}, "required": ["name", "region", "subregion", "continent", "latitude", "longitude", "alpha2"], "type": "object"}}}, "required": ["locations"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Locations"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.entities.locations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
