---
title: resource-tagging_cursor_result_info
page_id: schema-resource-tagging-cursor-result-info-4cc81357
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_cursor_result_info

```yaml
{"type": "object", "properties": {"count": {"description": "Indicates the number of results returned in the current page.", "type": "integer", "example": 20}, "cursor": {"description": "Provides a cursor for the next page of results. Include this value in the next request to continue pagination.", "type": "string", "example": "eyJhY2NvdW50X2lkIjoxMjM0NTY3ODkwfQ", "nullable": true}}}
```
