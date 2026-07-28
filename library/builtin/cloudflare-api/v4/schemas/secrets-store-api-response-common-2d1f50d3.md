---
title: secrets-store_api-response-common
page_id: schema-secrets-store-api-response-common-2d1f50d3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secrets-store_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/secrets-store_messages"}, "messages": {"$ref": "#/components/schemas/secrets-store_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
