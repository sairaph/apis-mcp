---
title: mcn_response
page_id: schema-mcn-response-af557218
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_response

```yaml
{"type": "object", "properties": {"messages": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_error"}}, "success": {"type": "boolean", "x-auditable": true}}, "required": ["result", "success", "errors", "messages"]}
```
