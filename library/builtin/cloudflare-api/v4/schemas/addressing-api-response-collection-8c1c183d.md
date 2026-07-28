---
title: addressing_api-response-collection
page_id: schema-addressing-api-response-collection-8c1c183d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_api-response-collection

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/addressing_api-response-common"}, {"properties": {"result_info": {"type": "object", "properties": {"count": {"description": "Total number of results for the requested service.", "type": "number", "example": 1}, "page": {"description": "Current page within paginated list of results.", "type": "number", "example": 1}, "per_page": {"description": "Number of results per page of results.", "type": "number", "example": 20}, "total_count": {"description": "Total results available without any search parameters.", "type": "number", "example": 2000}, "total_pages": {"description": "The number of total pages in the entire result set.", "type": "number", "example": 100}}}}, "type": "object"}]}
```
