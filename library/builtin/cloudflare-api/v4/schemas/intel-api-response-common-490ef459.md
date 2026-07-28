---
title: intel_api-response-common
page_id: schema-intel-api-response-common-490ef459
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/intel_messages"}, "messages": {"$ref": "#/components/schemas/intel_messages"}, "result": {"anyOf": [{"type": "object"}, {"items": {"oneOf": [{"type": "string"}, {"type": "object"}]}, "type": "array"}, {"type": "string"}]}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
