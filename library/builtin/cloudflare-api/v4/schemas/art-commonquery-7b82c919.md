---
title: art_CommonQuery
page_id: schema-art-commonquery-7b82c919
path: schemas
description: Defines fields that all query types share.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_CommonQuery

Defines fields that all query types share.

```yaml
{"description": "Defines fields that all query types share.", "type": "object", "properties": {"filters": {"description": "Filters to apply before aggregating results.", "type": "array", "items": {"$ref": "#/components/schemas/art_Filter"}, "example": []}, "from": {"description": "The start of the query time range (inclusive). RFC3339 format with timezone is required (e.g. `2024-11-05T00:00:00Z`).", "type": "string", "format": "date-time", "example": "2024-11-05T00:00:00Z"}, "groupBy": {"description": "Specifies the column names to group results by. Requires valid columns for the target dataset.", "type": "array", "items": {"type": "string"}, "example": ["country", "allowed"]}, "stats": {"description": "Specifies the stat names to include in results. Requires valid stats for the target dataset (e.g. `attemptsTotal`, `bytesTotal`).", "type": "array", "items": {"type": "string"}, "example": ["attemptsTotal"]}, "to": {"description": "Specifies the end of the query time range (exclusive). Requires RFC3339 format with timezone.", "type": "string", "format": "date-time", "example": "2024-11-06T00:00:00Z"}}, "required": ["from", "to", "groupBy", "stats", "filters"]}
```
