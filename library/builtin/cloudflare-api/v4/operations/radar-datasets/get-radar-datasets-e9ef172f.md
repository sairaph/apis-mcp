---
title: List datasets
page_id: operation-get-radar-datasets-a93845fa
path: operations/radar-datasets
description: Retrieves a list of datasets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/datasets
operation_ids:
    - radar-get-reports-datasets
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List datasets

`GET /radar/datasets`

Operation ID: `radar-get-reports-datasets`

Retrieves a list of datasets.

## Definition

```yaml
{"operationId": "radar-get-reports-datasets", "summary": "List datasets", "description": "Retrieves a list of datasets.", "parameters": [{"name": "limit", "in": "query", "description": "Limits the number of objects returned in the response.", "schema": {"description": "Limits the number of objects returned in the response.", "type": "integer", "example": 5, "default": 5, "exclusiveMinimum": true, "minimum": 0}}, {"name": "offset", "in": "query", "description": "Skips the specified number of objects before fetching the results.", "schema": {"description": "Skips the specified number of objects before fetching the results.", "type": "integer", "minimum": 0}}, {"name": "datasetType", "in": "query", "description": "Filters results by dataset type.", "schema": {"description": "Filters results by dataset type.", "type": "string", "example": "RANKING_BUCKET", "default": "RANKING_BUCKET", "enum": ["RANKING_BUCKET", "REPORT"]}}, {"name": "date", "in": "query", "description": "Filters results by the specified date.", "schema": {"description": "Filters results by the specified date.", "type": "string", "format": "date", "example": "2024-09-19"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"datasets": {"type": "array", "items": {"properties": {"description": {"type": "string", "example": "This dataset contains a list of the op 20000 domains globally"}, "id": {"type": "integer", "example": 3}, "meta": {"type": "object"}, "tags": {"type": "array", "items": {"example": "global", "type": "string"}}, "title": {"type": "string", "example": "Top bucket 20000 domains"}, "type": {"type": "string", "example": "RANKING_BUCKET"}}, "required": ["id", "title", "description", "type", "tags", "meta"], "type": "object"}}}, "required": ["datasets"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Datasets"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.datasets", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
