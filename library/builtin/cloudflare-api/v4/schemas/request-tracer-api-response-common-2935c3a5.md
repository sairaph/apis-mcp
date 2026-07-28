---
title: request-tracer_api-response-common
page_id: schema-request-tracer-api-response-common-2935c3a5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# request-tracer_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/request-tracer_messages"}, "messages": {"$ref": "#/components/schemas/request-tracer_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
