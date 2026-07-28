---
title: List ASes from global routing tables
page_id: operation-get-radar-bgp-routes-ases-fb895179
path: operations/radar-bgp
description: Retrieves all ASes in the current global routing tables with routing statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/routes/ases
operation_ids:
    - radar-get-bgp-routes-asns
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List ASes from global routing tables

`GET /radar/bgp/routes/ases`

Operation ID: `radar-get-bgp-routes-asns`

Retrieves all ASes in the current global routing tables with routing statistics.

## Definition

```yaml
{"operationId": "radar-get-bgp-routes-asns", "summary": "List ASes from global routing tables", "description": "Retrieves all ASes in the current global routing tables with routing statistics.", "parameters": [{"name": "location", "in": "query", "description": "Filters results by location. Specify an alpha-2 location code.", "schema": {"description": "Filters results by location. Specify an alpha-2 location code.", "type": "string", "example": "US", "maxLength": 2, "minLength": 2}}, {"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "sortBy", "in": "query", "description": "Sorts results by the specified field.", "schema": {"description": "Sorts results by the specified field.", "type": "string", "example": "ipv4", "enum": ["cone", "pfxs", "ipv4", "ipv6", "rpki_valid", "rpki_invalid", "rpki_unknown"]}}, {"name": "sortOrder", "in": "query", "description": "Sort order.", "schema": {"description": "Sort order.", "type": "string", "example": "desc", "enum": ["ASC", "DESC"]}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"asns": {"type": "array", "items": {"properties": {"asn": {"type": "integer"}, "coneSize": {"description": "AS's customer cone size.", "type": "integer"}, "country": {"description": "Alpha-2 code for the AS's registration country.", "type": "string", "example": "US"}, "ipv4Count": {"description": "Number of IPv4 addresses originated by the AS.", "type": "integer"}, "ipv6Count": {"description": "Number of IPv6 addresses originated by the AS.", "type": "string", "example": "1.21e24"}, "name": {"description": "Name of the AS.", "type": "string"}, "pfxsCount": {"description": "Number of total IP prefixes originated by the AS.", "type": "integer"}, "rpkiInvalid": {"description": "Number of RPKI invalid prefixes originated by the AS.", "type": "integer"}, "rpkiUnknown": {"description": "Number of RPKI unknown prefixes originated by the AS.", "type": "integer"}, "rpkiValid": {"description": "Number of RPKI valid prefixes originated by the AS.", "type": "integer"}}, "required": ["asn", "name", "coneSize", "country", "ipv4Count", "ipv6Count", "pfxsCount", "rpkiValid", "rpkiInvalid", "rpkiUnknown"], "type": "object"}}, "meta": {"type": "object", "properties": {"dataTime": {"description": "The timestamp of when the data is generated.", "type": "string", "example": "2024-06-03T14:00:00"}, "queryTime": {"description": "The timestamp of the query.", "type": "string", "example": "2024-06-03T14:00:00"}, "totalPeers": {"description": "Total number of route collector peers used to generate this data.", "type": "integer"}}, "required": ["dataTime", "queryTime", "totalPeers"]}}, "required": ["asns", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.routes", "x-fern-sdk-method-name": "ases", "x-forge-hidden": true}
```
