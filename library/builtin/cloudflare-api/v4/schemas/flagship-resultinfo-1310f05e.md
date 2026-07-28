---
title: flagship_ResultInfo
page_id: schema-flagship-resultinfo-1310f05e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# flagship_ResultInfo

```yaml
{"type": "object", "properties": {"count": {"description": "Number of items returned in this page.", "type": "integer", "minimum": 0}, "cursor": {"description": "Cursor to pass back to fetch the next page, or null when this is the last page.", "type": "string", "nullable": true}}, "required": ["count", "cursor"]}
```
