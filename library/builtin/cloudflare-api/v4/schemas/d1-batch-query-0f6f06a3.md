---
title: d1_batch-query
page_id: schema-d1-batch-query-0f6f06a3
path: schemas
description: A single query object or a batch query object
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_batch-query

A single query object or a batch query object

```yaml
{"description": "A single query object or a batch query object", "oneOf": [{"$ref": "#/components/schemas/d1_single-query"}, {"properties": {"batch": {"type": "array", "items": {"$ref": "#/components/schemas/d1_single-query"}}}, "required": ["batch"], "title": "multiple queries", "type": "object"}]}
```
