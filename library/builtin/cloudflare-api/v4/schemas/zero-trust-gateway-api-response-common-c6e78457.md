---
title: zero-trust-gateway_api-response-common
page_id: schema-zero-trust-gateway-api-response-common-c6e78457
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/zero-trust-gateway_messages"}, "messages": {"$ref": "#/components/schemas/zero-trust-gateway_messages"}, "success": {"description": "Indicate whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
