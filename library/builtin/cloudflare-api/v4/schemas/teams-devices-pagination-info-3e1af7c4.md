---
title: teams-devices_pagination_info
page_id: schema-teams-devices-pagination-info-3e1af7c4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_pagination_info

```yaml
{"type": "object", "properties": {"count": {"description": "Number of records in the response.", "type": "integer"}, "page": {"description": "The page size number of the response.", "type": "integer"}, "per_page": {"description": "The limit for the number of records in the response.", "type": "integer"}, "total_count": {"description": "Total number of records available.", "type": "integer"}, "total_pages": {"description": "Total number of pages available.", "type": "integer"}}, "example": {"count": 1, "page": 1, "per_page": 10, "total_count": 10, "total_pages": 1}, "required": ["page", "per_page", "count", "total_count"]}
```
