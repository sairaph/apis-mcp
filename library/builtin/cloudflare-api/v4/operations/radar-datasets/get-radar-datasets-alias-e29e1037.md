---
title: Get dataset CSV stream
page_id: operation-get-radar-datasets-alias-c6681d55
path: operations/radar-datasets
description: Retrieves the CSV content of a given dataset by alias or ID. When getting the content by alias the latest dataset is returned, optionally filtered by the latest available at a given date.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/datasets/{alias}
operation_ids:
    - radar-get-reports-dataset-download
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get dataset CSV stream

`GET /radar/datasets/{alias}`

Operation ID: `radar-get-reports-dataset-download`

Retrieves the CSV content of a given dataset by alias or ID. When getting the content by alias the latest dataset is returned, optionally filtered by the latest available at a given date.

## Definition

```yaml
{"operationId": "radar-get-reports-dataset-download", "summary": "Get dataset CSV stream", "description": "Retrieves the CSV content of a given dataset by alias or ID. When getting the content by alias the latest dataset is returned, optionally filtered by the latest available at a given date.", "parameters": [{"name": "alias", "in": "path", "description": "Dataset alias or ID.", "required": true, "schema": {"description": "Dataset alias or ID.", "type": "string", "example": "ranking_top_1000"}}], "responses": {"200": {"description": "Successful response.", "content": {"text/csv": {"schema": {"type": "string"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean", "example": false}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Datasets"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.datasets", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
