---
title: DataInfo
page_id: schema-datainfo-2813858c
path: schemas
description: Document basic information
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# DataInfo

Document basic information

```yaml
{"type": "object", "description": "Document basic information", "properties": {"num_pages": {"type": "integer", "description": "Total number of document pages", "example": 5}, "pages": {"type": "array", "description": "Document page count information", "items": {"$ref": "#/components/schemas/PageInfo"}}}, "required": ["num_pages"]}
```
