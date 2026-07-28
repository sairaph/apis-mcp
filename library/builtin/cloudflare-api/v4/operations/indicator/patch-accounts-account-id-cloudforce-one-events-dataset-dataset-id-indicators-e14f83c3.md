---
title: Updates an indicator
page_id: operation-patch-accounts-account-id-cloudforce-one-events-dataset-dataset-id-indic-acd65e88
path: operations/indicator
description: Updates an existing indicator's properties.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/{indicator_id}
operation_ids:
    - patch_IndicatorUpdate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Updates an indicator

`PATCH /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/{indicator_id}`

Operation ID: `patch_IndicatorUpdate`

Updates an existing indicator's properties.

## Definition

```yaml
{"operationId": "patch_IndicatorUpdate", "summary": "Updates an indicator", "description": "Updates an existing indicator's properties.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset ID.", "required": true, "schema": {"description": "Dataset ID.", "type": "string"}}, {"name": "indicator_id", "in": "path", "description": "Indicator UUID.", "required": true, "schema": {"description": "Indicator UUID.", "type": "string"}}], "requestBody": {"content": {"application/json": {"schema": {"type": "object", "properties": {"indicatorType": {"type": "string", "example": "domain"}, "relatedEvents": {"type": "array", "items": {"properties": {"datasetId": {"type": "string", "example": "dataset-uuid-123"}, "eventId": {"type": "string", "example": "event-uuid-456"}}, "required": ["datasetId", "eventId"], "type": "object"}}, "tags": {"type": "array", "items": {"anyOf": [{"type": "string"}, {"properties": {"categoryName": {"type": "string"}, "value": {"type": "string"}}, "required": ["value"], "type": "object"}]}}, "value": {"type": "string", "example": "updated-domain.com"}}}}}}, "responses": {"200": {"description": "Returns the updated indicator.", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "datasetId": {"description": "The dataset ID this indicator belongs to. Included in list responses.", "type": "string", "example": "dataset-uuid-123"}, "indicatorType": {"type": "string", "example": "domain"}, "relatedEvents": {"type": "array", "items": {"properties": {"datasetId": {"type": "string", "example": "dataset-uuid-123"}, "eventId": {"type": "string", "example": "event-uuid-456"}}, "required": ["datasetId", "eventId"], "type": "object"}}, "tags": {"type": "array", "items": {"properties": {"categoryName": {"type": "string"}, "uuid": {"type": "string"}, "value": {"type": "string"}}, "type": "object"}}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}, "value": {"type": "string", "example": "malicious-domain.com"}}, "required": ["uuid", "indicatorType", "value", "createdAt", "updatedAt"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Indicator"], "x-api-token-group": ["Cloudforce One Write"]}
```
