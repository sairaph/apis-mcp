---
title: api-shield_api-response-common
page_id: schema-api-shield-api-response-common-ac90d9af
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/api-shield_messages"}, "messages": {"$ref": "#/components/schemas/api-shield_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
