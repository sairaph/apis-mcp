---
title: ai-audit_api-response-common
page_id: schema-ai-audit-api-response-common-048318c4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# ai-audit_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/ai-audit_api_message"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/ai-audit_api_message"}}, "success": {"type": "boolean", "example": true}}, "required": ["success", "errors", "messages"]}
```
