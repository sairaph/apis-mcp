---
title: dls_RegionPublicPaginatedListResponse
page_id: schema-dls-regionpublicpaginatedlistresponse-32d98ace
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dls_RegionPublicPaginatedListResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/dls_coded_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/dls_coded_message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/dls_RegionPublic"}}, "result_info": {"type": "object", "properties": {"count": {"description": "Number of items in the current page.", "type": "integer"}, "cursor": {"description": "Opaque cursor for the next page. Empty string when there are no more results.", "type": "string"}, "per_page": {"description": "Maximum number of items per page.", "type": "integer"}}, "required": ["count", "cursor", "per_page"]}, "success": {"type": "boolean"}}, "required": ["success", "errors", "messages", "result", "result_info"]}
```
