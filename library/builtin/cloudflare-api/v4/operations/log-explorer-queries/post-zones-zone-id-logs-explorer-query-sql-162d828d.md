---
title: Run a log query
page_id: operation-post-zones-zone-id-logs-explorer-query-sql-946aba3d
path: operations/log-explorer-queries
description: |-
    Run a SQL query against zone-level datasets.

    Timestamp fields are RFC3339 strings. Filter with:
      WHERE {timestamp_field} >= now() - INTERVAL '30' DAY
      WHERE {timestamp_field} >= '2026-04-01T00:00:00Z'
      WHERE {timestamp_field} BETWEEN '2026-04-01T00:00:00Z' AND '2026-04-30T23:59:59Z'

    List configured zone datasets to see enabled zone-level datasets.
    List available zone datasets to inspect their schemas and timestamp fields.

    For more information about the datasets, and the meaning of each field, check out https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/logs/explorer/query/sql
operation_ids:
    - zones-logs-explorer-query-post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Run a log query

`POST /zones/{zone_id}/logs/explorer/query/sql`

Operation ID: `zones-logs-explorer-query-post`

Run a SQL query against zone-level datasets.

Timestamp fields are RFC3339 strings. Filter with:
  WHERE {timestamp_field} >= now() - INTERVAL '30' DAY
  WHERE {timestamp_field} >= '2026-04-01T00:00:00Z'
  WHERE {timestamp_field} BETWEEN '2026-04-01T00:00:00Z' AND '2026-04-30T23:59:59Z'

List configured zone datasets to see enabled zone-level datasets.
List available zone datasets to inspect their schemas and timestamp fields.

For more information about the datasets, and the meaning of each field, check out https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/

## Definition

```yaml
{"operationId": "zones-logs-explorer-query-post", "summary": "Run a log query", "description": "Run a SQL query against zone-level datasets.\n\nTimestamp fields are RFC3339 strings. Filter with:\n  WHERE {timestamp_field} >= now() - INTERVAL '30' DAY\n  WHERE {timestamp_field} >= '2026-04-01T00:00:00Z'\n  WHERE {timestamp_field} BETWEEN '2026-04-01T00:00:00Z' AND '2026-04-30T23:59:59Z'\n\nList configured zone datasets to see enabled zone-level datasets.\nList available zone datasets to inspect their schemas and timestamp fields.\n\nFor more information about the datasets, and the meaning of each field, check out https://developers.cloudflare.com/logs/logpush/logpush-job/datasets/\n", "parameters": [{"$ref": "#/components/parameters/lex_ZoneID"}], "requestBody": {"required": true, "content": {"text/plain": {"schema": {"description": "SQL query to execute.", "type": "string", "format": "binary"}}}}, "responses": {"200": {"description": "Query results.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lex_QueryResponse"}}}}, "400": {"$ref": "#/components/responses/lex_BadRequest"}, "403": {"$ref": "#/components/responses/lex_Forbidden"}, "404": {"$ref": "#/components/responses/lex_NotFound"}, "507": {"$ref": "#/components/responses/lex_ResourcesExhausted"}}, "security": [{"api_token": []}], "tags": ["Log Explorer Queries"], "x-api-token-group": ["Logs Write", "Logs Read"], "x-auditable": false, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "logexplorer", "x-fern-sdk-method-name": "query", "x-forge-hidden": true, "x-stability": "beta"}
```
