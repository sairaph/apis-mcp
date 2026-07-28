---
title: organizations-api_V4ErrorResponse
page_id: schema-organizations-api-v4errorresponse-71e3664e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_V4ErrorResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"type": "object"}, "success": {"type": "boolean", "enum": [false]}}, "required": ["success", "errors", "messages"]}
```
