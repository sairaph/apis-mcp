---
title: workers_ErrorWorkerTagInvalid
page_id: schema-workers-errorworkertaginvalid-6db2466f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_ErrorWorkerTagInvalid

```yaml
{"type": "object", "properties": {"code": {"description": "Code indicating that the Worker has a tag containing invalid characters.", "type": "integer", "enum": [100134]}, "message": {"description": "Message explaining that tags cannot contain certain characters like comma or ampersand.", "type": "string"}}, "required": ["code", "message"]}
```
