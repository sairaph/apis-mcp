---
title: Get ASPA objects snapshot
page_id: operation-get-radar-bgp-rpki-aspa-snapshot-7528b517
path: operations/radar-bgp
description: Retrieves current or historical ASPA (Autonomous System Provider Authorization) objects. ASPA objects define which ASNs are authorized upstream providers for a customer ASN.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/rpki/aspa/snapshot
operation_ids:
    - radar-get-bgp-rpki-aspa-snapshot
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get ASPA objects snapshot

`GET /radar/bgp/rpki/aspa/snapshot`

Operation ID: `radar-get-bgp-rpki-aspa-snapshot`

Retrieves current or historical ASPA (Autonomous System Provider Authorization) objects. ASPA objects define which ASNs are authorized upstream providers for a customer ASN.

## Definition

```yaml
{"operationId": "radar-get-bgp-rpki-aspa-snapshot", "summary": "Get ASPA objects snapshot", "description": "Retrieves current or historical ASPA (Autonomous System Provider Authorization) objects. ASPA objects define which ASNs are authorized upstream providers for a customer ASN.", "parameters": [{"name": "customerAsn", "in": "query", "description": "Filter by customer ASN (the ASN publishing the ASPA object).", "schema": {"description": "Filter by customer ASN (the ASN publishing the ASPA object).", "type": "integer", "example": 13335}}, {"name": "providerAsn", "in": "query", "description": "Filter by provider ASN (an authorized upstream provider in ASPA objects).", "schema": {"description": "Filter by provider ASN (an authorized upstream provider in ASPA objects).", "type": "integer", "example": 174}}, {"name": "date", "in": "query", "description": "Filters results by the specified datetime (ISO 8601).", "schema": {"description": "Filters results by the specified datetime (ISO 8601).", "type": "string", "format": "date-time", "example": "2024-09-19T00:00:00Z"}}, {"name": "includeAsnInfo", "in": "query", "description": "Include ASN metadata (name, country) in response.", "schema": {"description": "Include ASN metadata (name, country) in response.", "type": "boolean"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"asnInfo": {"type": "object", "properties": {"13335": {"type": "object", "properties": {"asn": {"description": "ASN number.", "type": "integer"}, "country": {"description": "Alpha-2 country code.", "type": "string"}, "name": {"description": "AS name.", "type": "string"}}, "required": ["asn", "name", "country"]}}, "required": ["13335"]}, "aspaObjects": {"type": "array", "items": {"properties": {"customerAsn": {"description": "The customer ASN publishing the ASPA object.", "type": "integer"}, "providers": {"type": "array", "items": {"description": "Authorized provider ASNs.", "type": "integer"}}}, "required": ["customerAsn", "providers"], "type": "object"}}, "meta": {"type": "object", "properties": {"dataTime": {"description": "Timestamp of the underlying data.", "type": "string", "format": "date-time"}, "queryTime": {"description": "Timestamp when the query was executed.", "type": "string", "format": "date-time"}, "totalCount": {"description": "Total number of ASPA objects.", "type": "integer"}}, "required": ["dataTime", "queryTime", "totalCount"]}}, "required": ["aspaObjects", "asnInfo", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.rpki.aspa", "x-fern-sdk-method-name": "snapshot", "x-forge-hidden": true}
```
