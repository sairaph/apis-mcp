---
title: List mirrored tags for an indicator dataset
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-dataset-id-indicat-a0d2c0af
path: operations/indicator
description: Returns all mirrored tags from the indicator dataset (DO mirror table). No pagination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/tags
operation_ids:
    - get_IndicatorTagsList
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List mirrored tags for an indicator dataset

`GET /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/tags`

Operation ID: `get_IndicatorTagsList`

Returns all mirrored tags from the indicator dataset (DO mirror table). No pagination.

## Definition

```yaml
{"operationId": "get_IndicatorTagsList", "summary": "List mirrored tags for an indicator dataset", "description": "Returns all mirrored tags from the indicator dataset (DO mirror table). No pagination.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset ID.", "required": true, "schema": {"description": "Dataset ID.", "type": "string"}}], "responses": {"200": {"description": "Returns an array of mirrored tags.", "content": {"application/json": {"schema": {"description": "Array of mirror tag rows", "type": "array", "items": {"type": "object"}}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "500": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Indicator"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
