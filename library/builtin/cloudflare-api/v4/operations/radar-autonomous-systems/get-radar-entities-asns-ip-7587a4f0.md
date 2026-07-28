---
title: Get AS details by IP address
page_id: operation-get-radar-entities-asns-ip-06a3417c
path: operations/radar-autonomous-systems
description: Retrieves the requested autonomous system information based on IP address. Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/entities/asns/ip
operation_ids:
    - radar-get-entities-asn-by-ip
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get AS details by IP address

`GET /radar/entities/asns/ip`

Operation ID: `radar-get-entities-asn-by-ip`

Retrieves the requested autonomous system information based on IP address. Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).

## Definition

```yaml
{"operationId": "radar-get-entities-asn-by-ip", "summary": "Get AS details by IP address", "description": "Retrieves the requested autonomous system information based on IP address. Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).", "parameters": [{"name": "ip", "in": "query", "description": "IP address.", "required": true, "schema": {"description": "IP address.", "type": "string", "format": "ip", "example": "8.8.8.8"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"asn": {"type": "object", "properties": {"aka": {"type": "string"}, "asn": {"type": "integer", "example": 714}, "country": {"type": "string", "example": "GB"}, "countryName": {"type": "string", "example": "United Kingdom"}, "estimatedUsers": {"type": "object", "properties": {"estimatedUsers": {"description": "Total estimated users.", "type": "integer", "example": 86099}, "locations": {"type": "array", "items": {"properties": {"estimatedUsers": {"description": "Estimated users per location.", "type": "integer", "example": 16710}, "locationAlpha2": {"type": "string", "example": "US"}, "locationName": {"type": "string", "example": "United States"}}, "required": ["locationName", "locationAlpha2"], "type": "object"}}}, "required": ["locations"]}, "name": {"type": "string", "example": "Apple Inc."}, "orgName": {"type": "string"}, "related": {"type": "array", "items": {"properties": {"aka": {"type": "string"}, "asn": {"type": "integer"}, "estimatedUsers": {"description": "Total estimated users.", "type": "integer", "example": 65345}, "name": {"type": "string"}}, "required": ["name", "asn"], "type": "object"}}, "source": {"description": "Regional Internet Registry.", "type": "string", "example": "RIPE"}, "website": {"type": "string", "example": "https://www.apple.com/support/systemstatus/"}}, "required": ["name", "country", "countryName", "related", "asn", "website", "orgName", "source", "estimatedUsers"]}}, "required": ["asn"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string", "example": "Not Found."}}, "required": ["error"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Autonomous Systems"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.entities.asns", "x-fern-sdk-method-name": "ip", "x-forge-hidden": true}
```
