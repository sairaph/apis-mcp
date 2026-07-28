---
title: workers_ErrorWorkerNamePreviewLengthLimit
page_id: schema-workers-errorworkernamepreviewlengthlimit-ef9b1a08
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerNamePreviewLengthLimit

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that the Worker name is too long to be used with previews enabled.", "type": "integer", "enum": [100315]}, "message": {"description": "Message explaining that Worker names with previews enabled cannot exceed 54 characters.", "type": "string"}}, "required": ["code", "message"]}
```
