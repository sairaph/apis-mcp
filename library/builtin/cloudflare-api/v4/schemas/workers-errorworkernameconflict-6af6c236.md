---
title: workers_ErrorWorkerNameConflict
page_id: schema-workers-errorworkernameconflict-6af6c236
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerNameConflict

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that a Worker with this name already exists.", "type": "integer", "enum": [10040]}, "message": {"description": "Message explaining that the Worker name is already in use and suggesting to choose a different name.", "type": "string"}}, "required": ["code", "message"]}
```
