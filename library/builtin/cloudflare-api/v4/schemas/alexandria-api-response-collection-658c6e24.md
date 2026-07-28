---
title: alexandria_api-response-collection
page_id: schema-alexandria-api-response-collection-658c6e24
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# alexandria_api-response-collection

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/alexandria_api-response-common"}, {"properties": {"result_info": {"type": "object", "properties": {"count": {"description": "Returns the total number of results for the requested service.", "type": "number", "example": 1}, "page": {"description": "Returns the current page within paginated list of results.", "type": "number", "example": 1}, "per_page": {"description": "Returns the number of results per page of results.", "type": "number", "example": 20}, "total_count": {"description": "Returns the total results available without any search parameters.", "type": "number", "example": 2000}}}}, "type": "object"}]}
```
