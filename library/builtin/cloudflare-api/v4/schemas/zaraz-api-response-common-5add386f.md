---
title: zaraz_api-response-common
page_id: schema-zaraz-api-response-common-5add386f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/zaraz_messages"}, "messages": {"$ref": "#/components/schemas/zaraz_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "x-auditable": true}}, "required": ["success", "errors", "messages", "result"]}
```
