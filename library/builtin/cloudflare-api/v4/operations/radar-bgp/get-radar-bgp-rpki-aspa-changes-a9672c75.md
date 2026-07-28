---
title: Get ASPA changes over time
page_id: operation-get-radar-bgp-rpki-aspa-changes-f39ed116
path: operations/radar-bgp
description: Retrieves ASPA (Autonomous System Provider Authorization) changes over time. Returns daily aggregated changes including additions, removals, and modifications of ASPA objects.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/rpki/aspa/changes
operation_ids:
    - radar-get-bgp-rpki-aspa-changes
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get ASPA changes over time

`GET /radar/bgp/rpki/aspa/changes`

Operation ID: `radar-get-bgp-rpki-aspa-changes`

Retrieves ASPA (Autonomous System Provider Authorization) changes over time. Returns daily aggregated changes including additions, removals, and modifications of ASPA objects.

## Definition

```yaml
{"operationId": "radar-get-bgp-rpki-aspa-changes", "summary": "Get ASPA changes over time", "description": "Retrieves ASPA (Autonomous System Provider Authorization) changes over time. Returns daily aggregated changes including additions, removals, and modifications of ASPA objects.", "parameters": [{"name": "dateStart", "in": "query", "description": "Start of the date range (inclusive). Alternative to `dateRange`; provide together with `dateEnd`.", "schema": {"description": "Start of the date range (inclusive). Alternative to `dateRange`; provide together with `dateEnd`.", "type": "string", "format": "date-time", "example": "2023-09-01T11:41:33.782Z"}}, {"name": "dateEnd", "in": "query", "description": "End of the date range (inclusive). Alternative to `dateRange`; provide together with `dateStart`.", "schema": {"description": "End of the date range (inclusive). Alternative to `dateRange`; provide together with `dateStart`.", "type": "string", "format": "date-time", "example": "2023-09-01T11:41:33.782Z"}}, {"name": "asn", "in": "query", "description": "Filter changes involving this ASN (as customer or provider).", "schema": {"description": "Filter changes involving this ASN (as customer or provider).", "type": "integer", "example": 13335}}, {"name": "includeAsnInfo", "in": "query", "description": "Include ASN metadata (name, country) in response.", "schema": {"description": "Include ASN metadata (name, country) in response.", "type": "boolean"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"asnInfo": {"type": "object", "properties": {"13335": {"type": "object", "properties": {"asn": {"description": "ASN number.", "type": "integer"}, "country": {"description": "Alpha-2 country code.", "type": "string"}, "name": {"description": "AS name.", "type": "string"}}, "required": ["asn", "name", "country"]}}, "required": ["13335"]}, "changes": {"type": "array", "items": {"properties": {"customersAdded": {"description": "Number of new ASPA objects created.", "type": "integer"}, "customersRemoved": {"description": "Number of ASPA objects deleted.", "type": "integer"}, "date": {"description": "Date of the changes in ISO 8601 format.", "type": "string", "format": "date-time"}, "entries": {"type": "array", "items": {"properties": {"customerAsn": {"description": "The customer ASN affected.", "type": "integer"}, "providers": {"type": "array", "items": {"description": "Provider ASNs involved in the change.", "type": "integer"}}, "type": {"type": "string", "enum": ["CustomerAdded", "CustomerRemoved", "ProvidersAdded", "ProvidersRemoved"]}}, "required": ["type", "customerAsn", "providers"], "type": "object"}}, "providersAdded": {"description": "Number of providers added to existing objects.", "type": "integer"}, "providersRemoved": {"description": "Number of providers removed from existing objects.", "type": "integer"}, "totalCount": {"description": "Running total of active ASPA objects after this day.", "type": "integer"}}, "required": ["date", "customersAdded", "customersRemoved", "providersAdded", "providersRemoved", "totalCount", "entries"], "type": "object"}}, "meta": {"type": "object", "properties": {"dataTime": {"description": "Timestamp of the underlying data.", "type": "string", "format": "date-time"}, "queryTime": {"description": "Timestamp when the query was executed.", "type": "string", "format": "date-time"}}, "required": ["dataTime", "queryTime"]}}, "required": ["changes", "asnInfo", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.rpki.aspa", "x-fern-sdk-method-name": "changes", "x-forge-hidden": true}
```
