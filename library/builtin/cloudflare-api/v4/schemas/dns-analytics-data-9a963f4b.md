---
title: dns-analytics_data
page_id: schema-dns-analytics-data-9a963f4b
path: schemas
description: Array with one row per combination of dimension values.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-analytics_data

Array with one row per combination of dimension values.

```yaml
{"description": "Array with one row per combination of dimension values.", "type": "array", "items": {"properties": {"dimensions": {"description": "Array of dimension values, representing the combination of dimension values corresponding to this row.", "type": "array", "items": {"description": "Dimension value.", "example": "NODATA", "type": "string"}}}, "required": ["dimensions"], "type": "object"}}
```
