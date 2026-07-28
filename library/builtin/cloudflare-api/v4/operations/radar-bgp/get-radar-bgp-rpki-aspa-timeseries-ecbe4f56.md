---
title: Get ASPA count time series
page_id: operation-get-radar-bgp-rpki-aspa-timeseries-c9f58146
path: operations/radar-bgp
description: Retrieves ASPA (Autonomous System Provider Authorization) object count over time. Supports filtering by RIR or location (country code) to generate multiple named series. If no RIR or location filter is specified, returns total count.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/rpki/aspa/timeseries
operation_ids:
    - radar-get-bgp-rpki-aspa-timeseries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get ASPA count time series

`GET /radar/bgp/rpki/aspa/timeseries`

Operation ID: `radar-get-bgp-rpki-aspa-timeseries`

Retrieves ASPA (Autonomous System Provider Authorization) object count over time. Supports filtering by RIR or location (country code) to generate multiple named series. If no RIR or location filter is specified, returns total count.

## Definition

```yaml
{"operationId": "radar-get-bgp-rpki-aspa-timeseries", "summary": "Get ASPA count time series", "description": "Retrieves ASPA (Autonomous System Provider Authorization) object count over time. Supports filtering by RIR or location (country code) to generate multiple named series. If no RIR or location filter is specified, returns total count.", "parameters": [{"name": "dateStart", "in": "query", "description": "Start of the date range (inclusive). Alternative to `dateRange`; provide together with `dateEnd`.", "schema": {"description": "Start of the date range (inclusive). Alternative to `dateRange`; provide together with `dateEnd`.", "type": "string", "format": "date-time", "example": "2023-09-01T11:41:33.782Z"}}, {"name": "dateEnd", "in": "query", "description": "End of the date range (inclusive). Alternative to `dateRange`; provide together with `dateStart`.", "schema": {"description": "End of the date range (inclusive). Alternative to `dateRange`; provide together with `dateStart`.", "type": "string", "format": "date-time", "example": "2023-09-01T11:41:33.782Z"}}, {"name": "name", "in": "query", "description": "Array of names used to label the series in the response.", "schema": {"description": "Array of names used to label the series in the response.", "type": "array", "items": {"example": "main_series", "type": "string"}}}, {"name": "rir", "in": "query", "description": "Filter by Regional Internet Registry (RIR). Multiple RIRs generate multiple series.", "schema": {"description": "Filter by Regional Internet Registry (RIR). Multiple RIRs generate multiple series.", "type": "array", "items": {"enum": ["RIPE_NCC", "ARIN", "APNIC", "LACNIC", "AFRINIC"], "type": "string"}, "example": "RIPE_NCC"}}, {"name": "location", "in": "query", "description": "Filters results by location. Specify a comma-separated list of alpha-2 location codes.", "schema": {"description": "Filters results by location. Specify a comma-separated list of alpha-2 location codes.", "type": "array", "items": {"type": "string"}, "example": "US"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"meta": {"type": "object", "properties": {"dataTime": {"description": "Timestamp of the underlying data.", "type": "string", "format": "date-time"}, "queryTime": {"description": "Timestamp when the query was executed.", "type": "string", "format": "date-time"}}, "required": ["dataTime", "queryTime"]}, "serie_0": {"type": "object", "properties": {"timestamps": {"type": "array", "items": {"format": "date-time", "type": "string"}}, "values": {"type": "array", "items": {"description": "A numeric string.", "example": "10", "pattern": "^-?\\d+(\\.\\d+)?$", "type": "string"}}}, "required": ["timestamps", "values"]}}, "required": ["serie_0", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.rpki.aspa", "x-fern-sdk-method-name": "timeseries", "x-forge-hidden": true}
```
