---
title: Get AS-level relationships by ASN
page_id: operation-get-radar-entities-asns-asn-rel-28d4f051
path: operations/radar-autonomous-systems
description: Retrieves AS-level relationship for given networks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/entities/asns/{asn}/rel
operation_ids:
    - radar-get-asns-rel
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get AS-level relationships by ASN

`GET /radar/entities/asns/{asn}/rel`

Operation ID: `radar-get-asns-rel`

Retrieves AS-level relationship for given networks.

## Definition

```yaml
{"operationId": "radar-get-asns-rel", "summary": "Get AS-level relationships by ASN", "description": "Retrieves AS-level relationship for given networks.", "parameters": [{"name": "asn", "in": "path", "description": "Retrieves all ASNs with provider-customer or peering relationships with the given ASN.", "required": true, "schema": {"description": "Retrieves all ASNs with provider-customer or peering relationships with the given ASN.", "type": "integer", "example": 3}}, {"name": "asn2", "in": "query", "description": "Retrieves the AS relationship of ASN2 with respect to the given ASN.", "schema": {"description": "Retrieves the AS relationship of ASN2 with respect to the given ASN.", "type": "integer"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"meta": {"type": "object", "properties": {"data_time": {"type": "string"}, "query_time": {"type": "string"}, "total_peers": {"type": "integer"}}, "required": ["data_time", "query_time", "total_peers"]}, "rels": {"type": "array", "items": {"properties": {"asn1": {"type": "integer"}, "asn1_country": {"type": "string"}, "asn1_name": {"type": "string"}, "asn2": {"type": "integer"}, "asn2_country": {"type": "string"}, "asn2_name": {"type": "string"}, "rel": {"type": "string"}}, "required": ["asn1", "asn1_country", "asn1_name", "asn2", "asn2_country", "asn2_name", "rel"], "type": "object"}}}, "required": ["rels", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Autonomous Systems"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.entities.asns", "x-fern-sdk-method-name": "rel", "x-forge-hidden": true}
```
