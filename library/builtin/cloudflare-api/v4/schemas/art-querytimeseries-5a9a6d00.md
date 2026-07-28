---
title: art_QueryTimeseries
page_id: schema-art-querytimeseries-5a9a6d00
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_QueryTimeseries

```yaml
{"allOf": [{"$ref": "#/components/schemas/art_CommonQuery"}, {"properties": {"resolution": {"description": "Time bucket size for grouping results. Controls the granularity of the returned time slots.\n", "type": "string", "example": "hour"}}, "required": ["resolution"], "type": "object"}]}
```
