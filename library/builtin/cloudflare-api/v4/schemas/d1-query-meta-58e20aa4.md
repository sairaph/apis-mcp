---
title: d1_query-meta
page_id: schema-d1-query-meta-58e20aa4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_query-meta

```yaml
{"type": "object", "properties": {"changed_db": {"description": "Denotes if the database has been altered in some way, like deleting rows.", "type": "boolean", "x-auditable": true}, "changes": {"description": "Rough indication of how many rows were modified by the query, as provided by SQLite's `sqlite3_total_changes()`.", "type": "number", "x-auditable": true}, "duration": {"description": "The duration of the SQL query execution inside the database. Does not include any network communication.", "type": "number", "x-auditable": true}, "last_row_id": {"description": "The row ID of the last inserted row in a table with an `INTEGER PRIMARY KEY` as provided by SQLite. Tables created with `WITHOUT ROWID` do not populate this.", "type": "number", "x-auditable": true}, "rows_read": {"description": "Number of rows read during the SQL query execution, including indices (not all rows are necessarily returned).", "type": "number", "x-auditable": true}, "rows_written": {"description": "Number of rows written during the SQL query execution, including indices.", "type": "number", "x-auditable": true}, "served_by_colo": {"$ref": "#/components/schemas/d1_served-by-colo"}, "served_by_primary": {"description": "Denotes if the query has been handled by the database primary instance.", "type": "boolean", "x-auditable": true}, "served_by_region": {"$ref": "#/components/schemas/d1_served-by-region"}, "size_after": {"description": "Size of the database after the query committed, in bytes.", "type": "number", "x-auditable": true}, "timings": {"description": "Various durations for the query.", "type": "object", "properties": {"sql_duration_ms": {"description": "The duration of the SQL query execution inside the database. Does not include any network communication.", "type": "number", "x-auditable": true}}}}}
```
