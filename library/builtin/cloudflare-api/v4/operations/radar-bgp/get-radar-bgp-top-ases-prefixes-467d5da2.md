---
title: Get top ASes by prefix count
page_id: operation-get-radar-bgp-top-ases-prefixes-a128752e
path: operations/radar-bgp
description: Retrieves the full list of autonomous systems on the global routing table ordered by announced prefixes count. The data comes from public BGP MRT data archives and updates every 2 hours.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/top/ases/prefixes
operation_ids:
    - radar-get-bgp-top-asns-by-prefixes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get top ASes by prefix count

`GET /radar/bgp/top/ases/prefixes`

Operation ID: `radar-get-bgp-top-asns-by-prefixes`

Retrieves the full list of autonomous systems on the global routing table ordered by announced prefixes count. The data comes from public BGP MRT data archives and updates every 2 hours.

## Definition

```yaml
{"operationId": "radar-get-bgp-top-asns-by-prefixes", "summary": "Get top ASes by prefix count", "description": "Retrieves the full list of autonomous systems on the global routing table ordered by announced prefixes count. The data comes from public BGP MRT data archives and updates every 2 hours.", "parameters": [{"name": "country", "in": "query", "description": "Alpha-2 country code.", "schema": {"description": "Alpha-2 country code.", "type": "string", "example": "NZ"}}, {"name": "limit", "in": "query", "description": "Maximum number of ASes to return.", "schema": {"description": "Maximum number of ASes to return.", "type": "integer", "example": 10}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"asns": {"type": "array", "items": {"properties": {"asn": {"type": "integer"}, "country": {"type": "string"}, "name": {"type": "string"}, "pfxs_count": {"type": "integer"}}, "required": ["asn", "country", "name", "pfxs_count"], "type": "object"}}, "meta": {"type": "object", "properties": {"data_time": {"type": "string"}, "query_time": {"type": "string"}, "total_peers": {"type": "integer"}}, "required": ["data_time", "query_time", "total_peers"]}}, "required": ["asns", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string", "example": "Not Found."}}, "required": ["error"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.top.ases", "x-fern-sdk-method-name": "prefixes", "x-forge-hidden": true}
```
