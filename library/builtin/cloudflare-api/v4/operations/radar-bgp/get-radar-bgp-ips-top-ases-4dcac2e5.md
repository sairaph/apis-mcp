---
title: Get top ASes by announced IP space
page_id: operation-get-radar-bgp-ips-top-ases-52fcc33d
path: operations/radar-bgp
description: Returns the top-N autonomous systems by announced IP space at the nearest 8-hour RIB boundary at or before the requested date. The snapped boundary is returned as `anchor_ts`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/ips/top/ases
operation_ids:
    - radar-get-bgp-ips-top-ases
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get top ASes by announced IP space

`GET /radar/bgp/ips/top/ases`

Operation ID: `radar-get-bgp-ips-top-ases`

Returns the top-N autonomous systems by announced IP space at the nearest 8-hour RIB boundary at or before the requested date. The snapped boundary is returned as `anchor_ts`.

## Definition

```yaml
{"operationId": "radar-get-bgp-ips-top-ases", "summary": "Get top ASes by announced IP space", "description": "Returns the top-N autonomous systems by announced IP space at the nearest 8-hour RIB boundary at or before the requested date. The snapped boundary is returned as `anchor_ts`.", "parameters": [{"name": "date", "in": "query", "description": "Filters results by the specified datetime (ISO 8601).", "schema": {"description": "Filters results by the specified datetime (ISO 8601).", "type": "string", "format": "date-time", "example": "2024-09-19T00:00:00Z"}}, {"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "maximum": 50, "minimum": 1}}, {"name": "metric", "in": "query", "description": "Ranking metric: IPv4 /24 count or IPv6 /48 count.", "schema": {"description": "Ranking metric: IPv4 /24 count or IPv6 /48 count.", "type": "string", "example": "v4_24s", "enum": ["v4_24s", "v6_48s"]}}, {"name": "country", "in": "query", "description": "Optional ISO 3166-1 alpha-2 country filter. Omit for global top-N.", "schema": {"description": "Optional ISO 3166-1 alpha-2 country filter. Omit for global top-N.", "type": "string", "example": "US", "maxLength": 2, "minLength": 2}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"anchorTs": {"type": "string", "format": "date-time", "example": "2026-04-18T16:00:00.000Z"}, "asns": {"type": "array", "items": {"properties": {"asn": {"type": "integer", "example": 749}, "v4_24s": {"type": "integer", "example": 875649}, "v6_48s": {"type": "integer"}}, "required": ["asn", "v4_24s", "v6_48s"], "type": "object"}}, "country": {"type": "string", "example": "US", "nullable": true}, "metric": {"type": "string", "example": "v4_24s"}}, "required": ["anchorTs", "metric", "country", "asns"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.ips.top", "x-fern-sdk-method-name": "ases", "x-forge-hidden": true}
```
