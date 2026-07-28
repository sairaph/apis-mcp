---
title: dls_PrefixBindingPaginatedListResponse
page_id: schema-dls-prefixbindingpaginatedlistresponse-7d81a393
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dls_PrefixBindingPaginatedListResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/dls_coded_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/dls_coded_message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/dls_PrefixBinding"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Number of items in the current page.", "type": "integer"}, "cursor": {"description": "Opaque cursor for the next page. Empty string when there are no more results.", "type": "string"}, "per_page": {"description": "Maximum number of items per page.", "type": "integer"}}, "required": ["count", "cursor", "per_page"]}, "success": {"type": "boolean"}}, "required": ["success", "errors", "messages", "result", "result_info"]}
```
