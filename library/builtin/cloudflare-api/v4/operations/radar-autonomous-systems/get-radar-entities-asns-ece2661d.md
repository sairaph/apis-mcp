---
title: List autonomous systems
page_id: operation-get-radar-entities-asns-2015c2ec
path: operations/radar-autonomous-systems
description: Retrieves a list of autonomous systems.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/entities/asns
operation_ids:
    - radar-get-entities-asn-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List autonomous systems

`GET /radar/entities/asns`

Operation ID: `radar-get-entities-asn-list`

Retrieves a list of autonomous systems.

## Definition

```yaml
{"operationId": "radar-get-entities-asn-list", "summary": "List autonomous systems", "description": "Retrieves a list of autonomous systems.", "parameters": [{"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "offset", "in": "query", "description": "Skips the specified number of objects before fetching the results.", "schema": {"description": "Skips the specified number of objects before fetching the results.", "type": "integer", "minimum": 0}}, {"name": "asn", "in": "query", "description": "Filters results by Autonomous System. Specify one or more Autonomous System Numbers (ASNs) as a comma-separated list.", "schema": {"description": "Filters results by Autonomous System. Specify one or more Autonomous System Numbers (ASNs) as a comma-separated list.", "type": "string", "example": "174,7922"}}, {"name": "location", "in": "query", "description": "Filters results by location. Specify an alpha-2 location code.", "schema": {"description": "Filters results by location. Specify an alpha-2 location code.", "type": "string", "example": "US", "maxLength": 2, "minLength": 2}}, {"name": "orderBy", "in": "query", "description": "Specifies the metric to order the ASNs by.", "schema": {"description": "Specifies the metric to order the ASNs by.", "type": "string", "default": "ASN", "enum": ["ASN", "POPULATION"]}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"asns": {"type": "array", "items": {"properties": {"aka": {"type": "string"}, "asn": {"type": "integer", "example": 714}, "country": {"type": "string", "example": "GB"}, "countryName": {"type": "string", "example": "United Kingdom"}, "estimatedUsers": {"type": "object", "properties": {"estimatedUsers": {"description": "Total estimated users.", "type": "integer", "example": 86099}}}, "name": {"type": "string", "example": "Apple Inc."}, "orgName": {"type": "string"}, "website": {"type": "string", "example": "https://www.apple.com/support/systemstatus/"}}, "required": ["name", "asn", "country", "countryName", "estimatedUsers"], "type": "object"}}}, "required": ["asns"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Autonomous Systems"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.entities.asns", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
