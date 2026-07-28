---
title: google-tag-gateway_api-response-common
page_id: schema-google-tag-gateway-api-response-common-06e41f67
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# google-tag-gateway_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/google-tag-gateway_messages"}, "messages": {"$ref": "#/components/schemas/google-tag-gateway_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
