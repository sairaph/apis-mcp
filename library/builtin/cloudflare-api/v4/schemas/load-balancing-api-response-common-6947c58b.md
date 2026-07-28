---
title: load-balancing_api-response-common
page_id: schema-load-balancing-api-response-common-6947c58b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# load-balancing_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/load-balancing_messages"}, "messages": {"$ref": "#/components/schemas/load-balancing_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
