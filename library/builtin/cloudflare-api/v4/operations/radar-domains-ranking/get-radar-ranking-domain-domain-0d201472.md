---
title: Get domain rank details
page_id: operation-get-radar-ranking-domain-domain-52b7684b
path: operations/radar-domains-ranking
description: Retrieves domain rank details. Cloudflare provides an ordered rank for the top 100 domains, but for the remainder it only provides ranking buckets like top 200 thousand, top one million, etc.. These are available through Radar datasets endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/ranking/domain/{domain}
operation_ids:
    - radar-get-ranking-domain-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get domain rank details

`GET /radar/ranking/domain/{domain}`

Operation ID: `radar-get-ranking-domain-details`

Retrieves domain rank details. Cloudflare provides an ordered rank for the top 100 domains, but for the remainder it only provides ranking buckets like top 200 thousand, top one million, etc.. These are available through Radar datasets endpoints.

## Definition

```yaml
{"operationId": "radar-get-ranking-domain-details", "summary": "Get domain rank details", "description": "Retrieves domain rank details. Cloudflare provides an ordered rank for the top 100 domains, but for the remainder it only provides ranking buckets like top 200 thousand, top one million, etc.. These are available through Radar datasets endpoints.", "parameters": [{"name": "domain", "in": "path", "description": "Domain name.", "required": true, "schema": {"description": "Domain name.", "type": "string", "example": "google.com", "pattern": "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9-]*[A-Za-z0-9])$"}}, {"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "rankingType", "in": "query", "description": "The ranking type.", "schema": {"description": "The ranking type.", "type": "string", "example": "POPULAR", "default": "POPULAR", "enum": ["POPULAR", "TRENDING_RISE", "TRENDING_STEADY"]}}, {"name": "name", "in": "query", "description": "Array of names used to label the series in the response.", "schema": {"description": "Array of names used to label the series in the response.", "type": "array", "items": {"example": "main_series", "type": "string"}}}, {"name": "includeTopLocations", "in": "query", "description": "Includes top locations in the response.", "schema": {"description": "Includes top locations in the response.", "type": "boolean"}}, {"name": "date", "in": "query", "description": "Filters results by the specified array of dates.", "schema": {"description": "Filters results by the specified array of dates.", "type": "array", "items": {"format": "date", "type": "string"}, "example": "2022-09-19"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"details_0": {"type": "object", "properties": {"bucket": {"description": "Only available in POPULAR ranking for the most recent ranking.", "type": "string", "example": "2000"}, "categories": {"type": "array", "items": {"properties": {"id": {"type": "integer", "example": 81}, "name": {"type": "string", "example": "Content Servers"}, "superCategoryId": {"type": "integer", "example": 26}}, "required": ["superCategoryId", "name", "id"], "type": "object"}}, "rank": {"type": "integer", "example": 3}, "top_locations": {"type": "array", "items": {"properties": {"locationCode": {"type": "string", "example": "US"}, "locationName": {"type": "string", "example": "United States"}, "rank": {"type": "integer", "example": 1}}, "required": ["rank", "locationName", "locationCode"], "type": "object"}}}, "required": ["categories"]}, "meta": {"type": "object", "properties": {"dateRange": {"type": "array", "items": {"properties": {"endTime": {"description": "Adjusted end of date range.", "type": "string", "format": "date-time", "example": "2022-09-17T10:22:57.555Z"}, "startTime": {"description": "Adjusted start of date range.", "type": "string", "format": "date-time", "example": "2022-09-16T10:22:57.555Z"}}, "required": ["startTime", "endTime"], "type": "object"}}}, "required": ["dateRange"]}}, "required": ["details_0", "meta"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Domains Ranking"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.ranking.domain", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
