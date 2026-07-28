---
title: waitingroom_api-response-common-2
page_id: schema-waitingroom-api-response-common-2-61673431
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_api-response-common-2

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/waitingroom_messages"}, "messages": {"$ref": "#/components/schemas/waitingroom_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
