---
title: workers_ErrorWorkerLimit
page_id: schema-workers-errorworkerlimit-5256b7a4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerLimit

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that the account has exceeded the maximum number of Workers allowed.", "type": "integer", "enum": [10037]}, "message": {"description": "Message explaining that the Worker limit has been exceeded and providing guidance.", "type": "string"}}, "required": ["code", "message"]}
```
