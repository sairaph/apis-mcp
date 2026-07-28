---
title: workers_ErrorWorkerTagLimit
page_id: schema-workers-errorworkertaglimit-3e23cb57
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerTagLimit

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that the Worker has exceeded the maximum number of tags allowed.", "type": "integer", "enum": [100103]}, "message": {"description": "Message explaining that the tag limit has been exceeded and suggesting to remove a tag.", "type": "string"}}, "required": ["code", "message"]}
```
