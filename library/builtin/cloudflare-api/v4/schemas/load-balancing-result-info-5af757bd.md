---
title: load-balancing_result_info
page_id: schema-load-balancing-result-info-5af757bd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_result_info

```yaml
{"type": "object", "properties": {"count": {"description": "Total number of results on the current page.", "type": "number", "example": 20}, "page": {"description": "Current page within paginated list of results.", "type": "number", "example": 1}, "per_page": {"description": "Number of results per page.", "type": "number", "example": 20}, "total_count": {"description": "Total results available without any search parameters.", "type": "number", "example": 2000}, "total_pages": {"description": "Total number of pages available.", "type": "number", "example": 100}}}
```
