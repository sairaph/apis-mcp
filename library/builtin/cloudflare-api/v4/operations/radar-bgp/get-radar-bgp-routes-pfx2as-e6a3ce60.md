---
title: Get prefix-to-ASN mapping
page_id: operation-get-radar-bgp-routes-pfx2as-fabb63c4
path: operations/radar-bgp
description: Retrieves the prefix-to-ASN mapping from global routing tables.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/bgp/routes/pfx2as
operation_ids:
    - radar-get-bgp-pfx2as
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get prefix-to-ASN mapping

`GET /radar/bgp/routes/pfx2as`

Operation ID: `radar-get-bgp-pfx2as`

Retrieves the prefix-to-ASN mapping from global routing tables.

## Definition

```yaml
{"operationId": "radar-get-bgp-pfx2as", "summary": "Get prefix-to-ASN mapping", "description": "Retrieves the prefix-to-ASN mapping from global routing tables.", "parameters": [{"name": "prefix", "in": "query", "schema": {"type": "string", "example": "1.1.1.0/24"}}, {"name": "origin", "in": "query", "description": "Lookup prefixes originated by the given ASN.", "schema": {"description": "Lookup prefixes originated by the given ASN.", "type": "integer"}}, {"name": "rpkiStatus", "in": "query", "description": "Return only results with matching rpki status: valid, invalid or unknown.", "schema": {"description": "Return only results with matching rpki status: valid, invalid or unknown.", "type": "string", "example": "INVALID", "enum": ["VALID", "INVALID", "UNKNOWN"]}}, {"name": "longestPrefixMatch", "in": "query", "description": "Return only results with the longest prefix match for the given prefix. For example, specify a /32 prefix to lookup the origin ASN for an IPv4 address.", "schema": {"description": "Return only results with the longest prefix match for the given prefix. For example, specify a /32 prefix to lookup the origin ASN for an IPv4 address.", "type": "boolean", "example": "true"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"meta": {"type": "object", "properties": {"data_time": {"type": "string"}, "query_time": {"type": "string"}, "total_peers": {"type": "integer"}}, "required": ["data_time", "query_time", "total_peers"]}, "prefix_origins": {"type": "array", "items": {"properties": {"origin": {"type": "integer"}, "peer_count": {"type": "integer"}, "prefix": {"type": "string"}, "rpki_validation": {"type": "string"}}, "required": ["origin", "peer_count", "prefix", "rpki_validation"], "type": "object"}}}, "required": ["prefix_origins", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar BGP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.bgp.routes", "x-fern-sdk-method-name": "pfx2as", "x-forge-hidden": true}
```
