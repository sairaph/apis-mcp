---
title: workers_ErrorWorkerInvalid
page_id: schema-workers-errorworkerinvalid-361d5d0c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerInvalid

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that the Worker request contains invalid data.", "type": "integer", "enum": [10021]}, "message": {"description": "Message explaining why the Worker request is invalid, such as malformed JSON.", "type": "string"}}, "required": ["code", "message"]}
```
