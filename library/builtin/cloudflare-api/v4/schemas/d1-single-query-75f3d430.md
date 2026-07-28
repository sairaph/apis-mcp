---
title: d1_single-query
page_id: schema-d1-single-query-75f3d430
path: schemas
description: A single query with or without parameters
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# d1_single-query

A single query with or without parameters

```yaml
{"description": "A single query with or without parameters", "type": "object", "properties": {"params": {"$ref": "#/components/schemas/d1_params"}, "sql": {"$ref": "#/components/schemas/d1_sql"}}, "required": ["sql"], "title": "single query"}
```
