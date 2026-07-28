---
title: Get Multi-Origin AS (MOAS) prefixes
page_id: operation-get-radar-bgp-routes-moas-151a80db
path: operations/radar-bgp
description: Retrieves all Multi-Origin AS (MOAS) prefixes in the global routing tables.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/routes/moas
operation_ids:
    - radar-get-bgp-pfx2as-moas
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Multi-Origin AS (MOAS) prefixes

`GET /radar/bgp/routes/moas`

Operation ID: `radar-get-bgp-pfx2as-moas`

Retrieves all Multi-Origin AS (MOAS) prefixes in the global routing tables.

## Definition

```yaml
{"operationId": "radar-get-bgp-pfx2as-moas", "summary": "Get Multi-Origin AS (MOAS) prefixes", "description": "Retrieves all Multi-Origin AS (MOAS) prefixes in the global routing tables.", "parameters": [{"name": "origin", "in": "query", "description": "Lookup MOASes originated by the given ASN.", "schema": {"description": "Lookup MOASes originated by the given ASN.", "type": "integer"}}, {"name": "prefix", "in": "query", "schema": {"type": "string", "example": "1.1.1.0/24"}}, {"name": "invalid_only", "in": "query", "description": "Lookup only RPKI invalid MOASes.", "schema": {"description": "Lookup only RPKI invalid MOASes.", "type": "boolean"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"meta": {"type": "object", "properties": {"data_time": {"type": "string"}, "query_time": {"type": "string"}, "total_peers": {"type": "integer"}}, "required": ["data_time", "query_time", "total_peers"]}, "moas": {"type": "array", "items": {"properties": {"origins": {"type": "array", "items": {"properties": {"origin": {"type": "integer"}, "peer_count": {"type": "integer"}, "rpki_validation": {"type": "string"}}, "required": ["origin", "peer_count", "rpki_validation"], "type": "object"}}, "prefix": {"type": "string"}}, "required": ["prefix", "origins"], "type": "object"}}}, "required": ["moas", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.routes", "x-fern-sdk-method-name": "moas", "x-forge-hidden": true}
```
