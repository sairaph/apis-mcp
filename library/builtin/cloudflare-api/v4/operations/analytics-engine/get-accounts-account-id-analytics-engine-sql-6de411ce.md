---
title: Execute an Analytics Engine SQL query via query parameter
page_id: operation-get-accounts-account-id-analytics-engine-sql-9f13641a
path: operations/analytics-engine
description: Executes a SQL query against Workers Analytics Engine data. Pass the SQL query as a query parameter. The response uses newline-delimited JSON (NDJSON) by default, or a single JSON object when the query includes a FORMAT JSON clause. Use the POST method for longer queries that may exceed URL length limits.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/analytics_engine/sql
operation_ids:
    - analytics-engine-sql-query-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Execute an Analytics Engine SQL query via query parameter

`GET /accounts/{account_id}/analytics_engine/sql`

Operation ID: `analytics-engine-sql-query-get`

Executes a SQL query against Workers Analytics Engine data. Pass the SQL query as a query parameter. The response uses newline-delimited JSON (NDJSON) by default, or a single JSON object when the query includes a FORMAT JSON clause. Use the POST method for longer queries that may exceed URL length limits.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/analytics-engine_AccountId"}]
```

## Definition

```yaml
{"operationId": "analytics-engine-sql-query-get", "summary": "Execute an Analytics Engine SQL query via query parameter", "description": "Executes a SQL query against Workers Analytics Engine data. Pass the SQL query as a query parameter. The response uses newline-delimited JSON (NDJSON) by default, or a single JSON object when the query includes a FORMAT JSON clause. Use the POST method for longer queries that may exceed URL length limits.", "parameters": [{"name": "query", "in": "query", "description": "Supply a SQL query to run against your Analytics Engine datasets. Use only SELECT statements. The service validates the query against a restricted SQL grammar before execution.", "required": true, "schema": {"type": "string", "example": "SELECT blob1, count() FROM my_dataset WHERE timestamp > NOW() - INTERVAL '1' DAY GROUP BY blob1 ORDER BY count() DESC LIMIT 10"}}], "responses": {"200": {"$ref": "#/components/responses/analytics-engine_QuerySuccess"}, "400": {"$ref": "#/components/responses/analytics-engine_BadRequest"}, "403": {"$ref": "#/components/responses/analytics-engine_Forbidden"}, "422": {"$ref": "#/components/responses/analytics-engine_UnprocessableQuery"}, "429": {"$ref": "#/components/responses/analytics-engine_TooManyRequests"}, "500": {"$ref": "#/components/responses/analytics-engine_InternalServerError"}, "503": {"$ref": "#/components/responses/analytics-engine_ServiceUnavailable"}, "504": {"$ref": "#/components/responses/analytics-engine_GatewayTimeout"}}, "security": [{"api_token": []}], "tags": ["Analytics Engine"], "x-api-token-group": ["Account Analytics Read"], "x-fern-availability": "generally-available", "x-fern-ignore": true, "x-fern-sdk-group-name": "analytics_engine.sql", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
