---
title: realtimekit_PagingResponse
page_id: schema-realtimekit-pagingresponse-2b63b910
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_PagingResponse

```yaml
{"type": "object", "properties": {"data": {"type": "array", "items": {"type": "object"}}, "paging": {"type": "object", "properties": {"end_offset": {"type": "number", "example": 30}, "start_offset": {"type": "number", "example": 1}, "total_count": {"type": "number", "example": 30, "minimum": 0}}, "required": ["total_count", "start_offset", "end_offset"]}, "success": {"type": "boolean", "example": true}}, "required": ["success", "data", "paging"], "title": "PagingResponse"}
```
