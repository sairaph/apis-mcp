---
title: Get BGP routing table stats
page_id: operation-get-radar-bgp-routes-stats-a5211ebd
path: operations/radar-bgp
description: Retrieves the BGP routing table stats.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/routes/stats
operation_ids:
    - radar-get-bgp-routes-stats
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get BGP routing table stats

`GET /radar/bgp/routes/stats`

Operation ID: `radar-get-bgp-routes-stats`

Retrieves the BGP routing table stats.

## Definition

```yaml
{"operationId": "radar-get-bgp-routes-stats", "summary": "Get BGP routing table stats ", "description": "Retrieves the BGP routing table stats.", "parameters": [{"name": "asn", "in": "query", "description": "Filters results by Autonomous System. Specify a single Autonomous System Number (ASN) as integer.", "schema": {"description": "Filters results by Autonomous System. Specify a single Autonomous System Number (ASN) as integer.", "type": "integer", "example": 174}}, {"name": "location", "in": "query", "description": "Filters results by location. Specify an alpha-2 location code.", "schema": {"description": "Filters results by location. Specify an alpha-2 location code.", "type": "string", "example": "US", "maxLength": 2, "minLength": 2}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"meta": {"type": "object", "properties": {"data_time": {"type": "string"}, "query_time": {"type": "string"}, "total_peers": {"type": "integer"}}, "required": ["data_time", "query_time", "total_peers"]}, "stats": {"type": "object", "properties": {"distinct_origins": {"type": "integer"}, "distinct_origins_ipv4": {"type": "integer"}, "distinct_origins_ipv6": {"type": "integer"}, "distinct_prefixes": {"type": "integer"}, "distinct_prefixes_ipv4": {"type": "integer"}, "distinct_prefixes_ipv6": {"type": "integer"}, "routes_invalid": {"type": "integer"}, "routes_invalid_ipv4": {"type": "integer"}, "routes_invalid_ipv6": {"type": "integer"}, "routes_total": {"type": "integer"}, "routes_total_ipv4": {"type": "integer"}, "routes_total_ipv6": {"type": "integer"}, "routes_unknown": {"type": "integer"}, "routes_unknown_ipv4": {"type": "integer"}, "routes_unknown_ipv6": {"type": "integer"}, "routes_valid": {"type": "integer"}, "routes_valid_ipv4": {"type": "integer"}, "routes_valid_ipv6": {"type": "integer"}}, "required": ["distinct_origins", "distinct_origins_ipv4", "distinct_origins_ipv6", "distinct_prefixes", "distinct_prefixes_ipv4", "distinct_prefixes_ipv6", "routes_invalid", "routes_invalid_ipv4", "routes_invalid_ipv6", "routes_total", "routes_total_ipv4", "routes_total_ipv6", "routes_unknown", "routes_unknown_ipv4", "routes_unknown_ipv6", "routes_valid", "routes_valid_ipv4", "routes_valid_ipv6"]}}, "required": ["stats", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.routes", "x-fern-sdk-method-name": "stats", "x-forge-hidden": true}
```
