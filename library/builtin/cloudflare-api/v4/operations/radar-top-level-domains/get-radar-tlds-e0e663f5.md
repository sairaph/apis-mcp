---
title: List TLDs
page_id: operation-get-radar-tlds-62f3a887
path: operations/radar-top-level-domains
description: Retrieves a list of TLDs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/tlds
operation_ids:
    - radar-get-tlds
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List TLDs

`GET /radar/tlds`

Operation ID: `radar-get-tlds`

Retrieves a list of TLDs.

## Definition

```yaml
{"operationId": "radar-get-tlds", "summary": "List TLDs", "description": "Retrieves a list of TLDs.", "parameters": [{"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "offset", "in": "query", "description": "Skips the specified number of objects before fetching the results.", "schema": {"description": "Skips the specified number of objects before fetching the results.", "type": "integer", "minimum": 0}}, {"name": "tldManager", "in": "query", "description": "Filters results by TLD manager.", "schema": {"description": "Filters results by TLD manager.", "type": "string", "maxLength": 100}}, {"name": "tldType", "in": "query", "description": "Filters results by TLD type.", "schema": {"description": "Filters results by TLD type.", "type": "string", "enum": ["GENERIC", "COUNTRY_CODE", "GENERIC_RESTRICTED", "INFRASTRUCTURE", "SPONSORED"]}}, {"name": "tld", "in": "query", "description": "Filters results by top-level domain. Specify a comma-separated list of TLDs.", "schema": {"description": "Filters results by top-level domain. Specify a comma-separated list of TLDs.", "type": "string", "example": "com,pt"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"tlds": {"type": "array", "items": {"properties": {"manager": {"description": "The organization that manages the TLD.", "type": "string", "example": "VeriSign Global Registry Services"}, "tld": {"description": "The actual TLD.", "type": "string", "example": "com"}, "type": {"description": "The type of TLD.", "type": "string", "example": "GENERIC"}}, "required": ["tld", "type", "manager"], "type": "object"}}}, "required": ["tlds"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Top-Level Domains"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.tlds", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
