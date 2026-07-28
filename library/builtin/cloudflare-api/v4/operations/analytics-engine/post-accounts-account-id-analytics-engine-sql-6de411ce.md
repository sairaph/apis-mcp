---
title: Execute an Analytics Engine SQL query via request body
page_id: operation-post-accounts-account-id-analytics-engine-sql-6528088f
path: operations/analytics-engine
description: Executes a SQL query against Workers Analytics Engine data. Pass the SQL query in the request body as plain text. The response uses newline-delimited JSON (NDJSON) by default, or a single JSON object when the query includes a FORMAT JSON clause. Prefer this method for longer queries that may exceed URL length limits.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/analytics_engine/sql
operation_ids:
    - analytics-engine-sql-query-post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Execute an Analytics Engine SQL query via request body

`POST /accounts/{account_id}/analytics_engine/sql`

Operation ID: `analytics-engine-sql-query-post`

Executes a SQL query against Workers Analytics Engine data. Pass the SQL query in the request body as plain text. The response uses newline-delimited JSON (NDJSON) by default, or a single JSON object when the query includes a FORMAT JSON clause. Prefer this method for longer queries that may exceed URL length limits.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/analytics-engine_AccountId"}]
```

## Definition

```yaml
{"operationId": "analytics-engine-sql-query-post", "summary": "Execute an Analytics Engine SQL query via request body", "description": "Executes a SQL query against Workers Analytics Engine data. Pass the SQL query in the request body as plain text. The response uses newline-delimited JSON (NDJSON) by default, or a single JSON object when the query includes a FORMAT JSON clause. Prefer this method for longer queries that may exceed URL length limits.", "requestBody": {"required": true, "content": {"text/plain": {"examples": {"aggregation_query": {"summary": "Aggregation with time filter", "value": "SELECT blob1, sum(double1) as total\nFROM my_dataset\nWHERE timestamp > NOW() - INTERVAL '7' DAY\nGROUP BY blob1\nORDER BY total DESC\nLIMIT 100\n"}, "json_format_query": {"summary": "Query with JSON output format", "value": "SELECT blob1, count() FROM my_dataset GROUP BY blob1 FORMAT JSON"}, "simple_query": {"summary": "Simple count query", "value": "SELECT count() FROM my_dataset"}}, "schema": {"description": "Supply a SQL query to run against your Analytics Engine datasets. Use only SELECT statements. The service validates the query against a restricted SQL grammar before execution.", "type": "string"}}}}, "responses": {"200": {"$ref": "#/components/responses/analytics-engine_QuerySuccess"}, "400": {"$ref": "#/components/responses/analytics-engine_BadRequest"}, "403": {"$ref": "#/components/responses/analytics-engine_Forbidden"}, "422": {"$ref": "#/components/responses/analytics-engine_UnprocessableQuery"}, "429": {"$ref": "#/components/responses/analytics-engine_TooManyRequests"}, "500": {"$ref": "#/components/responses/analytics-engine_InternalServerError"}, "503": {"$ref": "#/components/responses/analytics-engine_ServiceUnavailable"}, "504": {"$ref": "#/components/responses/analytics-engine_GatewayTimeout"}}, "security": [{"api_token": []}], "tags": ["Analytics Engine"], "x-api-token-group": ["Account Analytics Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "analytics_engine.sql", "x-fern-sdk-method-name": "query", "x-forge-hidden": true}
```
