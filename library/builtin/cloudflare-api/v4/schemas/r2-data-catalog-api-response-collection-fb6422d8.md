---
title: r2-data-catalog_api-response-collection
page_id: schema-r2-data-catalog-api-response-collection-fb6422d8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_api-response-collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/r2-data-catalog_api-response-single"}, {"properties": {"result_info": {"type": "object", "properties": {"count": {"description": "Indicates the number of results in this page.", "type": "integer"}, "page": {"description": "Specifies the current page number.", "type": "integer"}, "per_page": {"description": "Specifies the number of results per page.", "type": "integer"}, "total_count": {"description": "Specifies the total number of results.", "type": "integer"}}}}, "type": "object"}]}
```
