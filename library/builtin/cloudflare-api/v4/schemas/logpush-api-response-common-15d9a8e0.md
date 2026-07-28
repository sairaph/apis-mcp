---
title: logpush_api-response-common
page_id: schema-logpush-api-response-common-15d9a8e0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# logpush_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/logpush_messages"}, "messages": {"$ref": "#/components/schemas/logpush_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
