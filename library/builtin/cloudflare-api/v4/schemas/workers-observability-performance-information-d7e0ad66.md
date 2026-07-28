---
title: workers-observability_performance_information
page_id: schema-workers-observability-performance-information-d7e0ad66
path: schemas
description: Query performance statistics from the database. Includes execution time, rows scanned, and bytes read. Does not include network latency.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-observability_performance_information

Query performance statistics from the database. Includes execution time, rows scanned, and bytes read. Does not include network latency.

```yaml
{"description": "Query performance statistics from the database. Includes execution time, rows scanned, and bytes read. Does not include network latency.", "type": "object", "properties": {"abr_level": {"description": "The level of Adaptive Bit Rate (ABR) sampling used for the query. If empty the ABR level is 1", "type": "number"}, "bytes_read": {"description": "Number of uncompressed bytes read from the table.", "type": "number"}, "elapsed": {"description": "Time in seconds for the query to run.", "type": "number"}, "rows_read": {"description": "Number of rows scanned from the table.", "type": "number"}}, "required": ["elapsed", "rows_read", "bytes_read"]}
```
