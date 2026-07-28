---
title: art_DataSecurityFindingsSummaryQuery
page_id: schema-art-datasecurityfindingssummaryquery-76772b42
path: schemas
description: Query for aggregate findings summary.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_DataSecurityFindingsSummaryQuery

Query for aggregate findings summary.

```yaml
{"description": "Query for aggregate findings summary.", "type": "object", "properties": {"filters": {"description": "Filters to apply.", "type": "array", "items": {"$ref": "#/components/schemas/art_Filter"}, "example": []}, "from": {"description": "Start of the query time range (inclusive). RFC3339.", "type": "string", "format": "date-time", "example": "2024-11-01T00:00:00Z"}, "to": {"description": "End of the query time range (exclusive). RFC3339.", "type": "string", "format": "date-time", "example": "2024-11-08T00:00:00Z"}}, "required": ["from", "to", "filters"]}
```
