---
title: Lists indicators
page_id: operation-get-accounts-account-id-cloudforce-one-events-dataset-dataset-id-indicat-a545afdf
path: operations/indicator
description: This method is deprecated. Please use /events/indicators to retrieve a paginated list of indicators.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators
operation_ids:
    - get_IndicatorListLegacy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Lists indicators

`GET /accounts/{account_id}/cloudforce-one/events/dataset/{dataset_id}/indicators`

Operation ID: `get_IndicatorListLegacy`

This method is deprecated. Please use /events/indicators to retrieve a paginated list of indicators.

## Definition

```yaml
{"operationId": "get_IndicatorListLegacy", "summary": "Lists indicators", "description": "This method is deprecated. Please use /events/indicators to retrieve a paginated list of indicators.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "dataset_id", "in": "path", "description": "Dataset UUID.", "required": true, "schema": {"description": "Dataset UUID.", "type": "string", "format": "uuid"}}, {"name": "page", "in": "query", "schema": {"type": "number"}}, {"name": "pageSize", "in": "query", "schema": {"type": "number"}}, {"name": "name", "in": "query", "description": "Filter by indicator value (substring match)", "schema": {"description": "Filter by indicator value (substring match)", "type": "string"}}, {"name": "indicatorType", "in": "query", "schema": {"type": "string"}}, {"name": "relatedEvent", "in": "query", "description": "Filter indicators by related event UUID(s). Multiple UUIDs can be provided by repeating the parameter.", "schema": {"description": "Filter indicators by related event UUID(s). Multiple UUIDs can be provided by repeating the parameter.", "type": "array", "items": {"type": "string"}}}], "responses": {"200": {"description": "Returns a list of indicators.", "content": {"application/json": {"schema": {"type": "object", "properties": {"indicators": {"type": "array", "items": {"properties": {"createdAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "datasetId": {"description": "The dataset ID this indicator belongs to. Included in list responses.", "type": "string", "example": "dataset-uuid-123"}, "indicatorType": {"type": "string", "example": "domain"}, "relatedEvents": {"type": "array", "items": {"properties": {"datasetId": {"type": "string", "example": "dataset-uuid-123"}, "eventId": {"type": "string", "example": "event-uuid-456"}}, "required": ["datasetId", "eventId"], "type": "object"}}, "tags": {"type": "array", "items": {"properties": {"categoryName": {"type": "string"}, "uuid": {"type": "string"}, "value": {"type": "string"}}, "type": "object"}}, "updatedAt": {"type": "string", "format": "date-time", "example": "2022-04-01T00:00:00Z"}, "uuid": {"type": "string", "example": "12345678-1234-1234-1234-1234567890ab"}, "value": {"type": "string", "example": "malicious-domain.com"}}, "required": ["uuid", "indicatorType", "value", "createdAt", "updatedAt"], "type": "object"}}, "pagination": {"type": "object", "properties": {"page": {"type": "number"}, "pageSize": {"type": "number"}, "totalCount": {"type": "number"}, "totalPages": {"type": "number"}}, "required": ["page", "pageSize", "totalCount", "totalPages"]}}, "required": ["indicators", "pagination"]}}}}}, "deprecated": true, "security": [{"api_token": []}], "tags": ["Indicator"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
