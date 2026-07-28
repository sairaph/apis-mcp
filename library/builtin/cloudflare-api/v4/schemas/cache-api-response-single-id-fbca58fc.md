---
title: cache_api-response-single-id
page_id: schema-cache-api-response-single-id-fbca58fc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache_api-response-single-id

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/cache_messages"}, "messages": {"$ref": "#/components/schemas/cache_messages"}, "result": {"$ref": "#/components/schemas/cache_result"}, "success": {"description": "Indicates the API call's success or failure.", "type": "boolean", "example": true}}, "required": ["success", "errors", "messages", "result"]}
```
