---
title: intel_schemas-api-response-common
page_id: schema-intel-schemas-api-response-common-8c6c0c9c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_schemas-api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/intel_schemas-messages"}, "messages": {"$ref": "#/components/schemas/intel_schemas-messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
