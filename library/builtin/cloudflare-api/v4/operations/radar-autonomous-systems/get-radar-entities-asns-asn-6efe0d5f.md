---
title: Get AS details by ASN
page_id: operation-get-radar-entities-asns-asn-9d46a042
path: operations/radar-autonomous-systems
description: Retrieves the requested autonomous system information. (A confidence level below `5` indicates a low level of confidence in the traffic data - normally this happens because Cloudflare has a small amount of traffic from/to this AS). Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/entities/asns/{asn}
operation_ids:
    - radar-get-entities-asn-by-id
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get AS details by ASN

`GET /radar/entities/asns/{asn}`

Operation ID: `radar-get-entities-asn-by-id`

Retrieves the requested autonomous system information. (A confidence level below `5` indicates a low level of confidence in the traffic data - normally this happens because Cloudflare has a small amount of traffic from/to this AS). Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).

## Definition

```yaml
{"operationId": "radar-get-entities-asn-by-id", "summary": "Get AS details by ASN", "description": "Retrieves the requested autonomous system information. (A confidence level below `5` indicates a low level of confidence in the traffic data - normally this happens because Cloudflare has a small amount of traffic from/to this AS). Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).", "parameters": [{"name": "asn", "in": "path", "description": "Single Autonomous System Number (ASN) as integer.", "required": true, "schema": {"description": "Single Autonomous System Number (ASN) as integer.", "type": "integer", "example": 174}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"asn": {"type": "object", "properties": {"aka": {"type": "string"}, "asn": {"type": "integer", "example": 714}, "confidenceLevel": {"type": "integer", "example": 5}, "country": {"type": "string", "example": "GB"}, "countryName": {"type": "string", "example": "United Kingdom"}, "estimatedUsers": {"type": "object", "properties": {"estimatedUsers": {"description": "Total estimated users.", "type": "integer", "example": 86099}, "locations": {"type": "array", "items": {"properties": {"estimatedUsers": {"description": "Estimated users per location.", "type": "integer", "example": 16710}, "locationAlpha2": {"type": "string", "example": "US"}, "locationName": {"type": "string", "example": "United States"}}, "required": ["locationName", "locationAlpha2"], "type": "object"}}}, "required": ["locations"]}, "name": {"type": "string", "example": "Apple Inc."}, "orgName": {"type": "string"}, "related": {"type": "array", "items": {"properties": {"aka": {"type": "string"}, "asn": {"type": "integer", "example": 174}, "estimatedUsers": {"description": "Total estimated users.", "type": "integer", "example": 65345}, "name": {"type": "string", "example": "Cogent-174"}}, "required": ["name", "asn"], "type": "object"}}, "source": {"description": "Regional Internet Registry.", "type": "string", "example": "RIPE"}, "website": {"type": "string", "example": "https://www.apple.com/support/systemstatus/"}}, "required": ["name", "country", "countryName", "confidenceLevel", "related", "source", "asn", "website", "orgName", "estimatedUsers"]}}, "required": ["asn"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string", "example": "Not Found."}}, "required": ["error"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Autonomous Systems"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.entities.asns", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
