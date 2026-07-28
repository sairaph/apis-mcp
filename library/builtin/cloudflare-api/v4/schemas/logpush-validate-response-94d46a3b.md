---
title: logpush_validate_response
page_id: schema-logpush-validate-response-94d46a3b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# logpush_validate_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/logpush_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true, "properties": {"message": {"type": "string", "example": "", "x-auditable": true}, "valid": {"type": "boolean", "example": true, "x-auditable": true}}}}, "type": "object"}]}
```
