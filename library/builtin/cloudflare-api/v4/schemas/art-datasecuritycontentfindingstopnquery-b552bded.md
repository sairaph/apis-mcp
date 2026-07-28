---
title: art_DataSecurityContentFindingsTopNQuery
page_id: schema-art-datasecuritycontentfindingstopnquery-b552bded
path: schemas
description: Returns the top N integrations for content findings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_DataSecurityContentFindingsTopNQuery

Returns the top N integrations for content findings.

```yaml
{"description": "Returns the top N integrations for content findings.", "type": "object", "properties": {"filters": {"description": "Filters to apply. `findingType = content` is applied automatically for CASB data.", "type": "array", "items": {"$ref": "#/components/schemas/art_Filter"}, "example": []}, "from": {"description": "Start of the query time range (inclusive). RFC3339.", "type": "string", "format": "date-time", "example": "2024-11-01T00:00:00Z"}, "n": {"description": "Maximum number of integrations to return.", "type": "integer", "format": "int64", "example": 10, "minimum": 1}, "to": {"description": "End of the query time range (exclusive). RFC3339.", "type": "string", "format": "date-time", "example": "2024-11-08T00:00:00Z"}}, "required": ["from", "to", "n", "filters"]}
```
