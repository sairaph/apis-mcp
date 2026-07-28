---
title: workers_ErrorWorkerTagLengthLimit
page_id: schema-workers-errorworkertaglengthlimit-760e9e7b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerTagLengthLimit

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that the Worker has a tag that exceeds the maximum tag length.", "type": "integer", "enum": [100102]}, "message": {"description": "Message explaining why the tag is too long, including the maximum tag length.", "type": "string"}}, "required": ["code", "message"]}
```
