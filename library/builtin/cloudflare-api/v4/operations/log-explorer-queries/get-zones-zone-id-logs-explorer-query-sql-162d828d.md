---
title: Run a log query
page_id: operation-get-zones-zone-id-logs-explorer-query-sql-05281e9a
path: operations/log-explorer-queries
description: Run a SQL query against zone-level datasets. The SQL query is passed as the `query` query parameter. Because the query appears in the URL, it is subject to gateway and proxy URL-length limits and may be recorded in access logs. This endpoint is deprecated in favour of the POST variant which sends the query in the request body.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/logs/explorer/query/sql
operation_ids:
    - zones-logs-explorer-query-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Run a log query

`GET /zones/{zone_id}/logs/explorer/query/sql`

Operation ID: `zones-logs-explorer-query-get`

Run a SQL query against zone-level datasets. The SQL query is passed as the `query` query parameter. Because the query appears in the URL, it is subject to gateway and proxy URL-length limits and may be recorded in access logs. This endpoint is deprecated in favour of the POST variant which sends the query in the request body.

## Definition

```yaml
{"operationId": "zones-logs-explorer-query-get", "summary": "Run a log query", "description": "Run a SQL query against zone-level datasets. The SQL query is passed as the `query` query parameter. Because the query appears in the URL, it is subject to gateway and proxy URL-length limits and may be recorded in access logs. This endpoint is deprecated in favour of the POST variant which sends the query in the request body.", "parameters": [{"$ref": "#/components/parameters/lex_ZoneID"}, {"name": "query", "in": "query", "description": "SQL query to execute.", "required": true, "schema": {"type": "string", "maxLength": 4096}}], "responses": {"200": {"description": "Query results.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_QueryResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest"}, "403": {"$ref": "#/components/responses/lex_Forbidden"}, "404": {"$ref": "#/components/responses/lex_NotFound"}, "507": {"$ref": "#/components/responses/lex_ResourcesExhausted"}}, "deprecated": true, "security": [{"api_token": []}], "tags": ["Log Explorer Queries"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-auditable": false, "x-fern-ignore": true}
```
