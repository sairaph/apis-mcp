---
title: snippets_ResultInfo
page_id: schema-snippets-resultinfo-749a11a4
path: schemas
description: Additional information to navigate the results.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# snippets_ResultInfo

Additional information to navigate the results.

```yaml
{"description": "Additional information to navigate the results.", "type": "object", "properties": {"count": {"description": "Specify the number of results in the current page.", "type": "integer", "example": 25, "minimum": 0, "title": "Count", "x-auditable": true}, "page": {"$ref": "#/components/schemas/snippets_Page"}, "per_page": {"$ref": "#/components/schemas/snippets_PerPage"}, "total_count": {"description": "Specify the total number of results.", "type": "integer", "example": 100, "minimum": 0, "title": "Total Count", "x-auditable": true}, "total_pages": {"description": "Specify the total number of pages.", "type": "integer", "example": 10, "minimum": 1, "title": "Total Pages", "x-auditable": true}}, "required": ["page", "per_page", "total_pages", "count", "total_count"], "title": "Result Info"}
```
