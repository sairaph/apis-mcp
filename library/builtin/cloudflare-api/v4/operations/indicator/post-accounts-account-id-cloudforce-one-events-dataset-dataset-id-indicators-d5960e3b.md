---
title: Creates multiple indicators in bulk
page_id: operation-post-accounts-account-id-cloudforce-one-events-dataset-dataset-id-indica-6a95c58d
path: operations/indicator
description: Creates multiple indicators at once with their respective types and related datasets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/bulk
operation_ids:
    - post_IndicatorCreateBulk
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates multiple indicators in bulk

`POST /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/bulk`

Operation ID: `post_IndicatorCreateBulk`

Creates multiple indicators at once with their respective types and related datasets.

## Definition

```yaml
{"operationId": "post_IndicatorCreateBulk", "summary": "Creates multiple indicators in bulk", "description": "Creates multiple indicators at once with their respective types and related datasets.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset UUID.", "required": true, "schema": {"description": "Dataset UUID.", "type": "string", "format": "uuid"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"autoCreateType": {"description": "Global flag to automatically create indicator types if they don't exist. Individual indicators can override this with their own autoCreateType flag.", "type": "boolean"}, "indicators": {"type": "array", "items": {"properties": {"autoCreateType": {"description": "If true, automatically create the indicator type if it doesn't exist. If false (default), throw an error when the indicator type doesn't exist.", "type": "boolean"}, "indicatorType": {"type": "string", "example": "domain"}, "relatedEvents": {"type": "array", "items": {"properties": {"datasetId": {"type": "string", "example": "dataset-uuid-123"}, "eventId": {"type": "string", "example": "event-uuid-456"}}, "required": ["datasetId", "eventId"], "type": "object"}}, "tags": {"type": "array", "items": {"anyOf": [{"type": "string"}, {"properties": {"categoryName": {"type": "string"}, "value": {"type": "string"}}, "required": ["value"], "type": "object"}]}}, "value": {"type": "string", "example": "malicious-domain.com"}}, "required": ["indicatorType", "value"], "type": "object"}}}, "required": ["indicators"]}}}}, "responses": {"200": {"description": "Returns the number of created indicators.", "content": {"application/json": {"schema": {"description": "Number of created indicators", "type": "number"}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Indicator"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
