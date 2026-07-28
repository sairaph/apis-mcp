---
title: zero-trust-gateway_api-response-common-failure
page_id: schema-zero-trust-gateway-api-response-common-failure-3ae95eca
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_api-response-common-failure

```yaml
{"type": "object", "properties": {"errors": {"type": "object", "example": [{"code": 7003, "message": "No route for the URI"}], "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_messages"}], "minLength": 1}, "messages": {"type": "object", "example": [], "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_messages"}]}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"description": "Indicate whether the API call was successful.", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages", "result"]}
```
