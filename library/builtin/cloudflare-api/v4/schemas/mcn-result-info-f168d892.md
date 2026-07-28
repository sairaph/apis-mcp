---
title: mcn_result_info
page_id: schema-mcn-result-info-f168d892
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_result_info

```yaml
{"type": "object", "properties": {"count": {"description": "The number of items in the current result set.", "type": "integer", "example": 1}, "page": {"description": "The current page (starts from zero).", "type": "integer", "example": 1}, "per_page": {"description": "The maximum number of items per page.", "type": "integer", "example": 20}, "total_count": {"description": "The total number of items in the entire result set.", "type": "integer", "example": 2000}, "total_pages": {"description": "The number of total pages in the entire result set.", "type": "integer", "example": 200}}, "required": ["page", "per_page", "count", "total_count"]}
```
