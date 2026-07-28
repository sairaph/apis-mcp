---
title: custom-indicator-feeds_api-response-common
page_id: schema-custom-indicator-feeds-api-response-common-c1e8dbb3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-indicator-feeds_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/custom-indicator-feeds_messages-2"}, "messages": {"$ref": "#/components/schemas/custom-indicator-feeds_messages-2"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
