---
title: List Internet services categories
page_id: operation-get-radar-ranking-internet-services-categories-b5bd63d4
path: operations/radar-internet-services-ranking
description: Retrieves the list of Internet services categories.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/ranking/internet_services/categories
operation_ids:
    - radar-get-ranking-internet-services-categories
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Internet services categories

`GET /radar/ranking/internet_services/categories`

Operation ID: `radar-get-ranking-internet-services-categories`

Retrieves the list of Internet services categories.

## Definition

```yaml
{"operationId": "radar-get-ranking-internet-services-categories", "summary": "List Internet services categories", "description": "Retrieves the list of Internet services categories.", "parameters": [{"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5}}, {"name": "name", "in": "query", "description": "Array of names used to label the series in the response.", "schema": {"description": "Array of names used to label the series in the response.", "type": "array", "items": {"example": "main_series", "type": "string"}}}, {"name": "date", "in": "query", "description": "Filters results by the specified array of dates.", "schema": {"description": "Filters results by the specified array of dates.", "type": "array", "items": {"format": "date", "type": "string"}, "example": "2022-09-19"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"categories_0": {"type": "array", "items": {"properties": {"name": {"type": "string", "example": "Generative AI"}}, "required": ["name"], "type": "object"}}}, "required": ["categories_0"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Internet Services Ranking"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.ranking.internet-services", "x-fern-sdk-method-name": "categories", "x-forge-hidden": true}
```
