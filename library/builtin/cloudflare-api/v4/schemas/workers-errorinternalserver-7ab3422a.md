---
title: workers_ErrorInternalServer
page_id: schema-workers-errorinternalserver-7ab3422a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorInternalServer

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that an unknown internal server error has occurred.", "type": "integer", "enum": [10002]}, "message": {"description": "Message explaining that an unknown error occurred and providing guidance for reporting the issue.", "type": "string"}}, "required": ["code", "message"]}
```
