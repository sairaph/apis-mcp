---
title: Aggregate events by single or multiple columns with optional date filtering
page_id: operation-get-accounts-account-id-cloudforce-one-events-aggregate-b31bc40b
path: operations/event
description: Aggregate threat events by one or more columns (e.g., attacker, targetIndustry) with optional date filtering and daily grouping. Supports multi-dimensional aggregation for cross-analysis.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/events/aggregate
operation_ids:
    - get_EventAggregate
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Aggregate events by single or multiple columns with optional date filtering

`GET /accounts/{account_id}/cloudforce-one/events/aggregate`

Operation ID: `get_EventAggregate`

Aggregate threat events by one or more columns (e.g., attacker, targetIndustry) with optional date filtering and daily grouping. Supports multi-dimensional aggregation for cross-analysis.

## Definition

```yaml
{"operationId": "get_EventAggregate", "summary": "Aggregate events by single or multiple columns with optional date filtering", "description": "Aggregate threat events by one or more columns (e.g., attacker, targetIndustry) with optional date filtering and daily grouping. Supports multi-dimensional aggregation for cross-analysis.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "aggregateBy", "in": "query", "description": "Column(s) to aggregate by - single column or comma-separated list (e.g., 'attacker', 'targetIndustry', 'attacker,targetIndustry')", "required": true, "schema": {"description": "Column(s) to aggregate by - single column or comma-separated list (e.g., 'attacker', 'targetIndustry', 'attacker,targetIndustry')", "type": "string"}}, {"name": "datasetId", "in": "query", "description": "Dataset ID(s) to filter by. Can be a single dataset ID, comma-separated list, or array. If not provided, uses default dataset", "schema": {"description": "Dataset ID(s) to filter by. Can be a single dataset ID, comma-separated list, or array. If not provided, uses default dataset", "type": "array", "items": {"type": "string"}}}, {"name": "startDate", "in": "query", "description": "Start date for filtering (ISO 8601 format, e.g., '2024-01-01')", "schema": {"description": "Start date for filtering (ISO 8601 format, e.g., '2024-01-01')", "type": "string"}}, {"name": "endDate", "in": "query", "description": "End date for filtering (ISO 8601 format, e.g., '2024-12-31')", "schema": {"description": "End date for filtering (ISO 8601 format, e.g., '2024-12-31')", "type": "string"}}, {"name": "groupByDate", "in": "query", "description": "Whether to group results by date (daily aggregation)", "schema": {"description": "Whether to group results by date (daily aggregation)", "type": "boolean"}}, {"name": "limit", "in": "query", "description": "Maximum number of results to return", "schema": {"description": "Maximum number of results to return", "type": "number", "default": 100}}], "responses": {"200": {"description": "Returns aggregated event data.", "content": {"application/json": {"schema": {"type": "object", "properties": {"aggregateBy": {"description": "Column(s) that were aggregated by", "type": "string"}, "aggregations": {"description": "Array of aggregation results with dynamic fields based on aggregateBy columns", "type": "array", "items": {"allOf": [{"additionalProperties": {"nullable": true, "type": "string"}, "type": "object"}, {"properties": {"count": {"description": "Number of events for this aggregation", "type": "number"}, "date": {"description": "Date (if groupByDate is true)", "type": "string"}}, "required": ["count"], "type": "object"}]}}, "dateRange": {"description": "Date range used for filtering", "type": "object", "properties": {"endDate": {"type": "string"}, "startDate": {"type": "string"}}}, "total": {"description": "Total number of events in the aggregation", "type": "number"}}, "required": ["aggregations", "total", "aggregateBy"]}}}}, "400": {"description": "Bad Request.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string", "example": "An error occurred."}}, "required": ["message"], "type": "object"}}, "result": {"type": "object"}, "success": {"type": "boolean"}}, "required": ["result", "success", "errors"]}}}}}, "security": [{"api_token": []}], "tags": ["Event"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
