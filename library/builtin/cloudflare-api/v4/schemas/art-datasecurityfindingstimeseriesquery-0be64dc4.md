---
title: art_DataSecurityFindingsTimeseriesQuery
page_id: schema-art-datasecurityfindingstimeseriesquery-0be64dc4
path: schemas
description: Query for findings timeseries.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_DataSecurityFindingsTimeseriesQuery

Query for findings timeseries.

```yaml
{"description": "Query for findings timeseries.", "type": "object", "properties": {"filters": {"description": "Filters to apply.", "type": "array", "items": {"$ref": "#/components/schemas/art_Filter"}, "example": []}, "from": {"description": "Start of the query time range (inclusive). RFC3339.", "type": "string", "format": "date-time", "example": "2024-11-01T00:00:00Z"}, "to": {"description": "End of the query time range (exclusive). RFC3339.", "type": "string", "format": "date-time", "example": "2024-11-08T00:00:00Z"}}, "required": ["from", "to", "filters"]}
```
