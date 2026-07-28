---
title: Get Geolocation details
page_id: operation-get-radar-geolocations-geo-id-b28ee88f
path: operations/radar-geolocations
description: 'Retrieves the requested Geolocation information. Geolocation names can be localized by sending an `Accept-Language` HTTP header with a BCP 47 language tag (e.g., `Accept-Language: pt-PT`). The full quality-value chain is supported (e.g., `pt-PT,pt;q=0.9,en;q=0.8`).'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/geolocations/{geo_id}
operation_ids:
    - radar-get-geolocation-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Geolocation details

`GET /radar/geolocations/{geo_id}`

Operation ID: `radar-get-geolocation-details`

Retrieves the requested Geolocation information. Geolocation names can be localized by sending an `Accept-Language` HTTP header with a BCP 47 language tag (e.g., `Accept-Language: pt-PT`). The full quality-value chain is supported (e.g., `pt-PT,pt;q=0.9,en;q=0.8`).

## Definition

```yaml
{"operationId": "radar-get-geolocation-details", "summary": "Get Geolocation details", "description": "Retrieves the requested Geolocation information. Geolocation names can be localized by sending an `Accept-Language` HTTP header with a BCP 47 language tag (e.g., `Accept-Language: pt-PT`). The full quality-value chain is supported (e.g., `pt-PT,pt;q=0.9,en;q=0.8`).", "parameters": [{"name": "geo_id", "in": "path", "description": "Geolocation ID. Refer to [GeoNames](https://download.geonames.org/export/dump/readme.txt)", "required": true, "schema": {"description": "Geolocation ID. Refer to [GeoNames](https://download.geonames.org/export/dump/readme.txt)", "type": "string", "example": "3190509", "maxLength": 100}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"geolocation": {"type": "object", "properties": {"code": {"type": "string", "example": "PT-11"}, "geoId": {"type": "string", "example": "2267056"}, "latitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "locale": {"description": "BCP 47 locale code used for the geolocation name translation", "type": "string", "example": "pt-PT"}, "longitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "name": {"type": "string", "example": "Lisbon"}, "parent": {"type": "object", "properties": {"code": {"type": "string", "example": "PT-11"}, "geoId": {"type": "string", "example": "2267056"}, "latitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "locale": {"description": "BCP 47 locale code used for the geolocation name translation", "type": "string", "example": "pt-PT"}, "longitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "name": {"type": "string", "example": "Lisbon"}, "parent": {"type": "object", "properties": {"code": {"type": "string", "example": "PT-11"}, "geoId": {"type": "string", "example": "2267056"}, "latitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "locale": {"description": "BCP 47 locale code used for the geolocation name translation", "type": "string", "example": "pt-PT"}, "longitude": {"description": "A numeric string.", "type": "string", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$"}, "name": {"type": "string", "example": "Lisbon"}, "type": {"description": "The type of the geolocation.", "type": "string", "enum": ["CONTINENT", "COUNTRY", "ADM1"]}}, "required": ["geoId", "name", "type", "latitude", "longitude"]}, "type": {"description": "The type of the geolocation.", "type": "string", "enum": ["CONTINENT", "COUNTRY", "ADM1"]}}, "required": ["geoId", "name", "type", "latitude", "longitude", "parent"]}, "type": {"description": "The type of the geolocation.", "type": "string", "enum": ["CONTINENT", "COUNTRY", "ADM1"]}}, "required": ["geoId", "name", "type", "latitude", "longitude", "parent"]}}, "required": ["geolocation"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string", "example": "Not Found."}}, "required": ["error"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Geolocations"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.geolocations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
