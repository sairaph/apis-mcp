---
title: ai-audit_api-response-common-failure
page_id: schema-ai-audit-api-response-common-failure-466eb759
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# ai-audit_api-response-common-failure

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/ai-audit_api_message"}, "example": [{"code": 7003, "message": "No route for the URI"}], "minItems": 1}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/ai-audit_api_message"}}, "result": {"type": "object", "nullable": true}, "success": {"type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages"]}
```
