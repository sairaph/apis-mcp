---
title: r2_v4_response
page_id: schema-r2-v4-response-032adfd2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_v4_response

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/r2_errors"}, "messages": {"$ref": "#/components/schemas/r2_messages"}, "result": {"type": "object"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
