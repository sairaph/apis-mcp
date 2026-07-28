---
title: workers-observability_query_run
page_id: schema-workers-observability-query-run-3713b8e2
path: schemas
description: Represents a single execution of a query against Workers Observability data, including the query definition, execution status, and performance statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-observability_query_run

Represents a single execution of a query against Workers Observability data, including the query definition, execution status, and performance statistics.

```yaml
{"description": "Represents a single execution of a query against Workers Observability data, including the query definition, execution status, and performance statistics.", "type": "object", "properties": {"accountId": {"description": "Cloudflare account ID that owns this query run.", "type": "string"}, "created": {"description": "ISO-8601 timestamp when the query run was created.", "type": "string"}, "dry": {"description": "Whether this was a dry run (results not persisted).", "type": "boolean"}, "granularity": {"description": "Number of time-series buckets used for the query. Higher values produce more detailed series data.", "type": "number"}, "id": {"description": "Unique identifier for this query run.", "type": "string"}, "query": {"description": "A saved query definition with its parameters, metadata, and ownership information.", "allOf": [{"$ref": "#/components/schemas/workers-observability_query"}, {"properties": {"created": {"anyOf": [{"type": "string"}, {"type": "string"}]}, "updated": {"anyOf": [{"type": "string"}, {"type": "string"}]}}, "type": "object"}]}, "statistics": {"description": "Query performance statistics from the database (does not include network latency).", "type": "object", "properties": {"abr_level": {"description": "The level of Adaptive Bit Rate (ABR) sampling used for the query. If empty the ABR level is 1", "type": "number"}, "bytes_read": {"description": "Number of uncompressed bytes read from the table.", "type": "number"}, "elapsed": {"description": "Time in seconds for the query to run.", "type": "number"}, "rows_read": {"description": "Number of rows scanned from the table.", "type": "number"}}, "required": ["elapsed", "rows_read", "bytes_read"]}, "status": {"description": "Current execution status of the query run.", "type": "string", "enum": ["STARTED", "COMPLETED"]}, "timeframe": {"description": "Time range for the query execution", "type": "object", "properties": {"from": {"description": "Start timestamp for the query timeframe (Unix timestamp in milliseconds)", "type": "number"}, "to": {"description": "End timestamp for the query timeframe (Unix timestamp in milliseconds)", "type": "number"}}, "required": ["to", "from"]}, "updated": {"description": "ISO-8601 timestamp when the query run was last updated.", "type": "string"}, "userId": {"description": "ID of the user who initiated the query run.", "type": "string"}}, "required": ["id", "query", "accountId", "timeframe", "userId", "status", "granularity", "dry"]}
```
