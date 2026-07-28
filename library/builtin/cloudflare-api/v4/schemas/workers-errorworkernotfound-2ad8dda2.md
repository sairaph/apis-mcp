---
title: workers_ErrorWorkerNotFound
page_id: schema-workers-errorworkernotfound-2ad8dda2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerNotFound

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that the Worker does not exist.", "type": "integer", "enum": [10007]}, "message": {"description": "Message explaining that the Worker was not found.", "type": "string"}}, "required": ["code", "message"]}
```
