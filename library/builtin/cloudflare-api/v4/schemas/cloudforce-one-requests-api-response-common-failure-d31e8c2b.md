---
title: cloudforce-one-requests_api-response-common-failure
page_id: schema-cloudforce-one-requests-api-response-common-failure-d31e8c2b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_api-response-common-failure

```yaml
{"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"errors": {"type": "object", "properties": {"code": {"type": "integer", "example": 10433}, "message": {"type": "string", "example": "request error"}}}, "success": {"type": "boolean", "example": false}}, "type": "object"}]}
```
