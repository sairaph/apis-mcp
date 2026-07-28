---
title: art_QueryTopN
page_id: schema-art-querytopn-f28bd81b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_QueryTopN

```yaml
{"allOf": [{"$ref": "#/components/schemas/art_CommonQuery"}, {"properties": {"n": {"description": "Maximum number of results to return.", "type": "integer", "format": "int64", "example": 10, "minimum": 1}, "orderBy": {"description": "Specifies the stat name for sorting results in descending order. Requires a valid stat for the target dataset.", "type": "string", "example": "attemptsTotal"}}, "required": ["n", "orderBy"], "type": "object"}]}
```
