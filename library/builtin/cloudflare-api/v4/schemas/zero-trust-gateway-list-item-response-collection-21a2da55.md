---
title: zero-trust-gateway_list_item_response_collection
page_id: schema-zero-trust-gateway-list-item-response-collection-21a2da55
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_list_item_response_collection

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/zero-trust-gateway_items"}}, "type": "object"}, {"properties": {"result_info": {"type": "object", "properties": {"count": {"description": "Shows the total results returned based on your search parameters.", "type": "number", "example": 1}, "page": {"description": "Show the current page within paginated list of results.", "type": "number", "example": 1}, "per_page": {"description": "Show the number of results per page of results.", "type": "number", "example": 20}, "total_count": {"description": "Show the total results available without any search parameters.", "type": "number", "example": 2000}}}}, "type": "object"}]}
```
