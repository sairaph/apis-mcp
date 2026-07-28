---
title: Save query
page_id: operation-post-accounts-account-id-workers-observability-queries-0fb7b208
path: operations/saved-queries
description: Persist query for later use.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/workers/observability/queries
operation_ids:
    - queries.post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Save query

`POST /accounts/{account_id}/workers/observability/queries`

Operation ID: `queries.post`

Persist query for later use.

## Path Parameters

```yaml
[{"name": "account_id", "in": "path", "description": "Your Cloudflare account ID.", "required": true, "schema": {"type": "string"}}]
```

## Definition

```yaml
{"operationId": "queries.post", "summary": "Save query", "description": "Persist query for later use.", "requestBody": {"description": "Specify the new contents of the query.", "required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"type": "string", "example": "Query description", "maxLength": 1000, "nullable": true}, "name": {"description": "Query name", "type": "string", "maxLength": 250, "minLength": 1}, "parameters": {"properties": {"calculations": {"description": "Create Calculations to compute as part of the query.", "type": "array", "items": {"properties": {"alias": {"type": "string"}, "key": {"type": "string"}, "keyType": {"type": "string", "enum": ["string", "number", "boolean"]}, "operator": {"type": "string", "enum": ["uniq", "count", "max", "min", "sum", "avg", "median", "p001", "p01", "p05", "p10", "p25", "p75", "p90", "p95", "p99", "p999", "stddev", "variance", "COUNT_DISTINCT", "COUNT", "MAX", "MIN", "SUM", "AVG", "MEDIAN", "P001", "P01", "P05", "P10", "P25", "P75", "P90", "P95", "P99", "P999", "STDDEV", "VARIANCE"]}}, "required": ["operator"], "type": "object"}}, "datasets": {"description": "Set the Datasets to query. Leave it empty to query all the datasets.", "type": "array", "items": {"type": "string"}, "example": []}, "filterCombination": {"description": "Set a Flag to describe how to combine the filters on the query.", "type": "string", "enum": ["and", "or", "AND", "OR"]}, "filters": {"description": "Configure the Filters to apply to the query. Supports nested groups via kind: 'group'.", "type": "array", "items": {"$ref": "#/components/schemas/workers-observability_filter_node"}}, "groupBys": {"description": "Define how to group the results of the query.", "type": "array", "items": {"properties": {"type": {"type": "string", "enum": ["string", "number", "boolean"]}, "value": {"type": "string"}}, "required": ["type", "value"], "type": "object"}}, "havings": {"description": "Configure the Having clauses that filter on calculations in the query result.", "type": "array", "items": {"properties": {"key": {"type": "string"}, "operation": {"type": "string", "enum": ["eq", "neq", "gt", "gte", "lt", "lte"]}, "value": {"type": "number"}}, "required": ["key", "operation", "value"], "type": "object"}}, "limit": {"description": "Set a limit on the number of results / records returned by the query", "type": "integer", "maximum": 100, "minimum": 0}, "needle": {"description": "Define an expression to search using full-text search.", "type": "object", "properties": {"isRegex": {"type": "boolean"}, "matchCase": {"type": "boolean"}, "value": {"allOf": [{"anyOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}, {"maxLength": 1000, "type": "string"}]}}, "required": ["value"]}, "orderBy": {"description": "Configure the order of the results returned by the query.", "type": "object", "properties": {"order": {"description": "Set the order of the results", "type": "string", "enum": ["asc", "desc"]}, "value": {"description": "Configure which Calculation to order the results by.", "type": "string"}}, "required": ["value"]}}, "type": "object"}}, "required": ["description", "name", "parameters"]}}}}, "responses": {"200": {"description": "Successful request", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string", "enum": ["Successful request"]}}, "required": ["message"], "type": "object"}}, "result": {"$ref": "#/components/schemas/workers-observability_query"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["messages", "success", "errors", "result"]}}}}, "401": {"description": "Unauthorized", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Unauthorized"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "409": {"description": "Conflict", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Conflict"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}, "500": {"description": "Internal error", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "message": {"type": "string", "enum": ["Internal error"]}}, "required": ["message"], "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"message": {"type": "string"}}, "required": ["message"], "type": "object"}}, "success": {"type": "boolean", "enum": [false]}}, "required": ["errors", "success", "messages"]}}}}}, "tags": ["Saved Queries"], "x-api-token-group": ["Workers Observability Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "observability.queries", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
