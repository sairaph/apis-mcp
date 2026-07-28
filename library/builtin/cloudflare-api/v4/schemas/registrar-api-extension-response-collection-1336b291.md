---
title: registrar-api_extension-response-collection
page_id: schema-registrar-api-extension-response-collection-1336b291
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api_extension-response-collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/registrar-api_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/registrar-api_extension-item"}}, "result_info": {"description": "Cursor-based pagination metadata.", "type": "object", "properties": {"count": {"description": "Number of items in the current result set.", "type": "integer"}, "cursor": {"description": "Opaque token for the next page. Empty string when no more pages.", "type": "string"}, "per_page": {"description": "Number of items per page.", "type": "integer"}}}}, "type": "object"}]}
```
