---
title: List Origins
page_id: operation-get-radar-origins-cc60e3ad
path: operations/radar-origins
description: Retrieves a list of origins with their regions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/origins
operation_ids:
    - radar-get-origins
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Origins

`GET /radar/origins`

Operation ID: `radar-get-origins`

Retrieves a list of origins with their regions.

## Definition

```yaml
{"operationId": "radar-get-origins", "summary": "List Origins", "description": "Retrieves a list of origins with their regions.", "parameters": [{"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "offset", "in": "query", "description": "Skips the specified number of objects before fetching the results.", "schema": {"description": "Skips the specified number of objects before fetching the results.", "type": "integer", "minimum": 0}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"origins": {"type": "array", "items": {"properties": {"regions": {"type": "array", "items": {"properties": {"region": {"description": "The region code.", "type": "string", "example": "us-east-1"}}, "required": ["region"], "type": "object"}}, "slug": {"description": "The origin slug.", "type": "string", "example": "amazon"}}, "required": ["slug", "regions"], "type": "object"}}}, "required": ["origins"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Origins"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.origins", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
