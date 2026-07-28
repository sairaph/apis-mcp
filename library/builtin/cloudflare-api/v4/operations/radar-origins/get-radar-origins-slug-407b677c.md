---
title: Get Origin details
page_id: operation-get-radar-origins-slug-db69e09f
path: operations/radar-origins
description: Retrieves the requested origin information with its regions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/origins/{slug}
operation_ids:
    - radar-get-origin-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Origin details

`GET /radar/origins/{slug}`

Operation ID: `radar-get-origin-details`

Retrieves the requested origin information with its regions.

## Definition

```yaml
{"operationId": "radar-get-origin-details", "summary": "Get Origin details", "description": "Retrieves the requested origin information with its regions.", "parameters": [{"name": "slug", "in": "path", "description": "Origin slug.", "required": true, "schema": {"description": "Origin slug.", "type": "string", "example": "amazon", "enum": ["AMAZON", "GOOGLE", "MICROSOFT", "ORACLE"]}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"origin": {"type": "object", "properties": {"regions": {"type": "array", "items": {"properties": {"region": {"description": "The region code.", "type": "string", "example": "us-east-1"}}, "required": ["region"], "type": "object"}}, "slug": {"description": "The origin slug.", "type": "string", "example": "amazon"}}, "required": ["slug", "regions"]}}, "required": ["origin"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string", "example": "Not Found."}}, "required": ["error"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Origins"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.origins", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
