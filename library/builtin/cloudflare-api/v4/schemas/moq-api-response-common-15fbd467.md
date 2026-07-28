---
title: moq_api-response-common
page_id: schema-moq-api-response-common-15fbd467
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer"}, "message": {"type": "string"}}, "type": "object"}}, "messages": {"type": "array", "items": {"properties": {"code": {"type": "integer"}, "message": {"type": "string"}}, "type": "object"}}, "success": {"type": "boolean", "example": true}}, "required": ["success", "errors", "messages"]}
```
