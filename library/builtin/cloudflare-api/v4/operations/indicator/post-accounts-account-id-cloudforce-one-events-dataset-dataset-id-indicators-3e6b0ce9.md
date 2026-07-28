---
title: Creates a new indicator
page_id: operation-post-accounts-account-id-cloudforce-one-events-dataset-dataset-id-indica-adbb110b
path: operations/indicator
description: Creates a new indicator with the specified type and related datasets.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/create
operation_ids:
    - post_IndicatorCreate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a new indicator

`POST /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/create`

Operation ID: `post_IndicatorCreate`

Creates a new indicator with the specified type and related datasets.

## Definition

```yaml
{"operationId": "post_IndicatorCreate", "summary": "Creates a new indicator", "description": "Creates a new indicator with the specified type and related datasets.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset UUID.", "required": true, "schema": {"description": "Dataset UUID.", "type": "string", "format": "uuid"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"autoCreateType": {"description": "If true, automatically create the indicator type if it doesn't exist. If false (default), throw an error when the indicator type doesn't exist.", "type": "boolean"}, "indicatorType": {"type": "string", "example": "domain"}, "relatedEvents": {"type": "array", "items": {"properties": {"datasetId": {"type": "string", "example": "dataset-uuid-123"}, "eventId": {"type": "string", "example": "event-uuid-456"}}, "required": ["datasetId", "eventId"], "type": "object"}}, "tags": {"type": "array", "items": {"anyOf": [{"type": "string"}, {"properties": {"categoryName": {"type": "string"}, "value": {"type": "string"}}, "required": ["value"], "type": "object"}]}}, "value": {"type": "string", "example": "malicious-domain.com"}}, "required": ["indicatorType", "value"]}}}}, "responses": {"200": {"description": "Returns the created indicator.", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "datasetId": {"description": "The dataset ID this indicator belongs to. Included in list responses.", "type": "string", "example": "dataset-uuid-123"}, "indicatorType": {"type": "string", "example": "domain"}, "relatedEvents": {"type": "array", "items": {"properties": {"datasetId": {"type": "string", "example": "dataset-uuid-123"}, "eventId": {"type": "string", "example": "event-uuid-456"}}, "required": ["datasetId", "eventId"], "type": "object"}}, "tags": {"type": "array", "items": {"properties": {"categoryName": {"type": "string"}, "uuid": {"type": "string"}, "value": {"type": "string"}}, "type": "object"}}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}, "value": {"type": "string", "example": "malicious-domain.com"}}, "required": ["uuid", "indicatorType", "value", "createdAt", "updatedAt"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Indicator"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
