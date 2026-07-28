---
title: waf-product-api-bundle_api-response-common-failure-3
page_id: schema-waf-product-api-bundle-api-response-common-failure-3-5d2145bd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-product-api-bundle_api-response-common-failure-3

```yaml
{"type": "object", "properties": {"errors": {"example": [{"code": 13900, "message": "internal server error"}], "allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_messages"}], "minItems": 1}, "messages": {"example": [], "allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_messages"}]}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"description": "Defines whether the API call was successful.", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages", "result"]}
```
