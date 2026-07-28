---
title: resource-sharing_api-response-common
page_id: schema-resource-sharing-api-response-common-494527de
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-sharing_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/resource-sharing_v4errors"}, "result": {"anyOf": [{"type": "object"}, {"items": {}, "type": "array"}, {"type": "string"}]}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true}}, "required": ["success", "errors"]}
```
