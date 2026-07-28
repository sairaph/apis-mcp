---
title: Get real-time BGP routes for a prefix
page_id: operation-get-radar-bgp-routes-realtime-3fa837a6
path: operations/radar-bgp
description: Retrieves real-time BGP routes for a prefix, using public real-time data collectors (RouteViews and RIPE RIS).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/routes/realtime
operation_ids:
    - radar-get-bgp-routes-realtime
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get real-time BGP routes for a prefix

`GET /radar/bgp/routes/realtime`

Operation ID: `radar-get-bgp-routes-realtime`

Retrieves real-time BGP routes for a prefix, using public real-time data collectors (RouteViews and RIPE RIS).

## Definition

```yaml
{"operationId": "radar-get-bgp-routes-realtime", "summary": "Get real-time BGP routes for a prefix", "description": "Retrieves real-time BGP routes for a prefix, using public real-time data collectors (RouteViews and RIPE RIS).", "parameters": [{"name": "prefix", "in": "query", "schema": {"type": "string", "example": "1.1.1.0/24"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"meta": {"type": "object", "properties": {"asn_info": {"type": "array", "items": {"properties": {"as_name": {"description": "Name of the autonomous system.", "type": "string"}, "asn": {"description": "AS number.", "type": "integer"}, "country_code": {"description": "Alpha-2 code for the AS's registration country.", "type": "string"}, "org_id": {"description": "Organization ID.", "type": "string"}, "org_name": {"description": "Organization name.", "type": "string"}}, "required": ["asn", "as_name", "country_code", "org_id", "org_name"], "type": "object"}}, "collectors": {"type": "array", "items": {"properties": {"collector": {"description": "Public route collector ID.", "type": "string"}, "latest_realtime_ts": {"description": "Latest real-time stream timestamp for this collector.", "type": "string"}, "latest_rib_ts": {"description": "Latest RIB dump MRT file timestamp for this collector.", "type": "string"}, "latest_updates_ts": {"description": "Latest BGP updates MRT file timestamp for this collector.", "type": "string"}, "peers_count": {"description": "Total number of collector peers used from this collector.", "type": "integer"}, "peers_v4_count": {"description": "Total number of collector peers used from this collector for IPv4 prefixes.", "type": "integer"}, "peers_v6_count": {"description": "Total number of collector peers used from this collector for IPv6 prefixes.", "type": "integer"}}, "required": ["collector", "latest_realtime_ts", "latest_updates_ts", "latest_rib_ts", "peers_count", "peers_v4_count", "peers_v6_count"], "type": "object"}}, "data_time": {"description": "The most recent data timestamp for from the real-time sources.", "type": "string"}, "prefix_origins": {"type": "array", "items": {"properties": {"origin": {"description": "Origin ASN.", "type": "integer"}, "prefix": {"description": "IP prefix of this query.", "type": "string"}, "rpki_validation": {"description": "Prefix-origin RPKI validation: valid, invalid, unknown.", "type": "string"}, "total_peers": {"description": "Total number of peers.", "type": "integer"}, "total_visible": {"description": "Total number of peers seeing this prefix.", "type": "integer"}, "visibility": {"description": "Ratio of peers seeing this prefix to total number of peers.", "type": "number"}}, "required": ["origin", "prefix", "rpki_validation", "total_peers", "total_visible", "visibility"], "type": "object"}}, "query_time": {"description": "The timestamp of this query.", "type": "string"}}, "required": ["collectors", "asn_info", "prefix_origins", "data_time", "query_time"]}, "routes": {"type": "array", "items": {"properties": {"as_path": {"description": "AS-level path for this route, from collector to origin.", "type": "array", "items": {"type": "integer"}}, "collector": {"description": "Public collector ID for this route.", "type": "string"}, "communities": {"description": "BGP community values.", "type": "array", "items": {"type": "string"}}, "prefix": {"description": "IP prefix of this query.", "type": "string"}, "timestamp": {"description": "Latest timestamp of change for this route.", "type": "string"}}, "required": ["timestamp", "collector", "prefix", "as_path", "communities"], "type": "object"}}}, "required": ["routes", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.routes", "x-fern-sdk-method-name": "realtime", "x-forge-hidden": true}
```
