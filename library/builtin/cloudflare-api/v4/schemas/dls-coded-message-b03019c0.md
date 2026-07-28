---
title: dls_coded_message
page_id: schema-dls-coded-message-b03019c0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dls_coded_message

```yaml
{"type": "object", "properties": {"code": {"type": "integer", "minimum": 1000}, "error_chain": {"description": "Optional upstream error context for APIv4 errors that wrap downstream service failures.", "type": "array", "items": {"$ref": "#/components/schemas/dls_coded_message"}}, "message": {"type": "string"}}, "required": ["code", "message"]}
```
