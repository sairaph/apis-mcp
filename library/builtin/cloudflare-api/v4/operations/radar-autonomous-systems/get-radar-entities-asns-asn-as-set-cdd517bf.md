---
title: Get IRR AS-SETs that an AS is a member of
page_id: operation-get-radar-entities-asns-asn-as-set-e10eb8b5
path: operations/radar-autonomous-systems
description: Retrieves Internet Routing Registry AS-SETs that an AS is a member of.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/entities/asns/{asn}/as_set
operation_ids:
    - radar-get-asns-as-set
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get IRR AS-SETs that an AS is a member of

`GET /radar/entities/asns/{asn}/as_set`

Operation ID: `radar-get-asns-as-set`

Retrieves Internet Routing Registry AS-SETs that an AS is a member of.

## Definition

```yaml
{"operationId": "radar-get-asns-as-set", "summary": "Get IRR AS-SETs that an AS is a member of", "description": "Retrieves Internet Routing Registry AS-SETs that an AS is a member of.", "parameters": [{"name": "asn", "in": "path", "description": "Retrieves all AS-SETs that the given AS is a member of.", "required": true, "schema": {"description": "Retrieves all AS-SETs that the given AS is a member of.", "type": "integer", "example": 3}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"as_sets": {"type": "array", "items": {"properties": {"as_members_count": {"description": "The number of AS members in the AS-SET", "type": "integer"}, "as_set_members_count": {"description": "The number of AS-SET members in the AS-SET", "type": "integer"}, "as_set_upstreams_count": {"description": "The number of recursive upstream AS-SETs", "type": "integer"}, "asn_cone_size": {"description": "The number of unique ASNs in the AS-SETs recursive downstream", "type": "integer"}, "hierarchical_asn": {"description": "The AS number following hierarchical AS-SET name", "type": "integer"}, "inferred_asn": {"description": "The inferred AS number of the AS-SET", "type": "integer"}, "irr_sources": {"description": "The IRR sources of the AS-SET", "type": "array", "items": {"type": "string"}}, "name": {"description": "The name of the AS-SET", "type": "string"}, "peeringdb_asn": {"description": "The AS number matching PeeringDB record", "type": "integer"}}, "required": ["name", "as_members_count", "as_set_members_count", "irr_sources", "asn_cone_size", "as_set_upstreams_count"], "type": "object"}}, "paths": {"description": "Paths from the AS-SET that include the given AS to its upstreams recursively", "type": "array", "items": {"items": {"type": "string"}, "type": "array"}}}, "required": ["as_sets", "paths"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Autonomous Systems"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.entities.asns", "x-fern-sdk-method-name": "as-set", "x-forge-hidden": true}
```
