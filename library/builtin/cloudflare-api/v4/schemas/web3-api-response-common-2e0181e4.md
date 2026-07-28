---
title: web3_api-response-common
page_id: schema-web3-api-response-common-2e0181e4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# web3_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/web3_messages"}, "messages": {"$ref": "#/components/schemas/web3_messages"}, "result": {"description": "Provides the API response.", "type": "object", "anyOf": [{"type": "object"}, {"items": {}, "type": "array"}, {"type": "string"}]}, "success": {"description": "Specifies whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
