---
title: Get dataset download URL
page_id: operation-post-radar-datasets-download-9508608e
path: operations/radar-datasets
description: Retrieves an URL to download a single dataset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /radar/datasets/download
operation_ids:
    - radar-post-reports-dataset-download-url
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get dataset download URL

`POST /radar/datasets/download`

Operation ID: `radar-post-reports-dataset-download-url`

Retrieves an URL to download a single dataset.

## Definition

```yaml
{"operationId": "radar-post-reports-dataset-download-url", "summary": "Get dataset download URL", "description": "Retrieves an URL to download a single dataset.", "parameters": [{"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"datasetId": {"type": "integer", "example": 3, "x-auditable": true}}, "required": ["datasetId"]}}}}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"dataset": {"type": "object", "properties": {"url": {"type": "string", "example": "https://example.com/download"}}, "required": ["url"]}}, "required": ["dataset"]}}, "required": ["result"]}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Datasets"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.datasets", "x-fern-sdk-method-name": "download", "x-forge-hidden": true}
```
