---
title: Get location details
page_id: operation-get-radar-entities-locations-location-404ea2ad
path: operations/radar-locations
description: Retrieves the requested location information. (A confidence level below `5` indicates a low level of confidence in the traffic data - normally this happens because Cloudflare has a small amount of traffic from/to this location).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/entities/locations/{location}
operation_ids:
    - radar-get-entities-location-by-alpha2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get location details

`GET /radar/entities/locations/{location}`

Operation ID: `radar-get-entities-location-by-alpha2`

Retrieves the requested location information. (A confidence level below `5` indicates a low level of confidence in the traffic data - normally this happens because Cloudflare has a small amount of traffic from/to this location).

## Definition

```yaml
{"operationId": "radar-get-entities-location-by-alpha2", "summary": "Get location details", "description": "Retrieves the requested location information. (A confidence level below `5` indicates a low level of confidence in the traffic data - normally this happens because Cloudflare has a small amount of traffic from/to this location).", "parameters": [{"name": "location", "in": "path", "description": "Location alpha-2 code.", "required": true, "schema": {"description": "Location alpha-2 code.", "type": "string", "example": "US", "maxLength": 2, "minLength": 2}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"location": {"type": "object", "properties": {"alpha2": {"type": "string", "example": "AF"}, "confidenceLevel": {"type": "integer", "example": 5}, "continent": {"type": "string", "example": "AS"}, "latitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "longitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "name": {"type": "string", "example": "Afghanistan"}, "region": {"type": "string", "example": "Middle East"}, "subregion": {"type": "string", "example": "Southern Asia"}}, "required": ["name", "region", "subregion", "continent", "latitude", "longitude", "alpha2", "confidenceLevel"]}}, "required": ["location"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string", "example": "Not Found."}}, "required": ["error"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Locations"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.entities.locations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
