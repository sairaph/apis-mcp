---
title: Reads an indicator
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-dataset-id-indicat-cc9ae627
path: operations/indicator
description: Retrieves a specific indicator by its UUID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/{indicator_id}
operation_ids:
    - get_IndicatorRead
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reads an indicator

`GET /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators/{indicator_id}`

Operation ID: `get_IndicatorRead`

Retrieves a specific indicator by its UUID.

## Definition

```yaml
{"operationId": "get_IndicatorRead", "summary": "Reads an indicator", "description": "Retrieves a specific indicator by its UUID.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset ID.", "required": true, "schema": {"description": "Dataset ID.", "type": "string"}}, {"name": "indicator_id", "in": "path", "description": "Indicator UUID.", "required": true, "schema": {"description": "Indicator UUID.", "type": "string"}}], "responses": {"200": {"description": "Returns the indicator.", "content": {"application/json": {"schema": {"type": "object", "properties": {"createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "datasetId": {"description": "The dataset ID this indicator belongs to. Included in list responses.", "type": "string", "example": "dataset-uuid-123"}, "indicatorType": {"type": "string", "example": "domain"}, "relatedEvents": {"type": "array", "items": {"properties": {"datasetId": {"type": "string", "example": "dataset-uuid-123"}, "eventId": {"type": "string", "example": "event-uuid-456"}}, "required": ["datasetId", "eventId"], "type": "object"}}, "tags": {"type": "array", "items": {"properties": {"categoryName": {"type": "string"}, "uuid": {"type": "string"}, "value": {"type": "string"}}, "type": "object"}}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}, "value": {"type": "string", "example": "malicious-domain.com"}}, "required": ["uuid", "indicatorType", "value", "createdAt", "updatedAt"]}}}}, "404": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Indicator"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
