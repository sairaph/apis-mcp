---
title: analytics-engine_JsonFormatResponse
page_id: schema-analytics-engine-jsonformatresponse-0243c178
path: schemas
description: Response structure when FORMAT JSON is specified in the query. Contains the query result rows, column metadata, and a row count.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# analytics-engine_JsonFormatResponse

Response structure when FORMAT JSON is specified in the query. Contains the query result rows, column metadata, and a row count.

```yaml
{"description": "Response structure when FORMAT JSON is specified in the query. Contains the query result rows, column metadata, and a row count.\n", "type": "object", "properties": {"data": {"description": "Array of result rows. Each row is an object with keys corresponding to the selected columns.\n", "type": "array", "items": {"additionalProperties": true, "type": "object"}}, "meta": {"description": "Column metadata describing the name and type of each column in the result set.\n", "type": "array", "items": {"properties": {"name": {"description": "Column name.", "type": "string"}, "type": {"description": "Column data type.", "type": "string"}}, "required": ["name", "type"], "type": "object"}}, "rows": {"description": "Total number of rows in the result set.", "type": "integer"}}, "required": ["data", "meta", "rows"]}
```
