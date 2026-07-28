---
title: art_Filter
page_id: schema-art-filter-eee96a58
path: schemas
description: A filter to apply to the query.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_Filter

A filter to apply to the query.

```yaml
{"description": "A filter to apply to the query.", "type": "object", "properties": {"name": {"description": "Specifies the column name to filter on. Requires a valid column for the target dataset (e.g. `country`, `allowed`, `appId`).", "type": "string", "example": "country"}, "op": {"description": "Filter operator. Common values: `eq`, `neq`, `in`, `not_in`, `gt`, `lt`, `gte`, `lte`.\n", "type": "string", "example": "in"}, "values": {"description": "Values to match against. Type depends on the column.", "type": "array", "items": {"oneOf": [{"type": "string"}, {"type": "boolean"}, {"type": "integer"}, {"type": "number"}]}, "example": ["US", "CA", "GB"]}}, "required": ["name", "op", "values"]}
```
