---
title: observatory_api-response-common
page_id: schema-observatory-api-response-common-ea2830f5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# observatory_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/observatory_messages"}, "messages": {"$ref": "#/components/schemas/observatory_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "x-auditable": true}}, "required": ["success", "errors", "messages"]}
```
