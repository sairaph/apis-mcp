---
title: workers_ErrorAuth
page_id: schema-workers-errorauth-09b66036
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorAuth

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that the user is not authorized to perform this action.", "type": "integer", "enum": [10023]}, "message": {"description": "Message explaining that the user lacks access to this feature.", "type": "string"}}, "required": ["code", "message"]}
```
