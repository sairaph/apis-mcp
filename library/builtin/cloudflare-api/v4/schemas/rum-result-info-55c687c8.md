---
title: rum_result_info
page_id: schema-rum-result-info-55c687c8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rum_result_info

```yaml
{"type": "object", "properties": {"count": {"description": "The total number of items on the current page.", "type": "integer", "example": 10}, "page": {"description": "Current page within the paginated list of results.", "type": "integer", "example": 1}, "per_page": {"description": "The maximum number of items to return per page of results.", "type": "integer", "example": 10}, "total_count": {"description": "The total number of items.", "type": "integer", "example": 25}, "total_pages": {"description": "The total number of pages.", "type": "integer", "example": 3, "nullable": true}}}
```
