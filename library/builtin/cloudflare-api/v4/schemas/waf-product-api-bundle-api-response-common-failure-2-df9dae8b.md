---
title: waf-product-api-bundle_api-response-common-failure-2
page_id: schema-waf-product-api-bundle-api-response-common-failure-2-df9dae8b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-product-api-bundle_api-response-common-failure-2

```yaml
{"type": "object", "properties": {"errors": {"example": [{"code": 10000, "message": "Authentication error"}], "allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_messages"}], "minLength": 1}, "messages": {"example": [], "allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_messages"}]}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages", "result"]}
```
